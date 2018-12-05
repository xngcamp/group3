package main

import {
	"errors"
	"fmt"
}

func main() {
	var solve map[string]int
	solve("((1)23(45))(aB)", 0) = 10 // the opening brace at index 0 matches the closing brace at index 10
	solve("((1)23(45))(aB)", 1) = 3 
	solve("((1)23(45))(aB)", 2) = -1 // there is no opening bracket at index 2, so return -1
	solve("((1)23(45))(aB)", 6) = 9
	solve("((1)23(45))(aB)", 11) = 14
	solve("((>)|?(*'))(yZ)", 11) = 14
}

func Solution(str string, idx int) (int, error) {
    return 0, errors.New("Not an opening bracket") 
}
