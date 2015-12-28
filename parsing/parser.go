package parsing

type Parser interface {
	Parse(input string) *string
}
