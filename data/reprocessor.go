package data

import (
	"log"
)

type ReProcessor interface {
	Run() *error
}

func StartReProcessor(p ReProcessor) *error {
	if err := p.Run(); err != nil {
		log.Printf("error in reprocessor run err=%s\n", *err)
		return err
	}
	return nil
}
