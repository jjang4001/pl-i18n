package builtins

import (
	"pl-i18n/object"
)

func Last(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}
