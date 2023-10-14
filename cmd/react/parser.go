package react

import (
	"bufio"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
)

func decodeValue(value string) interface{} {
	if checkChunkTag(value[0]) == chunkTagImport {
		val := value[1:]
		var res chunkImport
		err := json.Unmarshal([]byte(val), &res)

		if err != nil {
			//panic(err)
			fmt.Println(err)
		}
		return res
	} else if value[0] == '[' {
		var res []interface{}
		err := json.Unmarshal([]byte(value), &res)

		if err != nil {
			fmt.Println(err)
		}

		return res
	} else if value[0] == 'H' {
		// https://github.com/facebook/react/blob/be67db46b60d94f9fbefccf2523429af25873e5b/packages/react-server/src/ReactFlightServer.js#L1364
		return value
	} else if value == "null" {
		return nil
	} else {
		return value
	}
}

func splitEachLine(rsc string) []RSCLine {
	var lines []RSCLine
	sc := bufio.NewScanner(strings.NewReader(rsc))
	for sc.Scan() {
		id, value, _ := strings.Cut(sc.Text(), ":")
		decoded := decodeValue(value)

		lines = append(lines, RSCLine{id, decoded})
	}
	return lines
}

// https://github.com/facebook/react/blob/main/packages/react-client/src/ReactFlightClient.js#L1132
type RSCLine struct {
	Id    string
	Value interface{}
}

func prettifyJSX(jsx *JSXElement) {
	props := jsx.Props
	placeholderKeys := []string{}
	deleteKeys := []string{
		"error", "errorStyles", "loading", "loadingStyles",
		"notFound", "notFoundStyles", "hasLoading", "styles", "style", "templateStyles",
	}
	for k := range props {
		if slices.Contains(placeholderKeys, k) {
			props[k] = "{{" + k + "}}"
		} else if slices.Contains(deleteKeys, k) {
			delete(props, k)
		} else {
			v := props[k]
			prettify(&v)
			props[k] = v
		}
	}
}

func prettify(value *any) {
	switch (*value).(type) {
	case []interface{}:
		if len((*value).([]interface{})) > 0 && (*value).([]interface{})[0] == "$" {
			jsx := ParseJsx((*value).([]interface{}))
			prettifyJSX(&jsx)
			*value = jsx
		} else {
			for i := range (*value).([]interface{}) {
				prettify(&(*value).([]interface{})[i])
			}
		}
	case map[string]any:
		for k := range (*value).(map[string]any) {
			v := (*value).(map[string]any)[k]
			prettify(&v)
			(*value).(map[string]any)[k] = v
		}
	}
}

func ParseRSC(rsc string) []RSCLine {
	lines := splitEachLine(rsc)
	for i := range lines {
		prettify(&lines[i].Value)
	}
	return lines
}
