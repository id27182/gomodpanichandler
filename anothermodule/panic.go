package anothermodule

import (
	"fmt"
)

func CausePanic() {
	panic(fmt.Errorf("panic from external anothermodule"))
}
