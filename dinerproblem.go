package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type chopStick struct{ sync.Mutex }
type philosopher struct {
	pId                 int
	leftChop, rightChop *chopStick
	permissionCounter   int
}

func (p *philosopher) askPermission() bool {
	var mut sync.Mutex
	var flag bool = false
	mut.Lock()
	if p.permissionCounter != 3 {
		p.permissionCounter = p.permissionCounter + 1
		p.leftChop.Lock()
		p.rightChop.Lock()
		flag = true
	}
	mut.Unlock()
	return flag
}
func eat(p *philosopher, wg *sync.WaitGroup, list *[]int) {
	if p.askPermission() {
		fmt.Printf("starting to eat %d \n", p.pId)
		fmt.Printf("eating...%d \n", p.pId)
		p.leftChop.Unlock()
		p.rightChop.Unlock()
		fmt.Printf("finishing eating %d \n", p.pId)
		*list = append(*list, p.pId)
	}
	wg.Done()

}
func generateRandomNumber() int {
	return rand.Intn(5)
}
func main() {
	var philoList []*philosopher
	var wg sync.WaitGroup
	var philExecutionSummary = make([]int, 0)
	philoList = make([]*philosopher, 5)
	CSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(chopStick)
	}
	for x := range philoList {
		philoList[x] = &philosopher{
			x + 1,
			CSticks[x],
			CSticks[(x+1)%5],
			0,
		}
	}
	for len(philExecutionSummary) < 15 {
		i := generateRandomNumber()
		k := generateRandomNumber()
		wg.Add(2)
		go eat(philoList[i], &wg, &philExecutionSummary)
		go eat(philoList[k], &wg, &philExecutionSummary)
		if runtime.NumGoroutine() > 3 {
			wg.Wait()
		}
	}
	wg.Wait()
	fmt.Printf("Summary of sequence of execution and number of times of execution %d %v", len(philExecutionSummary), philExecutionSummary)

}
