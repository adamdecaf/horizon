package reprocess

import (
	"log"
)

func SpawnProcessors() {
	log.Println("[reprocessors] Spawning processors")

	go SpawnNullProcessor()
	go SpawnTwitterMentionProcessor()
}
