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

func FromType(_type string) (queryType QueryType, err error) {

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
