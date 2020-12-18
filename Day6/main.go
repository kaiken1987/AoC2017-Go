// Day6 project main.go
package main

import (
	"fmt"
)

const regCnt = 16

type memory struct {
	register [regCnt]int8
}

func find(cur memory, hist []memory, cnt int) int {
	for i := 0; i < cnt; i++ {
		if hist[i] == cur {
			return i
		}
	}
	return -1
}

func main() {
	var current memory
	var inital = [regCnt]int8{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
	//var inital = [regCnt]int8{0, 2, 7, 0}
	current.register = inital

	cnt := 0
	history := make([]memory, 1)
	history[cnt] = current
	cnt++

	for true {
		var high int8 = 0
		highIdx := 0
		for i := 0; i < regCnt; i++ {
			if current.register[i] > high {
				high = current.register[i]
				highIdx = i
			}
		}
		current.register[highIdx] = 0
		for i := int8(1); i <= high; i++ {
			current.register[(highIdx+int(i))%regCnt]++
		}
		found := find(current, history, cnt)
		if found > 0 {
			fmt.Printf("loop size: %d\n", cnt-found)
			break
		}
		history = append(history, current)
		cnt++
	}
	fmt.Printf("Found on run %d\n State: ", cnt)
	for i := 0; i < regCnt; i++ {
		fmt.Printf("%d, ", current.register[i])
	}

}
