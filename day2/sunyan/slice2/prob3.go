package slice2

import "fmt"

func main() {
	fmt.Println(IsValidISBN("3-598-P1581-X"))
}
func IsValidISBN(str string) bool {
	isbn := []byte(str)
	sum, index, num := 0, 10, 0
	for i := 0; i < len(isbn); i++ {
		if isbn[i] != '-' {
			if isbn[i] >= '0' && isbn[i] <= '9' {
				sum += int(isbn[i]-'0') * index
				num++
			} else if isbn[i] == 'X' && i == len(isbn)-1 {
				sum += 10
				num++
			}
			index--
		}
	}
	if num == 10 {
		if sum%11 == 0 {
			return true
		}
	}
	return false
}
