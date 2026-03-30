package query

type ParserOptions struct{}

func DefaultParserOptions() *ParserOptions {
	return &ParserOptions{}
}

type Parser struct {
	tokenizer Tokenizer
}

func NewParser() *Parser {
	return &Parser{
		tokenizer: *NewTokenizer(),
	}
}

func (p *Parser) Parse(queryString string) (query *Query, err error) {

	tokens, err := p.tokenizer.Tokenize(queryString)
	if err != nil {
		return nil, err
	}

	queryType, err := ToQueryType(tokens[0].Value)
	if err != nil {
		return nil, err
	}

	switch queryType {
	case Select:
		selectQuery, err := p.ParseSelectQuery(tokens)
		if err != nil {
			return nil, err
		}

		return &Query{
			Type:     Select,
			Metadata: selectQuery,
		}, nil
	case Insert:
		insertQuery, err := p.ParseInsertQuery(tokens)
		if err != nil {
			return nil, err
		}

		return &Query{
			Type:     Insert,
			Metadata: insertQuery,
		}, nil
	case Update:
		updateQuery, err := p.ParseUpdateQuery(tokens)
		if err != nil {
			return nil, err
		}

		return &Query{
			Type:     Insert,
			Metadata: updateQuery,
		}, nil
	case Delete:
		deleteQuery, err := p.ParseDeleteQuery(tokens)
		if err != nil {
			return nil, err
		}

		return &Query{
			Type:     Insert,
			Metadata: deleteQuery,
		}, nil
	case Check:
		checkQuery, err := p.ParseCheckQuery(tokens)
		if err != nil {
			return nil, err
		}

		return &Query{
			Type:     Insert,
			Metadata: checkQuery,
		}, nil
	case Count:
		countQuery, err := p.ParseCountQuery(tokens)
		if err != nil {
			return nil, err
		}

		return &Query{
			Type:     Insert,
			Metadata: countQuery,
		}, nil
	}

	return query, nil
}

func (p *Parser) ParseSelectQuery(tokens []*Token) (query *SelectQuery, err error) {
	return query, nil
}

func (p *Parser) ParseInsertQuery(tokens []*Token) (query *InsertQuery, err error) {
	return query, nil
}

func (p *Parser) ParseUpdateQuery(tokens []*Token) (query *UpdateQuery, err error) {
	return query, nil
}

func (p *Parser) ParseDeleteQuery(tokens []*Token) (query *DeleteQuery, err error) {
	return query, nil
}

func (p *Parser) ParseCountQuery(tokens []*Token) (query *CountQuery, err error) {
	return query, nil
}

func (p *Parser) ParseCheckQuery(tokens []*Token) (query *CheckQuery, err error) {
	return query, nil
}
