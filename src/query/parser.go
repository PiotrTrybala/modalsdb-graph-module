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

		data, err := parser.ParseSelectQuery(tokens)
		if err != nil {
			return nil, err
		}

		return &Query{
			Type: queryType,
			Data: data,
		}, nil

	default:
		return nil, errors.ErrUnsupported
	}
}

func (parser *Parser) ParseSelectQuery(tokens []string) (data *SelectQuery, err error) {

	relationsKeyword := strings.ToLower(tokens[1])
	if relationsKeyword != KeywordRelations {
		return nil, fmt.Errorf("%s: %s", ErrInvalidKeyword, relationsKeyword)
	}

	directionType, err := ToDirectionType(tokens[2])
	if err != nil {
		return nil, err
	}

	// get all tokens belonging to nodes ids
	var sb strings.Builder
	for i := 3; i < len(tokens); i++ {
		cleaned := strings.ToLower(tokens[i])
		cleaned = strings.TrimSpace(cleaned)
		sb.WriteString(cleaned)
	}

	nodesIdsString := sb.String()

	nodesIds, err := parser.ParseNodesIdsList(nodesIdsString)
	if err != nil {
		return nil, err
	}

	return &SelectQuery{
		Direction: directionType,
		NodesIds:  nodesIds,
	}, nil
}

func (parser *Parser) ParseNodesIdsList(list string) (nodesIds []string, err error) {

	if list[0] != '[' {
		return nil, errors.New("invalid nodes ids list: no opening bracket")
	}

	// naive nodes ids list parsing

	if list[len(list)-1] != ';' {
		return nil, errors.New("invalid nodes ids list: no semicolon at the end")
	}

	if list[len(list)-2] != ']' { // by because of ;
		return nil, errors.New("invalid nodes ids list: no closing bracket")
	}
	trimmed := strings.Trim(list, "[];")
	nodesIds = strings.Split(trimmed, ",")
	return nodesIds, nil
}
