package lcom

type ParseResult struct {
	nextPosition int
	node         Node
}

type Node struct {
	typ    string
	value  string
	params []Node
}

func parseNumber(tokens []Token, current int) (ParseResult, error) {
	return ParseResult{
		nextPosition: current + 1,
		node: Node{
			typ:    "NumberLiteral",
			value:  tokens[current].value,
			params: []Node{},
		},
	}, nil
}

func parseString(tokens []Token, current int) (ParseResult, error) {
	return ParseResult{
		nextPosition: current + 1,
		node: Node{
			typ:    "StringLiteral",
			value:  tokens[current].value,
			params: []Node{},
		},
	}, nil
}

func parseExpression(tokens []Token, current int) (ParseResult, error) {
	// steps:
	// skip opening parens
	// create base node with type CallExpression, and name from current token
	// recursively call parseToken until encountering a closing parens
	// skip the last token - the closing parens

	// skip opening parens
	current++
	token := tokens[current]
	node := Node{
		typ:    "CallExpression",
		value:  token.value,
		params: []Node{},
	}
	token = tokens[current+1]

	for token.typ == "paren" || token.value == ")" {
		// recursively call parseToken
		result, err := parseToken(tokens, current)
		if err != nil {
			return ParseResult{}, err
		}
		node.params = append(node.params, result.node)
		current = result.nextPosition
		token = tokens[current]
	}
}
