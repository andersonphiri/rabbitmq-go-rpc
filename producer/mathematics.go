package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// SterlingApproximation computes the approximate number of digits of a factorial of n, for example if n = 20,
// the expected number of digits for 20! is, factorial(20) has 19 digits
// the formula: given a simple number say 235, number of digits = 1 + Floor(math.log(235) base 10) or Ceil (math.log(235) base 10) = 1 = Floor(2.371...) or Ceil(2.371..)
// = 3
// formula: log(n! base 10) = ( N*ln(N) - N + ln(2*pi*N)  ) / ln10
func SterlingApproximation(n int) int64 {
	ln10 := math.Log(10)
	floatN := float64(n)
	nLnN := floatN * math.Log(float64(n))
	pi := float64(22) / float64(7)
	ln2piN := math.Log( float64(2*n) * pi )
	result := ((nLnN)/ln10) - (floatN/ln10) + (ln2piN/ln10) 
	return int64(result)

}

func SterlingApproximationInt(n int) int {
	ln10 := math.Log(10)
	floatN := float64(n)
	nLnN := floatN * math.Log(float64(n))
	pi := float64(22) / float64(7)
	ln2piN := math.Log( float64(2*n) * pi )
	result := ((nLnN)/ln10) - (floatN/ln10) + (ln2piN/ln10) 
	return int(result)

}

func SterlingApproximationInt64(n int64) *big.Int {
	ln10 := math.Log(10)
	ln10Big := big.NewFloat(ln10)
	floatBigN := big.NewFloat(float64(n))
	nLnN :=  big.NewFloat(float64(n) * math.Log(float64(n)))
	pi := float64(22) / float64(7)
	ln2piN := big.NewFloat( math.Log( float64(2*n) * pi ))


	a := new(big.Float).Quo(nLnN, ln10Big)
	b := new(big.Float).Quo(floatBigN, ln10Big)
	c := new(big.Float).Quo(ln2piN, ln10Big)
	result := new(big.Float).Sub(a, b) //  a - b + c //(*nLnN   /ln10Big) - (floatN/ln10) + (ln2piN/ln10) 
	result = new(big.Float).Add(result, c)

	res := new(big.Int)
	result.Int(res)
	return res

}

func SterlingApproximationString(n string, base int ) (string, error) {
	nBig, ok := new(big.Int).SetString(n, base)
	if !ok {
		return "nil", fmt.Errorf("failed to parse string %s to int in base %d",n,base)
	}
	ln10 := math.Log(10)
	ln10Big := big.NewFloat(ln10)
	floatBigN := new(big.Float).SetInt(nBig)
	nBigFloat := new(big.Float).SetInt(nBig)
	float, _ := nBigFloat.Float64()
	nLnN :=  new(big.Float).Mul( nBigFloat, big.NewFloat(math.Log(  float  )))
	piRat := big.NewRat(22, 7)
	big2nPi := new(big.Float).Mul(floatBigN, new(big.Float).SetInt64(2))
	big2nPi = new(big.Float).Mul(big2nPi, new(big.Float).SetRat(piRat))
	big2nPiFloat, _ := big2nPi.Float64()
	ln2piN := big.NewFloat( math.Log( big2nPiFloat))


	a := new(big.Float).Quo(nLnN, ln10Big)
	b := new(big.Float).Quo(floatBigN, ln10Big)
	c := new(big.Float).Quo(ln2piN, ln10Big)
	result := new(big.Float).Sub(a, b) //  a - b + c //(*nLnN   /ln10Big) - (floatN/ln10) + (ln2piN/ln10) 
	result = new(big.Float).Add(result, c)

	res := new(big.Int)
	result.Int(res)
	return res.String(), nil

}

type StringArray []int

func (sr StringArray) String() string {
	var sb strings.Builder
	for _, v := range sr {
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	return sb.String()
}

func FactorialLarge(N int) (result []int) {
	expectedN := SterlingApproximationInt(N)
	result = make([]int, 0, expectedN)
	result = append(result, 1)
	var i int = 2
	for ; i <=N ; i++ {
		var carry int = 0
		for j := 0; j < len(result); j++ {
			val := result[j] * i + carry
			result[j] = val % 10 
			carry = val / 10
		}
		for carry != 0 {
			result = append(result, carry % 10)
			carry = carry / 10
		}
	}
	// rverse 
	Reverse(result, 0, len(result) - 1)
	return
}

func Reverse[T any](a []T, startIndex int, endIndex int) {
	for startIndex < endIndex {
		a[startIndex], a[endIndex] = a[endIndex], a[startIndex]
		startIndex++
		endIndex--
	}
}
