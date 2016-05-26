package main

import "testing"

func TestSum(t *testing.T){


	sum := func(arg0 int, arg1 int, arg2 int) int{
       return arg0+arg1+arg2
	}
	if sum(1,2,3)!=16{
		t.Log("我就是故意地...")
		t.FailNow()
	}
}


