package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Solution("((>)|?(*'))(yZ)", 11))
}

/**
  Day2slice1
*/
func Solution(str string, idx int) (int, error) {
	strs := []byte(str)
	var num int = 1
	if strs[idx] != '(' {
		return -1, nil
	}
	for i := idx + 1; i < len(strs); i++ {
		if strs[i] == '(' {
			num++
		} else if strs[i] == ')' {
			num--
		}
		if num == 0 {
			return i, nil
		}
	}
	return 0, errors.New("Not an opening bracket")
}
