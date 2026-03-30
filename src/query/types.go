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
	Type     QueryType
	Metadata any
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
	Ids       []string
}

type InsertQueryType int

const (
	InsertQueryNode InsertQueryType = iota
	InsertQueryRelation
)

type InsertQuery struct {
	Type InsertQueryType

	NodeId string

	RelationFrom string
	Relations    []string
	RelationTo   string
}

type UpdateQueryType int

const (
	UpdateQueryData UpdateQueryType = iota
	UpdateQueryRelation
)

type UpdateQueryRelationType int

const (
	UpdateQueryRelationOneWay UpdateQueryRelationType = iota
	UpdateQueryRelationTwoWay
)

type UpdateQuery struct {
	Type UpdateQueryType

	NodeData string

	RelationType UpdateQueryRelationType
	RelationFrom string
	RelationTo   string
	RelationData string
}

type DeleteQueryType int

const (
	DeleteQueryRelation DeleteQueryType = iota
	DeleteQueryNode
)

type DeleteQuery struct {
	Type DeleteQueryType

	NodeId string

	RelationFrom string
	RelationTo   string
}

type CheckQueryType int

const (
	CheckQueryRelation CheckQueryType = iota
	CheckQueryPath
	CheckQueryContainRelation
)

type CheckQuery struct {
	Type CheckQueryType

	Source      string
	Destination string

	Relations []string
}

type CountQueryType int

const (
	CountQueryNodes CountQueryType = iota
)

type CountQuery struct {
	Direction DirectionType
	NodeId    string
}
