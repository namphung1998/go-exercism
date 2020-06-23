package matrix

import (
	"fmt"
	"strings"
)

type Matrix struct {
	config string
}

func New(config string) *Matrix {
	rows := strings.Split(config, "\n")
	fmt.Println(rows)

	return nil
}
