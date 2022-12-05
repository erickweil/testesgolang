package inputhelper

import (
	"strings"
)

var txt string
var args []string
var txti int

func InitParser(_txt string) {
	txt = _txt
	txti = 0

	args = strings.Split(_txt," ")
}

func NextArgument() string {
	txti++
	return args[txti-1]
}
