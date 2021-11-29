package main

import (
	"fmt"
	"runtime"
)

type Transactions struct {
	name  string
	price int
}

func main() {
	runtime.GOMAXPROCS(2)

	transactions := []Transactions{
		{"Coklat", 1100},
		{"Relaxa", 2100},
		{"Nugget", 3100},
		{"Mentos", 4100},
	}

	var c = make(chan Transactions)

	go transaction(transactions, c)

	result := make([]Transactions, len(transactions))

	for i := range result {
		(result)[i] = <-c
	}
	printResult(result)

}

func transaction(transactions []Transactions, c chan Transactions) {
	for _, transactions := range transactions {
		c <- transactions
	}
}

func printResult(transactions []Transactions) {
	for _, transactions := range transactions {
		fmt.Println(transactions.name, "\t: ", transactions.price)
	}
}
