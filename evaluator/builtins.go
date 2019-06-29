package evaluator

import (
	"pl-i18n/evaluator/builtins"
	"pl-i18n/object"
)

var builtinFunctions = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: builtins.Len,
	},

	"first": &object.Builtin{
		Fn: builtins.First,
	},

	"last": &object.Builtin{
		Fn: builtins.Last,
	},

	"rest": &object.Builtin{
		Fn: builtins.Rest,
	},

	"push": &object.Builtin{
		Fn: builtins.Push,
	},

	"print": &object.Builtin{
		Fn: builtins.Print,
	},
}
