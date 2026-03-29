package query

// import (
// 	"fmt"
// 	"testing"

// 	"modalsdb.com/graph_module/src/query"
// )

// func TestSelectQueryParser(t *testing.T) {

// 	selectAllRelations := "SELECT RELATIONS IN [john, alice];"
// 	selectInRelations := "SELECT RELATIONS OUT [john, alice];"
// 	selectOutRelations := "SELECT RELATIONS ALL [john, alice];"

// 	parser := query.NewParser(nil)

// 	r1, err := parser.Parse(selectAllRelations)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	fmt.Println("result 1:", r1)

// 	r2, err := parser.Parse(selectInRelations)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	fmt.Println("result 2:", r2)

// 	r3, err := parser.Parse(selectOutRelations)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println("result 3:", r3)

// }
