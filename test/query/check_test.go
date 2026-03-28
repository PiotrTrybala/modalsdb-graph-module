package query

import (
	"fmt"
	"testing"

	"modalsdb.com/graph_module/src/query"
)

func TestCheckQueryParser(t *testing.T) {

	checkQuery1 := "CHECK john AND alice;"
	checkQuery2 := "CHECK john BETWEEN alice;"
	checkQuery3 := "CHECK john AND alice HAS [:works_on,:on_top_of];"

	parser := query.NewParser(nil)

	r1, err := parser.Parse(checkQuery1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r1)

	r2, err := parser.Parse(checkQuery2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r2)

	r3, err := parser.Parse(checkQuery3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r3)

}
