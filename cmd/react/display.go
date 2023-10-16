package react

import (
	"fmt"
	"github.com/fatih/color"
)

func DisplayRSC(lines []RSCLine) {
	indent := "  "
	for _, l := range lines {
		switch l.Value.(type) {
		case []interface{}:
			color.Yellow("%s", l.Id)
			displayValue(l.Value, indent)
		case chunkImport:
			color.Magenta("%s Import Chunk", l.Id)
			displayValue(l.Value, indent)
		default:
			color.Green("%s\n", l.Id)
			displayValue(l.Value, indent)
		}
	}
}

func displayValue(value any, indent string) {
	switch value.(type) {
	case []interface{}:
		for _, v := range value.([]interface{}) {
			displayValue(v, indent+indent)
		}
	case JSXElement:
		color.Cyan("%v", formatJSX(value.(JSXElement), indent))
	case chunkImport:
		for _, v := range value.(chunkImport).Chunks {
			color.Magenta("%s%v", indent, v)
		}
	default:
		color.Green("%s%v", indent, value)
	}
}

func formatJSX(jsx JSXElement, indent string) string {
	formatted := ""
	if jsx.Key == "" {
		formatted = fmt.Sprintf("%s<%s", indent, jsx.Type)
	} else {
		formatted = fmt.Sprintf("%s<%s key=%s", indent, jsx.Type, jsx.Key)
	}
	for k, v := range jsx.Props {
		if k == "children" {
			continue
		} else {
			formatted += fmt.Sprintf(" %s=%v", k, v)
		}
	}
	formatted += fmt.Sprintf(">")
	children := jsx.Props["children"]
	if children != nil {
		childIndent := "  "
		switch children.(type) {
		case []interface{}:
			for _, child := range children.([]interface{}) {
				switch child.(type) {
				case JSXElement:
					formatted += fmt.Sprintf("\n%v\n", formatJSX(child.(JSXElement), childIndent+indent))
				default:
					formatted += fmt.Sprintf("\n%s%v\n", childIndent+indent, child)
				}
			}
		default:
			formatted += fmt.Sprintf("\n%s%v\n", indent+childIndent, children)
		}
	}
	formatted += fmt.Sprintf("%s</%s>", indent, jsx.Type)

	return formatted
}
