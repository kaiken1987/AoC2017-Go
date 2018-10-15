// day4 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func day4(line string) bool {
	valid := true
	toks := strings.Split(line, " ")
	for i := 0; i < len(toks)-1; i++ {
		key := toks[i]
		for j := i + 1; j < len(toks); j++ {
			if strings.Compare(toks[j], key) == 0 {
				valid = false
				break
			}
		}
	}
	return valid
}
func compare(key, tok string) bool {
	len1 := utf8.RuneCountInString(key)
	b1 := []rune(key)
	len2 := utf8.RuneCountInString(tok)
	b2 := []rune(tok)

	for i := 0; i < len1; i++ {
		found := false
		for j := 0; j < len2; j++ {
			if b1[i] == b2[j] {
				found = true
			}
		}
		if found == false {
			return false
		}

	}
	for i := 0; i < len2; i++ {
		found := false
		for j := 0; j < len1; j++ {
			if b2[i] == b1[j] {
				found = true
			}
		}
		if found == false {
			return false
		}

	}
	return true
}
func day4b(line string) bool {
	toks := strings.Split(line, " ")
	for i := 0; i < len(toks)-1; i++ {
		key := toks[i]
		for j := i + 1; j < len(toks); j++ {
			if compare(toks[j], key) && compare(toks[j], key) {
				return false
			}
		}
	}
	return true
}

func main() {
	f, err := os.Open("input")
	file := bufio.NewReader(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	valid := 0
	for true {
		line, err := file.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if day4b(line) {
			valid++
		}
	}
	fmt.Printf("Valid codes: %d", valid)
}
