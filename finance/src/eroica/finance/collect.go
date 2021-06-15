package main

import (
	"log"
	"sync"

	da "eroica.finance/src/eroica/finance/data_access"
	"eroica.finance/src/eroica/finance/service"
)

func main() {
	defer da.CloseDb()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.CollectStockBasic()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.CollectTradeCal()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.CollectNamechange()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.CollectHsConst()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.CollectStockCompany()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.CollectStkManagers()
	}()
	wg.Wait()
	log.Println("Success.")
}
