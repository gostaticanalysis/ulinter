package typeassert

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "typeassert",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "typeassert is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			var hasTypeAssertion bool
			for _, node := range n.Rhs {
				_, ok := node.(*ast.TypeAssertExpr)
				hasTypeAssertion = hasTypeAssertion || ok
			}
			if !hasTypeAssertion {
				return
			}

			// if right hand has 2 or more values, assign statement can't assert boolean value which describes type assertion is succeeded
			if len(n.Rhs) > 1 {
				pass.Reportf(n.Pos(), "type assertion must be checked")
			}
			if len(n.Lhs) == 2 {
				return
			}

			// TypeAssertionExpr.Type is nil means foo.(type)
			tae, ok := n.Rhs[0].(*ast.TypeAssertExpr)
			if !ok {
				return
			}
			if tae.Type == nil {
				return
			}
			pass.Reportf(n.Pos(), "type assertion must be checked")
		}
	})

	return nil, nil
}
