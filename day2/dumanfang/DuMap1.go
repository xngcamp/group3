package main
import "fmt"

func main(){
	var inputString []string = "abababa"
	var outputString [len(inputString)]string

	maxLength := findMaxLength(inputString, outputString)

	//根据长度寻找第一个符合长度的不重复子串
	for i:=0; i<maxLength; i++{
		outputString = inputString[i,i+maxLength]
		for j:=0; j<maxLength; j++ {
			if outputString[j] = outputString[i] {
				break
			}

			if j=maxLength {
				for k:=0; k<maxLength; k++ {
					fmt.Println(outputString[j])
				}
				break
			}
		}
	}
}


//寻找最长不重复子串的长度
func findMaxLength(inputString string, outputString string) int {
	var curLength int
	var maxLength int

	var position [26]int
	for i:=0; i<26; i++{
		position[i] = -1
	}
	
	for i:=0; i<len(inputString); i++ {
		prevIndex := position[inputString[i]-'a']
		if prevIndex<0 ||　(i-prevIndex) > curLength {
			++curLength
		} else {
			if curLength >maxLength {
				maxLength = curLength
			}

			curLength = i - prevIndex
		}

		position[inputString[i]-'a'] = i

		if curLength > maxLength {
			maxLength = curLength
		}
	}
}
