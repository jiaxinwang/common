package lazy

import (
	"errors"
	"strings"
)

func disassembleTag(tag string) (name, foreignkeyTable, foreignkey string, err error) {
	part := strings.Split(tag, `;`)
	for _, v := range part {
		if strings.Contains(v, `:`) {
			pair := strings.Split(v, ":")
			if len(pair) != 2 {
				err = errors.New("wrong format")
				return
			}
			switch pair[0] {
			case `foreign`:
				f := strings.Split(pair[1], ".")
				if len(f) != 2 {
					err = errors.New("wrong format")
					return
				}
				foreignkeyTable = f[0]
				foreignkey = f[1]
			}
		} else {
			name = v
		}
	}
	return
}
