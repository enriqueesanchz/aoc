package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y, z int
}

type PairDistance struct {
	P1, P2 Point
	Distance float64
}

type Set map[Point]struct{}

func (s Set) Add(v Point) {
	s[v] = struct{}{}
}

func (s Set) Has(v Point) bool {
	_, ok := s[v]
	return ok
}

func main() {
	p1()
	p2()
}

func p1() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := parse(scanner)
	start := time.Now()
	distance_map := make_distances(points)
	pairs := make_slice(distance_map)
	circuits := connect_pairs(pairs, 1000)
	
	circuits_len := make([]int, 0)
	for _, circuit := range circuits {
		circuits_len = append(circuits_len, len(circuit))
	}
	sort.Slice(circuits_len, func(i, j int) bool {return circuits_len[i] > circuits_len[j]})

	count := 1
	for i := range 3 {
	count *= circuits_len[i]
	}
	fmt.Printf("[%v] %d\n", time.Since(start), count)
}

func parse(scanner *bufio.Scanner) []Point {
	points := make([]Point, 0)
	for scanner.Scan() {
		s := scanner.Text()
		point := strings.Split(s, ",")

		x, err := strconv.Atoi(point[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(point[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(point[2])
		if err != nil {
			panic(err)
		}

		points = append(points, Point{x, y, z})
	}

	return points
}

func distance(point1, point2 Point) float64 {
	var x, y, z float64
	x = float64(point1.x - point2.x)
	y = float64(point1.y - point2.y)
	z = float64(point1.z - point2.z)
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2))
}

func make_distances(points []Point) map[[2]Point]float64 {
	distance_map := make(map[[2]Point]float64)

	for i := range points {
		for j := i; j < len(points); j++ {
			key1 := [2]Point{points[i], points[j]}
			key2 := [2]Point{points[j], points[i]}
			if j == i {
				distance_map[key1] = float64(9999999999)
				distance_map[key2] = float64(9999999999)
			} else {
				distance_map[key1] = distance(points[i], points[j])
				distance_map[key2] = float64(9999999999)
			}
		}
	}

	return distance_map
}

func make_slice(distance_map map[[2]Point]float64) []PairDistance {
	pairs := make([]PairDistance, 0, len(distance_map))

	for k, v := range distance_map {
		// avoid duplicate reversed pairs
		if k[0] == k[1] {
			continue
		}

		pairs = append(pairs, PairDistance{
			P1:       k[0],
			P2:       k[1],
			Distance: v,
		})
	}

	// sort by shortest distance
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Distance < pairs[j].Distance
	})

	return pairs
}

func connect_pairs(pairs []PairDistance, n int) []Set {
	circuits := make([]Set, 0)
	var i int
	for i = range n {
		circuits = insert_circuit(circuits, pairs[i])
	}

	return circuits
}

func insert_circuit(circuits []Set, pair PairDistance) []Set {
	var idx1, idx2 = -1, -1

	// Find which circuits contain the points
	for i := range circuits {
		if circuits[i].Has(pair.P1) {
			idx1 = i
		}
		if circuits[i].Has(pair.P2) {
			idx2 = i
		}
	}

	// Case 1: both points are in the same circuit
	if idx1 != -1 && idx1 == idx2 {
		return circuits
	}

	// Case 2: both points in different circuits
	if idx1 != -1 && idx2 != -1 {
		for p := range circuits[idx2] {
			circuits[idx1].Add(p)
		}

		// remove the merged circuit
		circuits = slices.Delete(circuits, idx2, idx2+1)
		return circuits
	}

	// Case 3: only one point is in a circuit
	if idx1 != -1 {
		circuits[idx1].Add(pair.P2)
		return circuits
	}
	if idx2 != -1 {
		circuits[idx2].Add(pair.P1)
		return circuits
	}

	// Case 4: neither point is in a circuit
	newCircuit := make(Set)
	newCircuit.Add(pair.P1)
	newCircuit.Add(pair.P2)
	return append(circuits, newCircuit)
}

