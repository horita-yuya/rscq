package react

// Entry
// https://github.com/facebook/react/blob/main/packages/react-client/src/ReactFlightClient.js#L1112
// https://github.com/facebook/react/blob/main/packages/react-client/src/ReactFlightClient.js#L996C11-L996C11
type row string

const (
	chunkTagImport = row("I")
)

// https://github.com/facebook/react/blob/main/packages/react-server-dom-webpack/src/shared/ReactFlightImportMetadata.js#L19
// TODO: AsyncImport
type chunkImport struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Chunks []string `json:"chunks"`
}

func checkChunkTag(tag uint8) row {
	switch tag {
	case 0x49:
		return chunkTagImport
	default:
		return ""
	}
}
