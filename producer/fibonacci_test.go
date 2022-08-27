package main

import "testing"


func failOnErrorT(t *testing.T, er error, message string) {
	if er != nil {
		t.Fatalf("\t %s: %s -> error: %s",fail, message, er)
	}
}
func  TestFibBingInt(t *testing.T) {
	var table = []struct{
		Name string
		Input string
		Output string 
		Base int
		SkipCheck bool
	}{
		{Name: "ten", Input: "10" , Output: "55", Base: 10, SkipCheck: false},
		{Name: "tenPow18", Input: "1000000000000000000" , Output: "-----", Base: 10, SkipCheck: true},
	}
	for _, input := range table {
		tfunc := func (tf *testing.T)  {
			// given
			t.Logf("\tgiven the need to compute the fib of a huge number %s \n", input.Input)
			{
				// when 
				t.Logf("when fib(%s, base %d) is computed using big integers approach\n", input.Input, input.Base) 
				{
					// should 
					result , err := FibBingInt(input.Input, input.Base)
					
					if input.SkipCheck {
						tf.Logf("\t%s Test: %s fib(%s, base %d) has result %s \n", pass,input.Name, input.Input, input.Base, result)
						return
					} else {
						failOnErrorT(tf,err, "failed to compute bigint factorial")
					}
					if result != input.Output  {
						tf.Logf("\t%s Test: %s fib(%s, base %d) has result %s but should be %s\n", fail,input.Name, input.Input, input.Base, result, input.Output)
					} else {
						tf.Logf("\t%s Test: %s fib(%s, base %d) has result %s \n", fail,input.Name, input.Input, input.Base, result)
					}
				}
			}
		}
		t.Run(input.Name,tfunc)
	}
}