package main

import (
	"math"
	"testing"
)

func TestCoinChange(t *testing.T) {

	coins := []int{1, 3, 4}
	amount := 6

}

func solve(amount int, coins int[]) int{

	best := math.MaxInt

	if (amount == 0){
		
	}
	for(c : coins){

		if (c > amount){
			continue
		}

		1 + solve(amount-c)
	}

}
