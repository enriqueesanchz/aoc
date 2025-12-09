package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	corners := parse(scanner)
	max_corners, max_area := calc_max_area(corners)
	fmt.Println("Max area:", int64(max_area), max_corners)
}

type Point struct {
	x, y int
}

func parse(scanner *bufio.Scanner) []Point {
	corners := make([]Point, 0)
	for scanner.Scan() {
		s := scanner.Text()
		x, err := strconv.Atoi(strings.Split(s, ",")[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strings.Split(s, ",")[1])
		if err != nil {
			panic(err)
		}
		corners = append(corners, Point{x, y})
	}

	return corners
}

func area(corners [2]Point) float64 {
	return (math.Abs(float64(corners[0].x-corners[1].x))+1.0) * (math.Abs(float64(corners[0].y-corners[1].y))+1.0)
}

func calc_max_area(corners []Point) ([2]Point, float64) {
	max_area := 0.0
	corners_pair := [2]Point{}

	for i := range corners {
		for j := i; j < len(corners); j++ {
			a := area([2]Point{corners[i], corners[j]})
			if a > max_area {
				max_area = a
				corners_pair = [2]Point{corners[i], corners[j]}
			}
		}
	}
	return corners_pair, max_area
}
