package engine

type Request struct {
	Url string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	 Reqs []Request
	 Items []interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
