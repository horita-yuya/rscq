package react

import (
	"fmt"
	"github.com/fatih/color"
)

func DisplayRSC(lines []RSCLine) {
	indent := "  "
	sep := "--------------------"
	for _, l := range lines {
		if l.Id == "0" {
			color.Red("%s", l.Value)
		}
		switch l.Value.(type) {
		case []interface{}:
			color.Magenta("%s%s", l.Id, sep)
			displayValue(l.Value, indent)
			color.Magenta("%s%s", sep, l.Id)
		case chunkImport:
			color.Magenta("%s%s", l.Id, sep)
			displayValue(l.Value, indent)
			color.Magenta("%s%s", sep, l.Id)
		default:
			color.Magenta("%s%s\n", l.Id, sep)
			displayValue(l.Value, indent)
			color.Magenta("%s%s\n", sep, l.Id)
		}
	}
}

func displayValue(value any, indent string) {
	switch value.(type) {
	case []interface{}:
		for _, v := range value.([]interface{}) {
			displayValue(v, indent)
		}
	case JSXElement:
		color.Magenta("%v", formatJSX(value.(JSXElement), indent))
	case chunkImport:
		for _, v := range value.(chunkImport).Chunks {
			color.Magenta("%s%v", indent, v)
		}
	default:
		color.Magenta("%s%v", indent, value)
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
		case JSXElement:
			formatted += fmt.Sprintf("\n%v\n", formatJSX(children.(JSXElement), childIndent+indent))
		default:
			formatted += fmt.Sprintf("\n%s%v\n", indent+childIndent, children)
		}
	}
	formatted += fmt.Sprintf("%s</%s>", indent, jsx.Type)

	return formatted
}
