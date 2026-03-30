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

	OperatorRelation   = "->"
	OperatorBiRelation = "<->"

	KeywordOpeningBracket = '['
	KeywordClosingBracket = ']'
)

type Token struct {
	Type     TokenType
	Value    string
	EndPos   int64
	StartPos int64
}

func (t *Token) String() string {
	return fmt.Sprintf("%d: %s, start_pos = %d, end_pos = %d\n", t.Type, t.Value, t.StartPos, t.EndPos)
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

	runes := []rune(queryString)
	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if !inToken {
			if c == ' ' {
				continue
			}
			inToken = true
			startPos = i
			curToken = ""
		}

		switch c {
		case ' ':
			t.addLiteralOrKeyword(&tokens, curToken, startPos, i-1)
			inToken = false

		case '-':
			t.addLiteralOrKeyword(&tokens, curToken, startPos, i-1)
			tokens = append(tokens, &Token{
				Type:     TokenTypeOperator,
				Value:    OperatorRelation,
				StartPos: int64(i),
				EndPos:   int64(i + 1),
			})
			i += 1
			inToken = false

		case '<':
			t.addLiteralOrKeyword(&tokens, curToken, startPos, i-1)
			tokens = append(tokens, &Token{
				Type:     TokenTypeOperator,
				Value:    OperatorBiRelation,
				StartPos: int64(i),
				EndPos:   int64(i + 2),
			})
			i += 2
			inToken = false

		case KeywordSeparator:
			t.addLiteralOrKeyword(&tokens, curToken, startPos, i-1)
			tokens = append(tokens, &Token{
				Type:     TokenTypeSeparator,
				Value:    string(KeywordSeparator),
				StartPos: int64(i),
				EndPos:   int64(i),
			})
			inToken = false

		case KeywordOpeningBracket, KeywordClosingBracket:
			t.addLiteralOrKeyword(&tokens, curToken, startPos, i-1)
			tokens = append(tokens, &Token{
				Type:     TokenTypeBracket,
				Value:    string(c),
				StartPos: int64(i),
				EndPos:   int64(i),
			})
			inToken = false

		default:
			curToken += string(c)
			if i == len(runes)-1 {
				t.addLiteralOrKeyword(&tokens, curToken, startPos, i)
			}
		}
	}

	return tokens, nil
}

// func (t *Tokenizer) Tokenize(queryString string) (tokens []*Token, err error) {
// 	tokens = make([]*Token, 0)

// 	curToken := ""
// 	inToken := false
// 	startPos := 0
// 	endPos := 0
// 	for i, c := range queryString {

// 		if inToken {

// 			switch c {
// 			case ' ':
// 				endPos = i - 1

// 				curToken = strings.TrimSpace(curToken)
// 				tokenType := TokenTypeLiteral
// 				if t.IsKeyword(curToken) {
// 					tokenType = TokenTypeKeyword
// 				}

// 				tokens = append(tokens, &Token{
// 					Type:     tokenType,
// 					Value:    curToken,
// 					StartPos: int64(startPos),
// 					EndPos:   int64(endPos),
// 				})
// 				continue

// 			case '-':

// 				endPos = i - 1

// 				curToken = strings.TrimSpace(curToken)
// 				tokenType := TokenTypeLiteral
// 				if t.IsKeyword(curToken) {
// 					tokenType = TokenTypeKeyword
// 				}

// 				tokens = append(tokens, &Token{
// 					Type:     tokenType,
// 					Value:    curToken,
// 					StartPos: int64(startPos),
// 					EndPos:   int64(endPos),
// 				})

// 				curToken = ""
// 				inToken = false

// 				// TODO: Add check if relation operation is defined correctly

// 				tokens = append(tokens, &Token{
// 					Type:     TokenTypeOperator,
// 					Value:    OperatorRelation,
// 					StartPos: int64(i),
// 					EndPos:   int64(i + 1),
// 				})

// 				i += 2

// 				continue

// 			case '<':

// 				endPos = i - 1

// 				curToken = strings.TrimSpace(curToken)
// 				tokenType := TokenTypeLiteral
// 				if t.IsKeyword(curToken) {
// 					tokenType = TokenTypeKeyword
// 				}

// 				tokens = append(tokens, &Token{
// 					Type:     tokenType,
// 					Value:    curToken,
// 					StartPos: int64(startPos),
// 					EndPos:   int64(endPos),
// 				})

// 				curToken = ""
// 				inToken = false

// 				tokens = append(tokens, &Token{
// 					Type:     TokenTypeOperator,
// 					Value:    OperatorBiRelation,
// 					EndPos:   int64(i + 2),
// 					StartPos: int64(i),
// 				})

// 				i += 3

// 				continue

// 			case KeywordSeparator:
// 				endPos = i - 1
// 				curToken = strings.TrimSpace(curToken)
// 				tokenType := TokenTypeLiteral
// 				if t.IsKeyword(curToken) {
// 					tokenType = TokenTypeKeyword
// 				}

// 				tokens = append(tokens, &Token{
// 					Type:     tokenType,
// 					Value:    curToken,
// 					StartPos: int64(startPos),
// 					EndPos:   int64(endPos),
// 				})

// 				curToken = ""
// 				inToken = false

// 				tokens = append(tokens, &Token{})

// 				continue

// 			case KeywordClosingBracket:
// 				continue
// 			}

// 		} else {

// 		}

// 		curToken += string(c)
// 	}

// 	return tokens, nil
// }

func (t *Tokenizer) IsKeyword(literal string) bool {
	switch literal {
	case KeywordAll, KeywordAnd, KeywordBetween, KeywordCheck, KeywordCount, KeywordData,
		KeywordDelete, KeywordHas, KeywordIn, KeywordInsert, KeywordNode, KeywordOut, KeywordPath, KeywordRelation, KeywordRelations,
		KeywordRestore, KeywordSelect, KeywordUpdate, KeywordWith:
		return true
	default:
		return false
	}
}

func (t *Tokenizer) addLiteralOrKeyword(tokens *[]*Token, val string, start, end int) {
	val = strings.TrimSpace(val)
	if val == "" {
		return
	}

	tokenType := TokenTypeLiteral
	if t.IsKeyword(val) {
		tokenType = TokenTypeKeyword
	}

	*tokens = append(*tokens, &Token{
		Type:     tokenType,
		Value:    val,
		StartPos: int64(start),
		EndPos:   int64(end),
	})
}
