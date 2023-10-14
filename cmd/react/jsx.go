package react

import "fmt"

//https://github.com/facebook/react/blob/main/packages/react-server/src/ReactFlightServer.js#L511

type JSXElement struct {
	Type  string
	Key   string
	Props map[string]any
}

func (jsx JSXElement) String() string {
	formatted := ""
	if jsx.Key == "" {
		formatted = fmt.Sprintf("<%s", jsx.Type)
	} else {
		formatted = fmt.Sprintf("<%s key=%s", jsx.Type, jsx.Key)
	}
	for k, v := range jsx.Props {
		if k == "children" {
			continue
		} else {
			formatted += fmt.Sprintf(" %s=%v", k, v)
		}
	}
	formatted += ">"
	children := jsx.Props["children"]
	if children != nil {
		formatted += fmt.Sprintf("\n    %v\n", children)
	}
	formatted += fmt.Sprintf("</%s>", jsx.Type)
	return formatted
}

func ParseJsx(elements []interface{}) JSXElement {
	t := elements[1].(string)
	ps := elements[3].(map[string]interface{})

	var k string
	if elements[2] == nil {
		k = ""
	} else {
		k = elements[2].(string)
	}

	return JSXElement{
		t,
		k,
		ps,
	}
}
