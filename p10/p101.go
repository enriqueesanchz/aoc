package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func generateCombinations(buttons [][]int) [][][]int {
    var result [][][]int
    n := len(buttons)

    var backtrack func(start, k int, curr [][]int)
    backtrack = func(start, k int, curr [][]int) {
        if k == 0 {
            combo := make([][]int, len(curr))
            copy(combo, curr)
            result = append(result, combo)
            return
        }

        for i := start; i <= n-k; i++ {
            backtrack(i+1, k-1, append(curr, buttons[i]))
        }
    }

    // generate combinations by size (1 to n)
    for size := 1; size <= n; size++ {
        backtrack(0, size, [][]int{})
    }

    return result
}

type row struct {
	lights string
	buttons [][]int
	jolts []int
}

func parse(scanner *bufio.Scanner) []row {
	var rows []row
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Split(s, "] ")
		l := strings.Trim(fields[0], "[]")

		fields = strings.Split(fields[1], " {")
		button_str := strings.ReplaceAll(fields[0], "(", "")
		button_str = strings.ReplaceAll(button_str, ")", "")
		button_strs := strings.Fields(button_str)
		var buttons [][]int
		for _, b := range button_strs {
			var button []int
			for c := range strings.SplitSeq(b, ",") {
				var val int
				fmt.Sscanf(c, "%d", &val)
				button = append(button, val)
			}
			buttons = append(buttons, button)
		}

		var jolts []int
		jolt_str := strings.Trim(fields[1], "{}")
		for jolt := range strings.SplitSeq(jolt_str, ",") {
			var val int
			fmt.Sscanf(jolt, "%d", &val)
			jolts = append(jolts, val)
		}

		r := row{
			lights: l,
			buttons: buttons,
			jolts: jolts,
		}
		rows = append(rows, r)
	}
	return rows
}

func goal(r row) []int {
	var g []int

	for _, c := range r.lights {
		if c == '#' {
			g = append(g, 1)
		} else {
			g = append(g, 0)
		}
	}

	return g
}

func check_valid(r row, comb [][]int) bool {
	light_state := make([]int, len(r.lights))
	for _, c := range comb {
		for _, b := range c {
			light_state[b] ^= 1
		}
	}

	g := goal(r)
	for i := range g {
		if g[i] != light_state[i] {
			return false
		}
	}

	return true
}

func p1() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := parse(scanner)
	
	count := 0
	for _, r := range rows {
		combs := generateCombinations(r.buttons)
		for _, comb := range combs {
			if check_valid(r, comb) {
				count += len(comb)
				break
			}
		}
	}
	fmt.Printf("[p1] %d\n", count)
}

func main() {
	p1()
}
