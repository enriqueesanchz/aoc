package main

import (
	"fmt"
	"os"
	"bufio"
)

func p2() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	conns := parse(scanner)

	start := "svr"
	stop := "fft"
	count_1 := get_count(conns, start, stop)

	start = "fft"
	stop = "dac"
	count_2 := get_count(conns, start, stop)

	start = "dac"
	stop = "out"
	count_3 := get_count(conns, start, stop)

	start = "svr"
	stop = "dac"
	count_4 := get_count(conns, start, stop)

	start = "dac"
	stop = "fft"
	count_5 := get_count(conns, start, stop)

	start = "fft"
	stop = "out"
	count_6 := get_count(conns, start, stop)

	total := count_1*count_2*count_3+count_4*count_5*count_6
	fmt.Printf("[p2] %d\n", total)
}

func memo_dfs(conns map[string][]string, node, start, stop string, memo map[string]int64) int64 {
	if val, ok := memo[node]; ok {
		return val
	}
	
	if node == start {
		return 0
	}

	if node == stop {
		return 1
	}

	var c int64
	next := conns[node]
	for _, n := range next {
		c += memo_dfs(conns, n, start, stop, memo)
	}

	memo[node] = c
	return c
}

func get_count(conns map[string][]string, start, stop string) int64 {
	var count int64
	memo := make(map[string]int64, 0)
	for _, n := range conns[start] {
		count += memo_dfs(conns, n, start, stop, memo)
	}
	return count
}
