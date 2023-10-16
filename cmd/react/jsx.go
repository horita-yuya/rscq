package react

//https://github.com/facebook/react/blob/main/packages/react-server/src/ReactFlightServer.js#L511

type JSXElement struct {
	Type  string
	Key   string
	Props map[string]any
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
