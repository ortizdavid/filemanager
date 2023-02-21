package helpers

import (
	"fmt"
	"strings"
)

func Contains(s []string, str string) bool {
	for _, item := range s {
		if str == item {
			return true
		}
	}
	return false
}

func PrintChar(ch string, chSize int) {
	for i := 0; i < chSize; i++ {
		fmt.Print(ch)
	}
	fmt.Println()
}

func GetNameFromUrl(url string) string {
	splitedUrl := strings.Split(url, "/")
	name := splitedUrl[len(splitedUrl)-1]
	return name
}