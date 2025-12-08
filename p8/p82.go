package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func p2() {
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
	count := connect_pairs_2(pairs, points)
	fmt.Printf("[%v] %d\n", time.Since(start), count)
}

func create_set(points []Point) []Set {
	circuits := make([]Set, 0)
	for _, point := range points {
		circuit := make(Set)
		circuit.Add(point)
		circuits = append(circuits, circuit)
	}
	return circuits
}

func connect_pairs_2(pairs []PairDistance, points []Point) int {
	circuits := create_set(points)
	var i int
	for len(circuits) != 1 {
		i++
		circuits = insert_circuit(circuits, pairs[i])
	}

	return pairs[i].P1.x * pairs[i].P2.x
}
