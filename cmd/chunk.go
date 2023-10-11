package cmd

import "fmt"

type chunk string

const (
	chunkTagImport = chunk("I")
)

type chunkImport struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Chunks []string `json:"chunks"`
}

func (ci chunkImport) String() string {
	return fmt.Sprintf("Import id:%s name:%s chunks[ %d elements ]", ci.Id, ci.Name, len(ci.Chunks))
}

func checkChunkTag(tag uint8) chunk {
	switch tag {
	case 0x49:
		return chunkTagImport
	default:
		return ""
	}
}
