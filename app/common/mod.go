package common

type Common struct {
	Error *Error
}

func New() *Common {
	return &Common{
		Error: NewError(),
	}
}
