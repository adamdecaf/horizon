package reprocess

import (
	"log"
)

type Processor interface {
	Run() *error
}

func StartProcessor(p Processor) *error {
	if err := p.Run(); err != nil {
		log.Printf("error in processor run err=%s\n", *err)
		return err
	}
	return nil
}
