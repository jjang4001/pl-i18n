package builtins

import (
	"fmt"
	"pl-i18n/object"
)

func Print(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return NULL
}
