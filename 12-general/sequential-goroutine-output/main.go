package main

// evens = 2,4,6,8,10,...n
// odds = 1,3,5,7,9,...n
// Output: 1,2,3,4,5,6,7,8,9,10,...n

import (
	"fmt"
	"sync"
)

var n = 100

func gerenateEvens(oddChan, evenChan chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 2; i <= n; i = i + 2 {
		// if i != n-1 {
		<-oddChan
		// }
		fmt.Println(i)
		evenChan <- i
	}
	close(evenChan)
}

func gerenateOdds(oddChan, evenChan chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 1; i <= n; i = i + 2 {
		fmt.Println(i)
		oddChan <- i
		if i != n {
			<-evenChan
		}
	}
	close(oddChan)
}

func main() {
	var wg sync.WaitGroup
	oddChan := make(chan int)
	evenChan := make(chan int)

	wg.Add(2)
	go gerenateOdds(oddChan, evenChan, &wg)
	go gerenateEvens(oddChan, evenChan, &wg)
	wg.Wait()

}
