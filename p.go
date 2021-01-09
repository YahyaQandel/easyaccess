package prnt
import (
	"fmt"
)

type anything interface {
}

func Ln(x anything) {
	fmt.Println(x)
}