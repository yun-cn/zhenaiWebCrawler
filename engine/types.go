package engine

//Request Types
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult  是上面的返回类型
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser return nil
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
