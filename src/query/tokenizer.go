package query

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidToken = errors.New("invalid token type")

type TokenType int

const (
	TokenTypeLiteral TokenType = iota
	TokenTypeKeyword
	TokenTypeSeparator
	TokenTypeBracket
	TokenTypeSemicolon
	TokenTypeOperator
)

const (
	KeywordIn  = "in"
	KeywordOut = "out"
	KeywordAll = "all"

	KeywordRelations = "relations"
	KeywordRelation  = "relation"

	KeywordSelect  = "select"
	KeywordInsert  = "insert"
	KeywordUpdate  = "update"
	KeywordDelete  = "delete"
	KeywordCheck   = "check"
	KeywordCount   = "count"
	KeywordRestore = "restore"

	KeywordPath    = "path"
	KeywordBetween = "between"
	KeywordAnd     = "and"
	KeywordHas     = "has"
	KeywordData    = "data"
	KeywordNode    = "node"
	KeywordWith    = "with"

	KeywordSemicolon = ';'

	KeywordSeparator = ','

	KeywordRelationship = ':'

	KeywordOperatorRelation   = "->"
	KeywordOperatorBiRelation = "<->"

	KeywordOpeningBracket = '['
	KeywordClosingBracket = ']'
)

type Token struct {
	Type     TokenType
	Value    string
	EndPos   int64
	StartPos int64
}

type Tokenizer struct {
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) Tokenize(queryString string) (tokens []*Token, err error) {
	tokens = make([]*Token, 0)

	curToken := ""
	inToken := false
	startPos := 0
	endPos := 0
	for i, c := range queryString {

		if inToken {

			if c == ' ' {
				endPos = i - 1
				fmt.Printf("literal = %s, start_pos = %d, end_pos = %d\n", strings.TrimSpace(curToken), startPos, endPos)

				inToken = false
				curToken = ""
				continue
			}

			if c == KeywordSeparator {
				endPos = i - 1
				fmt.Printf("literal = %s, start_pos = %d, end_pos = %d\n", curToken, startPos, endPos)
				inToken = false
				curToken = ""
				fmt.Printf("separator = %s, start_pos = %d, end_pos = %d\n", string(c), i, i)
				continue
			}

			if c == KeywordClosingBracket {
				endPos = i - 1
				fmt.Printf("literal = %s, start_pos = %d, end_pos = %d\n", curToken, startPos, endPos)

				inToken = false
				curToken = ""

				fmt.Printf("closing bracket = %s, start_pos = %d, end_pos = %d\n", curToken, startPos, endPos)

				continue
			}

		} else {

			if c == ' ' {
				continue
			}

			if c == KeywordOpeningBracket {
				fmt.Printf("bracket opening =  %s, start_pos = %d, end_pos = %d\n", string(c), i, i)
				continue
			}

			if c == KeywordSeparator {
				fmt.Printf("separator = %s, start_pos = %d, end_pos = %d\n", string(c), i, i)
				continue
			}

			if c == KeywordSemicolon {
				fmt.Printf("semicolon = %s, start_pos = %d, end_pos = %d\n", string(c), i, i)
				continue
			}

			if c == KeywordClosingBracket {
				fmt.Printf("bracket closing = %s, start_pos = %d, end_pos = %d\n", string(c), i, i)
				continue
			}

			inToken = true
			startPos = i
		}

		curToken += string(c)
	}

	return tokens, nil
}
