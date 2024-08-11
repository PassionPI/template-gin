package common

import "app-ink/app/core/dependency"

type Common struct {
	Parser *Parser
}

func New(dep *dependency.Dependency) *Common {
	return &Common{
		Parser: NewParser(),
	}
}
