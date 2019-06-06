package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func newMatrix(n int) [][]int64 {
	m := make([][]int64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int64, n)
	}
	return m
}

func main() {
		if len(os.Args) < 2 {
			fmt.Println("usage: go run matrix.go N")
		}

		// Parse the N argument from the commandline.
		n, err := strconv.Atoi(os.Args[1])
		if err != nil || n <= 0 {
			log.Fatalf("need a positive in for N, not %v", os.Args[1])
		}


	// Initialize our matrices with some values.
	a := newMatrix(n)
	b := newMatrix(n)
	c := newMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = int64(i*n + j)
			b[i][j] = int64(j*n + i)
			// c is already initialized to all zeros.
		}
	}

	begin := time.Now()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cSum := int64(0)
			for k := 0; k < n; k++ {
				cSum += int64(a[i][k] * b[k][j])
			}
			c[i][j] = cSum
		}
	}

	end := time.Now()
	fmt.Println("time: ", end.Sub(begin).Nanoseconds())

	// Print contents of C for debugging.
	//fmt.Println(c)

	total := int64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			total += int64(c[i][j])
		}
	}

	fmt.Println("sum: ", total)
	// This sum should be 450 for N=3, 3680 for N=4, and 18250 for N=5.
	knownSums := map[int]int{
		3: 450,
		4: 3680,
		5: 18250,
	}
	want := int64(knownSums[n])
	if want != 0 && total != want {
		fmt.Printf("Wanted sum %v but got %v\n", want, total)
	}
}
