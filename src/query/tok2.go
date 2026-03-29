package query

// import (
// 	"errors"
// 	"strings"
// )

// const (
// 	TokenTerminator = ";"
// 	TokenSeparator  = ","

// 	TokenOpenList  = "["
// 	TokenCloseList = "]"

// 	TokenDirection     = "->"
// 	TokenBidirectional = "<->"

// 	TokenRelationDefinition = ":"
// )

// var ErrInvalidKeyword = errors.New("invalid keyword")

// const (
// 	KeywordIn  = "in"
// 	KeywordOut = "out"
// 	KeywordAll = "all"

// 	KeywordRelations = "relations"
// 	KeywordRelation  = "relation"

// 	KeywordSelect  = "select"
// 	KeywordInsert  = "insert"
// 	KeywordUpdate  = "update"
// 	KeywordDelete  = "delete"
// 	KeywordCheck   = "check"
// 	KeywordCount   = "count"
// 	KeywordRestore = "restore"

// 	KeywordPath    = "path"
// 	KeywordBetween = "between"
// 	KeywordAnd     = "and"
// 	KeywordHas     = "has"
// 	KeywordData    = "data"
// 	KeywordNode    = "node"
// 	KeywordWith    = "with"
// )

// type TokenType int

// const (
// 	TokenTypeVariable TokenType = iota
// 	TokenTypeKeyword
// 	TokenTypeList
// 	TokenTypeData
// 	TokenTypeOther
// )

// type Token struct {
// 	Type  TokenType
// 	Value string
// }

// type Tokenizer struct{}

// func (t *Tokenizer) Tokenize(queryString string) (tokens []*Token, err error) {
// 	tokens = make([]*Token, 0)
// 	rawTokens := strings.Split(queryString, " ")

// 	for idx, token := range rawTokens {

// 		if keyword, err := t.toKeywordToken(token); err == nil {
// 			tokens = append(tokens, keyword)
// 		}

// 		if

// 	}

// 	return []string{}, nil
// }

// func (t *Tokenizer) toKeywordToken(rawToken string) (keywordToken *Token, err error) {
// 	return keywordToken, nil
// }

// func (t *Tokenizer) isDataToken(rawToken string) bool {
// 	return false
// }

// func (t *Tokenizer) isListToken(rawToken string) bool {
// 	return false
// }

// func (t *Tokenizer) toOtherToken(rawToken string) (otherToken *Token, err error) {
// 	return otherToken, nil
// }
