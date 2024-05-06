package lcom

import (
	"fmt"
	"regexp"
)

type MToken struct {
	consumedChars int
	token         Token
}

type Token struct {
	typ   string
	value string
}

func (t Token) IsNil() bool {
	return t.typ == "" && t.value == ""
}

func (t Token) String() string {
	// Assuming a maximum length of 10 for `typ`. Adjust the width as needed.
	return fmt.Sprintf("type: %-10s, value: %s", t.typ, t.value)
}

func itsWhitespace(input string, current int) (MToken, error) {
	char := string(input[current])

	re, err := regexp.Compile(`\s`)
	if err != nil {
		return MToken{0, Token{}}, fmt.Errorf("failed to compile regex pattern: %s", err)
	}

	if re.MatchString(char) {
		return MToken{consumedChars: 1}, nil
	}

	return MToken{consumedChars: 0}, nil
}

func tokenizeCharacter(typ string, value string, input string, current int) (MToken, error) {
	// check if beyond the end of the input
	if current >= len(input) {
		return MToken{0, Token{}}, nil
	}

	// check if the character matches
	if string(input[current]) == value {
		return MToken{1, Token{typ, value}}, nil
	}

	// no match
	return MToken{0, Token{}}, nil
}

func tokenizeParenOpen(input string, current int) (MToken, error) {
	return tokenizeCharacter("paren", "(", input, current)
}

func tokenizeParenClose(input string, current int) (MToken, error) {
	return tokenizeCharacter("paren", ")", input, current)
}

func tokenizePattern(typ string, pattern string, input string, current int) (MToken, error) {
	// Check if beyond the end of the input
	if current >= len(input) {
		return MToken{0, Token{}}, nil
	}

	// Compile the regex pattern once outside the loop
	re, err := regexp.Compile(pattern)
	if err != nil {
		return MToken{0, Token{}}, fmt.Errorf("failed to compile regex pattern: %s", err)
	}

	// Keep eating characters until we find a non-match
	consumedChars := 0
	value := ""
	for current+consumedChars < len(input) && re.MatchString(string(input[current+consumedChars])) {
		value += string(input[current+consumedChars])
		consumedChars++
	}

	// Return the token if we consumed at least one character
	if consumedChars > 0 {
		return MToken{consumedChars, Token{typ, value}}, nil
	}

	return MToken{0, Token{}}, nil
}

func tokenizeNumber(input string, current int) (MToken, error) {
	return tokenizePattern("number", "[0-9]", input, current)
}

func tokenizeName(input string, current int) (MToken, error) {
	return tokenizePattern("name", "[a-z]", input, current)
}

func tokenizeString(input string, current int) (MToken, error) {
	// Check if beyond the end of the input
	if current >= len(input) {
		return MToken{0, Token{}}, nil
	}

	// Fail if the first character is not a quote
	if input[current] != '"' {
		return MToken{0, Token{}}, nil
	}

	// Process characters until the closing quote or end of string
	value := ""
	consumedChars := 1 // Start after the opening quote
	for current+consumedChars < len(input) && input[current+consumedChars] != '"' {
		value += string(input[current+consumedChars])
		consumedChars++
	}

	// Check if the loop ended without finding a closing quote
	if current+consumedChars >= len(input) {
		return MToken{0, Token{}}, fmt.Errorf("unterminated string at position %d", current)
	}

	// Include the closing quote in consumed characters
	return MToken{consumedChars + 1, Token{"string", value}}, nil
}

func Tokenize(input string) ([]Token, error) {
	tokenizers := []func(string, int) (MToken, error){
		itsWhitespace,
		tokenizeParenOpen,
		tokenizeParenClose,
		tokenizeNumber,
		tokenizeName,
		tokenizeString,
	}

	current := 0
	tokens := []Token{}
	for current < len(input) {
		found := false
		for _, tokenizer := range tokenizers {
			mtoken, err := tokenizer(input, current)
			if err != nil {
				return nil, err
			}
			if mtoken.consumedChars > 0 {
				current += mtoken.consumedChars

				// check if token is null
				if mtoken.token.IsNil() {
					continue
				}

				found = true
				tokens = append(tokens, mtoken.token)
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("no tokenizer found for input: %s", input[current:])
		}
	}

	return tokens, nil
}
