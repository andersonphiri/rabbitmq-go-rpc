package main

import (
	// "math"
	"math/big"
	"testing"
)

const (
	pass = "\u2713"
	fail = "\u2717"
)


func TestSterlingApproximation(t *testing.T) {
	var table = []struct{
		Name string
		Input int
		Output int64 
		OutputRound int64
		SkipCheck bool
	}{
		{Name: "20", Input: 20 , Output: 19, OutputRound: 19},
		{Name: "10Pow9", Input: 1_000_000_000 , Output: 8_565_705_527 , OutputRound: 8_566_000_000},
		// {Name: "20", Input: 20 , Output: 19},
		// {Name: "20", Input: 20 , Output: 19},
		{Name: "10Pow18", Input: 100000_000_000_000_000 , Output: -1 , OutputRound: -1, SkipCheck: true},
	}
	for _, input := range table {
		tfunc := func (tf *testing.T)  {
			// given
			t.Logf("\tgiven the need to compute the number of digits for %d!\n", input.Input)
			{
				// when 
				t.Logf("when %d! is computed using sterling approximation\n", input.Input) 
				{
					// should 
					result := SterlingApproximation(input.Input)
					if input.SkipCheck {
						t.Logf("\t%s Test: %s, %d! should be %d digits and it got %d, expected rounded to %d (test to be igonred)\n", pass,input.Name, input.Input, input.Output, result, input.OutputRound)
						return
					}
					if result != input.Output  {
						t.Fatalf("\t%s Test: %s, %d! should be %d digits but it got %d, expected rounded to %d\n", fail,input.Name, input.Input, input.Output, result, input.OutputRound)
					} else {
						t.Logf("\t%s Test: %s, %d! should be %d digits and it got %d, expected rounded to %d\n", pass,input.Name, input.Input, input.Output, result, input.OutputRound)
					}
				}
			}
		}
		t.Run(input.Name,tfunc)
	}
}


func TestSterlingApproximationInt64(t *testing.T) {
	var table = []struct{
		Name string
		Input int64
		Output *big.Int 
		OutputRound int64
		SkipCheck bool
	}{
		{Name: "20", Input: 20 , Output: big.NewInt(19), OutputRound: 19},
		{Name: "10Pow9", Input: 1_000_000_000 , Output: big.NewInt(8_565_705_527) , OutputRound: 8_565_705_527, SkipCheck: true},
		// {Name: "20", Input: 20 , Output: 19},
		// {Name: "20", Input: 20 , Output: 19},
		{Name: "10Pow18", Input: 1000_000_000_000_000_000 , Output: big.NewInt(1) , OutputRound: -1, SkipCheck: true},
		//{Name: "maxInt64", Input: math.MaxInt64 , Output: big.NewInt(1) , OutputRound: -1, SkipCheck: true},
	}
	for _, input := range table {
		tfunc := func (tf *testing.T)  {
			// given
			t.Logf("\tgiven the need to compute the number of digits for %d!\n", input.Input)
			{
				// when 
				t.Logf("when %d! is computed using sterling approximation\n", input.Input) 
				{
					// should 
					result := SterlingApproximationInt64(input.Input)
					if input.SkipCheck {
						t.Logf("\t%s Test: %s, %d! should be %d digits and it got %s, expected rounded to %d (test to be igonred)\n", 
						pass,input.Name, input.Input, input.Output.Int64(), result.String(), input.OutputRound)
						return
					}
					if result.Int64() != input.OutputRound  {
						t.Fatalf("\t%s Test: %s, %d! should be %d digits but it got %s, expected rounded to %d\n", 
						fail,input.Name, input.Input, input.Output.Int64(), result.String(), input.OutputRound)
					} else {
						t.Logf("\t%s Test: %s, %d! should be %d digits and it got %s, expected rounded to %d\n", 
						pass,input.Name, input.Input, input.Output, result.String(), input.OutputRound)
					}
				}
			}
		}
		t.Run(input.Name,tfunc)
	}
}


func TestFactorialLarge(t *testing.T) {
	table := []struct {
		expected string
		actual   string
		n        int
		name     string
		skipCheck bool
	}{
		{expected: "5040", n: 7, name: "Testing7"},
		{expected: "6402373705728000", n: 18, name: "Testing18"},
		{expected: "15511210043330985984000000", n: 25, name: "Testing25"},
		{expected: "...", n: 10000, name: "Testing1000", skipCheck: true},
	}
	for _, testcase := range table {
		td := func(t *testing.T) {
			res := FactorialLarge(testcase.n)
			testcase.actual = StringArray(res).String()
			if testcase.skipCheck {
				t.Logf("%s\t FactorialLarge(%d) was\t\t:%s, test skipped", pass,testcase.n, testcase.actual)
				return
			}
			if testcase.expected != testcase.actual {
				t.Fatalf("%s\texpected\tFactorialLarge(%d)\tshould be:\t\t%s but it was:\t\t%s\n", fail,
					testcase.n, testcase.expected, testcase.actual)
			} else {
				t.Logf("%s\t is:\t\t:%s", pass, testcase.actual)
			}
		}
		t.Run(testcase.name, td)
	}

}

