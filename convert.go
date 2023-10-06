package defcase

import (
	"github.com/gopher-summer/defcase/cases/PascalCase"
	"github.com/gopher-summer/defcase/cases/SCREAMING_SNAKE"
	"github.com/gopher-summer/defcase/cases/camelCase"
	"github.com/gopher-summer/defcase/cases/kebab-case"
	"github.com/gopher-summer/defcase/cases/none"
	"github.com/gopher-summer/defcase/cases/snake_case"
)

func Convert(name string, c Case) string {
	switch c {
	case NoCase:
		return none.Func(name)
	case Snak_case:
		return snake.Func(name)
	case CamelCase:
		return camel.Func(name)
	case Kebab_Case:
		return kebab.Func(name)
	case PascalCase:
		return Pascal.Func(name)
	case SCREAMING_SNAKE_CASE:
		return SCREAMING_SNAKE.Func(name)
	}

	return name
}
