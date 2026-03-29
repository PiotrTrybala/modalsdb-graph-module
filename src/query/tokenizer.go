package query

const (
	TokenTerminator = ";"
	TokenSeparator  = ","
	TokenOpenList   = "["
	TokenCloseList  = "]"
)

type Tokenizer struct{}

func (t *Tokenizer) Tokenize(queryStrin string) (tokens []string, err error) {
	return []string{}, nil
}
