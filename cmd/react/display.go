package react

import (
	"fmt"
	"github.com/fatih/color"
)

func DisplayRSC(lines []RSCLine) {
	indent := "  "
	sep := "--------------------"
	for _, l := range lines {
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

	propsIndent := "  "
	jsxIndent := "  "
	jsxCount := 0
	childrenCount := 0
	for k, v := range jsx.Props {
		if k == "children" {
			childrenCount += 1
			continue
		} else {
			switch v.(type) {
			case JSXElement:
				jsxCount += 1
				formatted += fmt.Sprintf("\n%s%s={\n%v\n%s}\n", propsIndent+indent, k, formatJSX(v.(JSXElement), jsxIndent+propsIndent+indent), propsIndent+indent)
			default:
				formatted += fmt.Sprintf("\n%s%s={%v}", propsIndent+indent, k, v)
			}
		}
	}

	if len(jsx.Props) == 0 || len(jsx.Props) == 1 && childrenCount == 1 {
		formatted += fmt.Sprintf(">")
	} else if len(jsx.Props) == jsxCount {
		formatted += fmt.Sprintf("\n%s>", propsIndent+indent)
	} else {
		formatted += fmt.Sprintf("\n%s>", indent)
	}

	children := jsx.Props["children"]
	if children != nil {
		switch children.(type) {
		case []interface{}:
			for _, child := range children.([]interface{}) {
				switch child.(type) {
				case JSXElement:
					formatted += fmt.Sprintf("\n%v\n", formatJSX(child.(JSXElement), jsxIndent+indent))
				default:
					formatted += fmt.Sprintf("\n%s%v\n", jsxIndent+indent, child)
				}
			}
		case JSXElement:
			formatted += fmt.Sprintf("\n%v\n", formatJSX(children.(JSXElement), jsxIndent+indent))
		default:
			formatted += fmt.Sprintf("\n%s%v\n", indent+jsxIndent, children)
		}
	} else {
		// align heads of closing tags.
		// <div>
		// </div>
		formatted += "\n"
	}
	formatted += fmt.Sprintf("%s</%s>", indent, jsx.Type)

	return formatted
}
