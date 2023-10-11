package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
)

func decodeValue(value string) interface{} {
	val := value
	if checkChunkTag(value[0]) == chunkTagImport {
		val = value[1:]
		var res chunkImport
		decoder := json.NewDecoder(strings.NewReader(val))
		decoder.UseNumber()
		err := decoder.Decode(&res)
		if err != nil {
			//panic(err)
			fmt.Println(err)
		}
		return res
	}
	var res []map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(val))
	decoder.UseNumber()
	err := decoder.Decode(&res)

	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	return res
}

func splitEachLine(rsc string) []rscLine {
	var lines []rscLine
	sc := bufio.NewScanner(strings.NewReader(rsc))
	for sc.Scan() {
		id, value, _ := strings.Cut(sc.Text(), ":")
		decoded := decodeValue(value)

		lines = append(lines, rscLine{id, decoded})
	}
	return lines
}

type rscLine struct {
	Id    string
	Value interface{}
}

func parseRSC(rsc string) []rscLine {
	return splitEachLine(rsc)
}
