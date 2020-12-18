// Day6 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"

	//	"sort"
	"strings"
)

type prog struct {
	name   string
	weight int
	childs []int
}

func parse(line string) (string, int, []string) {
	var weight int
	var childs []string
	var name string
	fmt.Sscanf(line, "%s (%d)", &name, &weight)

	items := strings.Split(line, " ")
	name = items[0]
	for i := 3; i < len(items); i++ {
		childs = append(childs, strings.Split(items[i], ",")[0])
	}
	fmt.Printf("%s (%d)", name, weight)
	return name, weight, childs
}

func find(name string, progs []prog, cnt int) int {
	for i := 0; i < cnt; i++ {
		if progs[i].name == name {
			return i
		}
	}
	return -1
}
func findParent(child int, progs []prog, cnt int) int {
	for i, p := range progs {
		for j, _ := range p.childs {
			if progs[i].childs[j] == child {
				//if progs[i].childs[j].name == child.name {
				return i
			}
		}
	}
	return -1
}
func main() {
	f, err := os.Open("input")
	file := bufio.NewReader(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var programs []prog
	cnt := 0
	for true {
		line, err := file.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		line = strings.TrimSpace(line)
		var p prog
		var c []string
		p.name, p.weight, c = parse(line)
		search := cnt
		//test if already loaded
		ptr := find(p.name, programs, cnt)
		if ptr == -1 {
			programs = append(programs, p)
			ptr = cnt
			cnt++
		} else {
			programs[ptr] = p
		}

		p.weight = 0

		for _, child := range c {
			cp := find(child, programs, search)
			if cp == -1 {
				p.name = child
				programs = append(programs, p)
				cp = cnt
				cnt++
			}
			programs[ptr].childs = append(programs[ptr].childs, cp)
		}
	}
	f.Close()

	//sort.Slice(programs, func(i, j int) bool { return programs[i].weight < programs[j].weight })

	f, err = os.Create("C:\\Code\\Go\\AoC2017\\Day7\\output")
	out := bufio.NewWriter(f)
	for _, p := range programs {
		s := fmt.Sprintf("%s (%d)", p.name, p.weight)
		if len(p.childs) > 0 {
			s += " ->"
			for i, c := range p.childs {
				s += " " + programs[c].name
				if i != len(p.childs)-1 {
					s += ","
				}
			}
		}
		s += "\r\n"
		out.WriteString(s)
	}
	out.Flush()
	f.Close()

	parent := findParent(0, programs, cnt)
	var root prog
	depth := 0
	for parent != -1 {
		root = programs[parent]
		parent = findParent(parent, programs, cnt)
		depth++
	}
	fmt.Printf("Depth %d\nRoot: %s w: %d", depth, root.name, root.weight)
}
