package cycle

import "github.com/jhagege/leftpad"

var DEFAULT_CHAR = ' '

func DoublePad(s string, i int) string {
	return leftpad.Format(s+s, i)
}
