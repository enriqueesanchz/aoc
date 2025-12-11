package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	p1()
	p2()
}

func p1() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	conns := parse(scanner)

	count := 0
	for _, n := range conns["you"] {
		count += dfs(conns, n, "you")
	}
	fmt.Printf("[p1] %d\n", count)
}

func parse(scanner *bufio.Scanner) map[string][]string {
	conns := make(map[string][]string, 0)
	for scanner.Scan() {
		s := scanner.Text()

		fields := strings.Split(s, ": ")
		from := fields[0]
		to := strings.Fields(fields[1])
		conns[from] = to
	}

	return conns
}

func dfs(conns map[string][]string, node string, stop string) int {
	if node == stop {
		return 0
	}

	if node == "out" {
		return 1
	}

	c := 0
	next := conns[node]
	for _, n := range next {
		c += dfs(conns, n, stop)
	}
	return c
}
