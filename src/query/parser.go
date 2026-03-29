package query

// import (
// 	"errors"
// 	"fmt"
// 	"slices"
// 	"strings"
// )

// type ParserOptions struct{}

// type Parser struct{}

// func NewParser(opt *ParserOptions) *Parser {
// 	return &Parser{}
// }

// func (parser *Parser) Parse(queryString string) (query *Query, err error) {

// 	fmt.Println("query:", queryString)

// 	// tokenize input
// 	tokens := strings.Split(queryString, " ")
// 	for i := 0; i < len(tokens); i++ {
// 		tokens[i] = strings.ToLower(tokens[i])
// 	}

// 	fmt.Println("tokens:", tokens)

// 	// determine query type

// 	queryType, err := ToQueryType(tokens[0])
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("query type:", queryType)
// 	switch queryType {
// 	case Select:

// 		data, err := parser.ParseSelectQuery(tokens)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return &Query{
// 			Type: queryType,
// 			Data: data,
// 		}, nil

// 	case Insert:
// 		data, err := parser.ParseInsertQuery(tokens)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &Query{
// 			Type: queryType,
// 			Data: data,
// 		}, nil

// 	case Count:
// 		data, err := parser.ParseCountQuery(tokens)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &Query{
// 			Type: queryType,
// 			Data: data,
// 		}, nil

// 	case Check:
// 		data, err := parser.ParseCheckQuery(tokens)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return &Query{
// 			Type: queryType,
// 			Data: data,
// 		}, nil
// 	default:
// 		return nil, errors.ErrUnsupported
// 	}
// }

// func (parser *Parser) ParseSelectQuery(tokens []string) (data *SelectQuery, err error) {

// 	relationsKeyword := strings.ToLower(tokens[1])
// 	if relationsKeyword != KeywordRelations {
// 		return nil, fmt.Errorf("%s: %s", ErrInvalidKeyword, relationsKeyword)
// 	}

// 	directionType, err := ToDirectionType(tokens[2])
// 	if err != nil {
// 		return nil, err
// 	}

// 	// get all tokens belonging to nodes ids
// 	var sb strings.Builder
// 	for i := 3; i < len(tokens); i++ {
// 		cleaned := strings.ToLower(tokens[i])
// 		cleaned = strings.TrimSpace(cleaned)
// 		sb.WriteString(cleaned)
// 	}

// 	nodesIdsString := sb.String()

// 	nodesIds, err := parser.ParseNodesIdsList(nodesIdsString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &SelectQuery{
// 		Direction: directionType,
// 		NodesIds:  nodesIds,
// 	}, nil
// }

// func (parser *Parser) ParseNodesIdsList(list string) (nodesIds []string, err error) {

// 	if list[0] != '[' {
// 		return nil, errors.New("invalid nodes ids list: no opening bracket")
// 	}

// 	// naive nodes ids list parsing

// 	if list[len(list)-1] != ';' {
// 		return nil, errors.New("invalid nodes ids list: no semicolon at the end")
// 	}

// 	if list[len(list)-2] != ']' { // by because of ;
// 		return nil, errors.New("invalid nodes ids list: no closing bracket")
// 	}
// 	trimmed := strings.Trim(list, "[];")
// 	nodesIds = strings.Split(trimmed, ",")
// 	return nodesIds, nil
// }

// func (parser *Parser) ParseInsertQuery(tokens []string) (data any, err error) {
// 	if len(tokens) == 2 {
// 		return parser.ParseInsertNodeQuery(tokens)
// 	}
// 	return parser.ParseInsertRelationQuery(tokens)
// }

// func (parser *Parser) ParseInsertNodeQuery(tokens []string) (query *InsertNodeQuery, err error) {

// 	name := strings.TrimSpace(tokens[1])
// 	cleaned := strings.ToLower(name)

// 	return &InsertNodeQuery{
// 		Name: cleaned,
// 	}, nil
// }

// func (parser *Parser) ParseInsertRelationQuery(tokens []string) (query *InsertRelationQuery, err error) {

// 	relationKeyword := strings.TrimSpace(tokens[1])
// 	if relationKeyword != KeywordRelation {
// 		return nil, fmt.Errorf("%s: %s", ErrInvalidKeyword, relationKeyword)
// 	}

// 	relation := strings.TrimSpace(tokens[2])

// 	hasRelationship := strings.Contains(relation, KeywordRelationship)
// 	hasBidirectionalRelationship := strings.Contains(relation, KeywordBidirectionalRelationship)

// 	components := []string{}
// 	if hasBidirectionalRelationship {
// 		components = strings.Split(relation, KeywordBidirectionalRelationship)
// 	} else if hasRelationship {
// 		components = strings.Split(relation, KeywordRelationship)
// 	} else if hasBidirectionalRelationship && hasRelationship {
// 		return nil, errors.New("insert relation query can only have one type of relation")
// 	} else {
// 		return nil, errors.New("none of relationship operators were found")
// 	}

// 	if len(components) > 3 {
// 		return nil, errors.New("invalid relationship definition")
// 	}

// 	from := strings.ToLower(components[0])
// 	relations, err := parser.ParserRelations(components[1])
// 	if err != nil {
// 		return nil, err
// 	}
// 	to := strings.ToLower(components[2])

// 	return &InsertRelationQuery{
// 		From:      from,
// 		Relations: relations,
// 		To:        to,
// 	}, nil
// }

// func (parser *Parser) ParserRelations(rawString string) (relations []string, err error) {

// 	if rawString[0] != '[' {
// 		return nil, errors.New("invalid opening of relationship rule")
// 	}

// 	if rawString[len(rawString)-2] != ']' {
// 		return nil, errors.New("invalid closing of relationship rule")
// 	}

// 	if rawString[len(rawString)-1] != ';' {
// 		return nil, errors.New("invalid termination character")
// 	}

// 	cleaned := strings.Trim(rawString, "[];")
// 	cleaned = strings.ToLower(cleaned)
// 	names := strings.Split(cleaned, ",")
// 	if len(names) > 2 {
// 		return nil, errors.New("invalid number of relationships specified")
// 	}

// 	for _, name := range names {
// 		relations = append(relations, strings.Trim(name, ":"))
// 	}

// 	return relations, nil
// }

// func (parser *Parser) ParseCountQuery(tokens []string) (query *CountQuery, err error) {

// 	if len(tokens) != 3 {
// 		return nil, errors.New("invalid count query: not enough parameters")
// 	}

// 	direction, err := ToDirectionType(tokens[1])
// 	if err != nil {
// 		return nil, err
// 	}

// 	name := strings.ToLower(tokens[2])
// 	name = strings.TrimSpace(name)

// 	return &CountQuery{
// 		Direction: direction,
// 		Name:      name,
// 	}, nil
// }

// func (parser *Parser) ParseCheckQuery(tokens []string) (query *CheckQuery, err error) {

// 	hasAnd := slices.Contains(tokens, KeywordAnd)
// 	hasBetween := slices.Contains(tokens, KeywordBetween)

// 	if hasAnd {
// 		if tokens[2] != KeywordAnd {
// 			return nil, errors.New("invalid query format")
// 		}

// 		source := tokens[1]
// 		relations := []string{}
// 		destination := tokens[3]

// 		if len(tokens) > 4 && tokens[4] == KeywordHasRelations {
// 			relations, err = parser.ParserRelations(tokens[5])
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		return &CheckQuery{
// 			Destination:  destination,
// 			Source:       source,
// 			Action:       CheckActionAnd,
// 			HasRelations: relations,
// 		}, nil
// 	} else if hasBetween {

// 		if tokens[2] != KeywordBetween {
// 			return nil, errors.New("invalid query format")
// 		}

// 		source := tokens[1]
// 		destination := tokens[3]

// 		return &CheckQuery{
// 			Source:       source,
// 			Destination:  destination,
// 			Action:       CheckActionBetween,
// 			HasRelations: []string{},
// 		}, nil

// 	}

// 	fmt.Println(tokens)

// 	return nil, errors.New("invalid check query")
// }
