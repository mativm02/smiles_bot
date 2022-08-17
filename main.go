package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/matisidler/smiles-bot/logic"
	"github.com/matisidler/smiles-bot/models"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	var locker sync.Mutex
	days := []string{"04", "05", "11", "12", "18", "19", "25", "26", "28"}
	var monthList []models.Day
	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			list, err := logic.GetRequest("2022", "10", days[index], "EZE", "CUN", index != 0 && index%2 != 0)
			if err != nil {
				fmt.Println(err)
				return
			}
			locker.Lock()
			monthList = append(monthList, list...)
			locker.Unlock()
		}(i)
	}
	wg.Wait()
	monthList = logic.Remove0Values(monthList)
	logic.LookForBestPrices(monthList)
	if len(monthList) < 10 {
		fmt.Println(monthList)
		return
	}
	fmt.Println(monthList[0:10])
	fmt.Println(time.Since(start))
}
