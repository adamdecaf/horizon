package parsing

type SingleParser interface {
	Parse(input string) (string, error)
}

type MultiParser interface {
	Parse(input string) ([]string, error)
}
