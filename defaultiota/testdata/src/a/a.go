package a

const (
	StatusUnknown int = iota // want `StatusUnknown has invalid suffix for zero value with iota. suffix must be Invalid`
	StatusEnabled
)

func main() {
}
