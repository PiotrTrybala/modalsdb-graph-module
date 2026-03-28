package query

import (
	"errors"
	"fmt"
	"strings"
)

type ParserOptions struct{}

type Parser struct{}

func NewParser(opt *ParserOptions) *Parser {
	return &Parser{}
}

func (parser *Parser) Parse(queryString string) (query *Query, err error) {

	fmt.Println("query:", queryString)

	// tokenize input
	tokens := strings.Split(queryString, " ")

	fmt.Println("tokens:", tokens)

	// determine query type

	queryType, err := ToQueryType(tokens[0])
	if err != nil {
		return nil, err
	}

	fmt.Println("query type:", queryType)
	switch queryType {
	case Select:

		data, err := parser.ParseSelectQuery(tokens[])

	default:
		return nil, errors.ErrUnsupported
	}
}

func (parser *Parser) ParseSelectQuery(tokens []string) (data *SelectQuery, err error) {

	return &SelectQuery{}, nil
}
