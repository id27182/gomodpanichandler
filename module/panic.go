package module

import (
	"fmt"
)

func CausePanic() {
	panic(fmt.Errorf("panic from external module"))
}