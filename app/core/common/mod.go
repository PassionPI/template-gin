package common

type Common struct {
	Error  *Error
	Parser *Parser
}

func New() *Common {
	return &Common{
		Error:  NewError(),
		Parser: NewParser(),
	}
}
