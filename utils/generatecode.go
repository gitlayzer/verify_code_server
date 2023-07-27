package utils

import (
	"fmt"
	"math/rand"
)

func GenerateCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}
