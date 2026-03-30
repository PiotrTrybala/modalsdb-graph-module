package main

import (
	"fmt"

	"modalsdb.com/graph_module/src/query"
)

func main() {
	queries := []string{
		"SELECT RELATIONS IN [john, alice];",
		"INSERT pizza<->[:has,:on_top]<->cheese;",
	}

	tokenizer := query.NewTokenizer()

	for _, query := range queries {
		parsed, err := tokenizer.Tokenize(query)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("parsed:", query, "\n", parsed)
	}

}
