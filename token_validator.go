package main

import (
	"fmt"
	"flag"
	"strings"
	"reflect"
)

//splitTokenString split string to token pares
func splitTokenString(inputString string) []string {
	s := strings.Split(inputString, ",")
	return s
}


func splitStringToLetters(str, string)[]string{
	lettersArray := make([]string, 0)
	var it string
	for item := range str{
		it = string(str[item])
		lettersArray = append(lettersArray, it)
	}
			
}

func checkTokenString(tokenItemArray, inputStr []string) bool {
	// Think!!!
	m := make(map[string]bool)
	for i := 0; i < len(tokenItemArray); i++ {
			m[tokenItemArray[i]] = true
	}
	for _, item := range inputStr{
		if _, ok := m[item]; !ok {
			return false
		}
	}
	return true
}

func splitArrayToLetters(arr []string) []string {
	
	lettersArray := make([]string, 0)
	leterString := strings.Join(arr, "")
	var it string
	for item := range leterString{
		it = string(leterString[item])
		lettersArray = append(lettersArray, it)
	}
	return lettersArray
}

func main()  {
	tokens := flag.String("tokens", "", "token separate by ',' ")
	inputString := flag.String("inputstr", "", "string for validation!")
	flag.Parse()
	fmt.Printf("Tokens: %v Input Str: %v \n", *tokens, *inputString)
	ss:= splitTokenString(*tokens)
	tokenLetters := splitArrayToLetters(ss)
	if ss != nil{
		fmt.Println(ss)
	} 
}
