package query

import (
	"fmt"
	"testing"

	"modalsdb.com/graph_module/src/query"
)

func TestInsertQueryParser(t *testing.T) {

	insertNodeQuery := "INSERT john;"
	insertOneRelationQuery := "INSERT john->[:works_for]->google"
	insertBidirectional := "INSERT pizza->[:has,:on_top]<-cheese"

	parser := query.NewParser(nil)

	r1, err := parser.Parse(insertNodeQuery)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("result 1:", r1)

	r2, err := parser.Parse(insertOneRelationQuery)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("result 2:", r2)

	r3, err := parser.Parse(insertBidirectional)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("result 3:", r3)

}
