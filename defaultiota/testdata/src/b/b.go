package b

const (
	StatusInvalid int = iota
	StatusValid
)

func IsValid(i int) bool {
	return i == StatusValid
}
