package main

import (
	"fmt"

	"modalsdb.com/graph_module/src/query"
)

func main() {
	q := "SELECT RELATIONS IN [john, alice];"

	tokenizer := query.NewTokenizer()

	parsed, err := tokenizer.Tokenize(q)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("parsed:", parsed)

}
