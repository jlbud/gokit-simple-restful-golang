package main

type Service interface {
	Add(a, b int) int
}

type ArithmeticService struct{}

// 服务层
func (s *ArithmeticService) Add(a, b int) int {
	return a + b
}
