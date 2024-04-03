package main

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s = Starting point of Sam's house
 *  2. INTEGER t = Ending point of Sam's house
 *  3. INTEGER a = Apple Tree Distance point
 *  4. INTEGER b = Orange Tree Distance point
 *  5. INTEGER_ARRAY apples = Slice of Apples fall down
 *  6. INTEGER_ARRAY oranges = Slice of Orange fall down
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var applesFalls int32
	var orangesFalls int32
	var d int32
	// Calculate Apples falls count on Sam's House
	for _, apple := range apples {
		d = apple + a
		if d >= s && d <= t {
			applesFalls = applesFalls + 1
		}
	}
	// Calculate Oranges falls count on Sam's House
	for _, orange := range oranges {
		d = orange + b
		if d >= s && d <= t {
			orangesFalls = orangesFalls + 1
		}
	}

	fmt.Printf("%d\n%d\n", applesFalls, orangesFalls)

}

func main() {
	var s int32 = 7
	var t int32 = 11
	var a int32 = 5
	var b int32 = 15
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	countApplesAndOranges(s, t, a, b, apples, oranges)
}
