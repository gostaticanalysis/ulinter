package defaultiota

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "defaultiota",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "defaultiota is ..."

var validIotaSuffix string

func init() {
	Analyzer.Flags.StringVar(&validIotaSuffix, "suffix", "Invalid", "valid suffix for iota")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ValueSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ValueSpec:
			// check is iota
			if len(n.Values) != 1 {
				return
			}
			val, ok := n.Values[0].(*ast.Ident)
			if !ok {
				return
			}
			if val.Name != "iota" {
				return
			}

			if len(n.Names) != 1 {
				return
			}
			name := n.Names[0]
			if strings.HasSuffix(name.Name, validIotaSuffix) {
				return
			}
			pass.Reportf(n.Pos(), "%s has invalid suffix for zero value with iota. suffix must be %s", name.Name, validIotaSuffix)
		}
	})

	return nil, nil
}
