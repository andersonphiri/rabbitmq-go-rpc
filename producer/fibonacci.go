package main

import (
	"fmt"
	"math/big"
)

func Fib(n int) int {
	if n <= 2 {
		return 1
	}
	prev := 1
	curr := 2
	for i :=  4; i < n ; i++ {
		temp := curr + prev
		prev = curr
		curr = temp
	}
	return curr
}

func FibInt64(n int) int64 {
	if n <= 2 {
		return 1
	}
	// var prevPrev int64 = 1
	var prev int64 = 1
	var curr int64 = 2
	for i := 4; i <= n; i++ {
		temp := curr + prev
		prev = curr
		curr = temp
	}
	return curr
}

func FibBingInt( n string, base int) (string, error) {
	
	limit, ok := new(big.Int).SetString(n, base)
	if !ok {
		return "", fmt.Errorf("unable to parse string(%s) to big int", n)
	}
	prev := big.NewInt(1)
	curr := big.NewInt(2)
	if limit.Cmp(prev) <= 0 {
		return prev.String(), nil
	}
	
	increment := big.NewInt(1)

	for i := big.NewInt(4); i.Cmp(limit) <= 0; i = i.Add(i, increment) {
		temp := new(big.Int).Add(curr, prev)
		_ = prev.Set(curr)
		_ = curr.Set(temp)
	}
	return curr.String(), nil
}


