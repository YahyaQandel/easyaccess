package easyaccess
import (
	"fmt"
)

type anyType interface {
}

func Ln(x anyType) {
	fmt.Println(x)
}
