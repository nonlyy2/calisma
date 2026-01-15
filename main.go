package main

import "fmt"

type Student struct {
	Name       string
	Age        int
	Grades     []int
	IsStudying bool
}

func main() {
	me := Student{
		Name:       "Assylkhan",
		Age:        21,
		Grades:     []int{90, 85, 40},
		IsStudying: true,
	}

	fmt.Println("Студент:", me)
	fmt.Println("Имя:", me.Name)
	fmt.Println("Учится сейчас?", me.IsStudying)
}
