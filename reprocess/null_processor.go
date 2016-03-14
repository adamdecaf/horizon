package reprocess

import (
	"log"
)

type NullProcessor struct {
	Processor
}

func (p NullProcessor) Run() *error {
	log.Println("NullProcessor.Run()")
	return nil
}

func SpawnNullProcessor() *error {
	log.Printf("[Spawn] NullProcessor")
	processor := NullProcessor{}
	return StartProcessor(processor)
}
