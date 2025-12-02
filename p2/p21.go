package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(r string) (int, int){
	a, b := strings.Split(r, "-")[0], strings.Split(r, "-")[1]
	a_int, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	b_int, err := strconv.Atoi(b)
	if err != nil {
		panic(err)
	}
	return a_int, b_int
}

func must[T any](v T, err error) T { if err != nil { panic(err) }; return v }

func count_invalid(a, b int, count int) int {
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)

		char_count := make(map[rune]int)
		for j := 0; j < len(s); j++ {
			char_count[rune(s[j])]++
		}

		mid := len(s) / 2
		first_half := s[:mid]
		second_half := s[mid:]
		if first_half == second_half {
			count += i
		}
	}
	return count
}

func main() {
	records, err := csv.NewReader(must(os.Open("input.csv"))).ReadAll()
	if err != nil {
		panic(err)
	}
	count := 0
	for _, record := range records[0] {
		a, b := parse(record)
		count = count_invalid(a, b, count)
	}

	fmt.Printf("count: %d\n", count)

	count2 := 0
	for _, record := range records[0] {
		a, b := parse(record)
		count2 = count_invalid_2(a, b, count2)
	}

	fmt.Printf("count2: %d\n", count2)
}
