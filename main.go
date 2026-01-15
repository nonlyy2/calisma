package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (me Student) getAverage() float64 {
	nums := me.Grades
	total := 0
	for _, i := range nums {
		total += i
	}
	return float64(total) / float64(len(nums))
}

func main() {
	me := Student{
		Name:   "Assylkhan",
		Grades: []int{10, 20, 30},
	}
	res := me.getAverage()
	fmt.Println("avg: ", res)
}
