//comments created by gemini to understand code after time

package basics_w_comments // Пакет называется так же, как и папка

import "fmt"

// Student — простая структура для примера
type Student struct {
	Name string
	Age  int
}

// Пример 1: Функция принимает КОПИЮ студента
// Изменения здесь НЕ повлияют на оригинал
func changeNameByValue(s Student) {
	s.Name = "Alex" // Меняем имя у копии
	fmt.Println("Внутри функции (по значению):", s.Name)
}

// Пример 2: Функция принимает УКАЗАТЕЛЬ (адрес в памяти)
// Изменения здесь ПОВЛИЯЮТ на оригинал
// *Student означает "указатель на Student"
func changeNameByPointer(s *Student) {
	s.Name = "Bob" // Go сам ходит по адресу и меняет оригинал
	fmt.Println("Внутри функции (по указателю):", s.Name)
}

// Функция для демонстрации (можно вызывать извне)
func LearnPointers() {
	// 1. Создаем студента
	student := Student{Name: "John", Age: 20}
	fmt.Println("Исходное имя:", student.Name)

	// 2. Передаем по значению (создается копия)
	changeNameByValue(student)
	fmt.Println("После changeNameByValue:", student.Name) // Осталось John!

	// 3. Передаем по указателю (через &)
	// &student означает "дай мне адрес ячейки памяти, где лежит student"
	changeNameByPointer(&student)
	fmt.Println("После changeNameByPointer:", student.Name) // Стало Bob!
}
