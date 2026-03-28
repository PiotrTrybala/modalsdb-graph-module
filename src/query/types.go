package query

import (
	"errors"
	"fmt"
	"strings"
)

type QueryType int

const (
	Select QueryType = iota
	Insert
	Delete
	Check
	Update
	Count
	None
)

var InvalidType = errors.New("invalid query type")

func ToQueryType(_type string) (queryType QueryType, err error) {

	_type = strings.ToLower(_type)

	switch _type {
	case "select":
		return Select, nil
	case "insert":
		return Insert, nil
	case "delete":
		return Delete, nil
	case "update":
		return Update, nil
	case "check":
		return Check, nil
	case "count":
		return Count, nil
	default:
		return None, fmt.Errorf("%s: %s", InvalidType, _type)
	}

}

type Query struct {
	Type QueryType
	Data any
}

type DirectionType int

const (
	DirectionIn DirectionType = iota
	DirectionOut
	DirectionAll
	DirectionNone
)

var InvalidDirectionType = errors.New("invalid direction type")

func ToDirectionType(direction string) (directionType DirectionType, err error) {
	direction = strings.ToLower(direction)

	switch direction {
	case "in":
		return DirectionIn, nil
	case "out":
		return DirectionOut, nil
	case "all":
		return DirectionAll, nil
	default:
		return DirectionNone, fmt.Errorf("%s: %s", InvalidDirectionType, direction)
	}
}

type SelectQuery struct {
	Direction DirectionType
	NodesIds  []string
}

func (q SelectQuery) String() string {
	return fmt.Sprintf(`direction = %d,nodesIds = %v
	`, q.Direction, q.NodesIds)
}

var ErrInvalidKeyword = errors.New("invalid keyword")

const (
	KeywordRelations                 = "relations"
	KeywordRelation                  = "relation"
	KeywordRelationship              = "->"
	KeywordBidirectionalRelationship = "<->"
)

type InsertNodeQuery struct {
	Name string
}

type InsertRelationQuery struct {
	From      string
	Relations []string // max two relations from->to and to->from
	To        string
}
