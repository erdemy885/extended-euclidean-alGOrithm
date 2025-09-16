package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("First argument was not an integer")
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Second argument was not an integer")
	}
	if a < 0 || b < 0 {
		log.Fatal("You must enter two non-negative integers")
	}

	// gets table
	table := ExtEuc(a, b)

	//prints outputs
	fmt.Printf("Extended Euclidean Algorithm for %d and %d:\n", a, b)
	header := "|   u_1   |   v_1   |   u_2   |   v_2   |   u_3   |   v_3   |    q    |"
	printSep(len(header))
	fmt.Println(header)
	printSep(len(header))
	for _, row := range table {
		fmt.Printf("|%9d|%9d|%9d|%9d|%9d|%9d|%9d|\n", row[0], row[1], row[2], row[3], row[4], row[5], row[6])
	}
	printSep(len(header))
	lastRow := table[len(table)-1]
	fmt.Printf("gcd(%d, %d) = %d\n", a, b, lastRow[4])
	fmt.Printf("%d * %d + %d * %d = %d\n", lastRow[0], table[0][4], lastRow[2], table[0][5], lastRow[4])
}

// ExtEuc computes the Extended Euclidean Algorithm for two non-negative integers a and b.
// It returns a slice of 7-element integer arrays where each array represents a row of the
// table obtained during the computation of the Extended Euclidean Algorithm. The rows have
// the following structure:
//
//	[u_1, v_1, u_2, v_2, u_3, v_3, q]
//
// Where:
//
//	a * u_1 + b * u_2 = u_3
//	a * v_1 + b * v_2 = v_3
//	u_3 * q + v_3 = old_u_3
func ExtEuc(a, b int) [][7]int {
	// ensure a is greater than b
	if b > a {
		a, b = b, a
	}

	// store rows of division algorithm
	rows := make([][7]int, 0)
	rows = append(rows, [7]int{1, 0, 0, 1, a, b, 0})

	// continue until v_3 is 0
	for rows[len(rows)-1][5] != 0 {
		prevRow := rows[len(rows)-1]

		var row [7]int

		// new q is greatest int less than or equal to quotient of u_3 and v_3
		row[6] = prevRow[4] / prevRow[5]

		// new u_i is old v_i
		for i := 0; i < 6; i += 2 {
			row[i] = prevRow[i+1]
		}

		// new v_i = old u_i - (current q)(old v_i)
		for i := 1; i < 7; i += 2 {
			row[i] = prevRow[i-1] - prevRow[i]*row[6]
		}

		// add row to the rows slice
		rows = append(rows, row)
	}

	//return table
	return rows
}

// prints separator of length i with a newline after it
func printSep(i int) {
	for range i {
		fmt.Print("-")
	}
	fmt.Println()
}
