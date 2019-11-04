package main

import "testing"

func TestArithmeticServiceAdd(t *testing.T) {
	s := new(ArithmeticService)
	c := s.Add(1, 2)
	t.Log("total is ", c)
}
