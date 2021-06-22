package main

import "testing"

func TestCalculate(t *testing.T) {
	result := calculate(2)

	t.Logf("result %d", result)
	if result != 4 {
		t.Log("an error has occured")
		t.FailNow()
	}
	t.Log("This will be printed only with -v")
}
