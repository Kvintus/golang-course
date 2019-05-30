package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Square struct {
	sideLength float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type Triangle struct {
	height float64
	base float64
}

func (t Triangle) getArea() float64 {
	return (t.height * t.base) / 2
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}



