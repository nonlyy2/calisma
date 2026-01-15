package basics

import "fmt"

type Student struct {
	Name   string
	Age    int
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

func (me *Student) haveBirthday() {
	me.Age++
}

func main() {
	me := Student{
		Name:   "Assylkhan",
		Age:    21,
		Grades: []int{10, 20, 30},
	}

	fmt.Println(me.Age)
	me.haveBirthday()
	fmt.Println(me.Age)

	// res := me.getAverage()
	// fmt.Println("avg: ", res)
}
