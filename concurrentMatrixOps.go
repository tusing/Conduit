package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/* CONDUIT ASSIGNMENT: PARALLELISM (Tushar)
- This is my first program in Go. Essentially, my
    "Hello World". I'm not exactly sure what I'm
    doing...
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	/* TASK 1: FIBONACCI */
	fmt.Println("\n------  TASK 1  ------")
	/* Basic idea:
	   f(x):
	       x < 2 return x
	       i = new thread f(x-1)
	       j = new thread f(x-2)
	       return i+j
	*/

	fmt.Println("----------------------\n")

	/* TASK 2: VECTOR ADDITION
	   - Number of threads must be able to divide into
	       array length
	   - This issue can be fixed by appending an array
	       with 0's s.t. len(array) % num_threads = 0,
	       or separately dealing with the tail case
	*/
	fmt.Println("\n------  TASK 2  ------")
	a := rand.Perm(60)
	b := rand.Perm(60)
	fmt.Println("TOTAL SUM", task_2(a, b, runtime.NumCPU()))
	fmt.Println("----------------------\n")

	/* TASK 3: IN-PLACE MATRIX TRANSPOSE
	   - Currently only works for square matrices.
	   - A naive solution to this would be to pad the matrix
	       with 0's to make it a square matrix.
	   - This is an inefficient method of computing matrix
	       transpose.
	*/
	fmt.Println("\n------  TASK 3  ------")
	var wg sync.WaitGroup

	// Make a 5x5 matrix.
	var matrix [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	original := matrix // We'll use this for Task 4.
	fmt.Println("Original Matrix: ", matrix)

	// Transpose.
	for i := 0; i < len(matrix); i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < len(matrix); j++ {
				if i != j && j > i {
					temp := matrix[i][j]
					matrix[i][j] = matrix[j][i]
					matrix[j][i] = temp
				}
			}
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Transposed Matrix: ", matrix)
	fmt.Println("----------------------\n")

	/* TASK 4: MATRIX MULTIPLY
	   For the sake of simplicity, we'll be multiplying
	   the matrix generated in TASK 3 by itself.
	   - This only works on square matrices.
	*/
	fmt.Println("\n------  TASK 4  ------")
	transpose := matrix // Transpose of original.
	// Since we already have the transpose of the
	// second matrix, we can proceed. Otherwise, we'd
	// take the transpose of the second matrix, as above.
	fmt.Println("Multiplying by itself: ", original)
	var result [5][5]int
	for i := 0; i < len(original); i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < len(matrix); j++ {
				result[i][j] = dot_product(original[i],
					transpose[j])
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Result:", result)
	fmt.Println("----------------------\n")
}

func task_2(a []int, b []int, num_channels int) int {
	c := make(chan int)
	stride := len(a) / num_channels

	for i := 0; i < num_channels; i++ {
		i_start := i * stride
		i_end := i*stride + stride
		go add_arrays(a[i_start:i_end], b[i_start:i_end], c)
	}

	sum := 0
	for i := 0; i < num_channels; i++ {
		sum += <-c
	}

	return sum
}

func add_arrays(arr1 []int, arr2 []int, c chan int) {
	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i] + arr2[i]
	}
	fmt.Println("Added", arr1, arr2, " -> ", sum)
	c <- sum
}

/*TODO: Figure out how to avoid needing to put
[5] before int, so we can just generically use
[]int. */
func dot_product(arr1 [5]int, arr2 [5]int) int {
	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i] * arr2[i]
	}
	return sum
}
