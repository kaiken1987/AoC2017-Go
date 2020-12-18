// Day5 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day5(addrs []int, cnt int) {
	steps := 0
	add := 0
	for add >= 0 && add <= cnt {
		new := addrs[add]
		addrs[add]++
		add = add + new
		steps++
	}
	fmt.Printf("Exit Steps: %d", steps)
}
func day5b(addrs []int, cnt int) {
	steps := 0
	add := 0
	for add >= 0 && add < cnt {
		new := addrs[add]
		if addrs[add] >= 3 {
			addrs[add]--
		} else {
			addrs[add]++
		}
		add = add + new
		steps++
	}
	fmt.Printf("Exit Steps: %d\n", steps)
}
func main() {
	f, err := os.Open("input")
	//f, err := os.Open("test")
	file := bufio.NewReader(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cnt := 0
	addrs := make([]int, 1050)
	for true {
		line, err := file.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		line = strings.TrimSpace(line)
		addrs[cnt], _ = strconv.Atoi(line)
		cnt++
	}
	day5b(addrs, cnt)
}
