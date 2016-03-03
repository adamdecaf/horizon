package parsing

import(
	"log"
)

type NullSingleParser struct {
	SingleParser
}

func (p NullSingleParser) Parse(input string) (string, error) {
	log.Printf("NullSingleParser.Run() -- input=%s\n", input)
	return "", nil
}

type NullMultiParser struct {
	MultiParser
}

func (p NullMultiParser) Parse(input string) ([]string, error) {
	log.Printf("NullMultiParser.Run() -- input=%s\n", input)
	return nil, nil
}
