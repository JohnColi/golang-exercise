package main

import "fmt"

// 練習 1：struct 內嵌（embedding）
//
// 目標：
// 1. 定義兩個 struct，讓其中一個內嵌另一個
// 2. 觀察「外層」能不能直接呼叫「內層」的方法（方法提升）
// 3. 想一想：這跟傳統 OOP 的繼承有什麼不同？

// TODO: 定義內層 struct，例如 Person / Animal / Engine ...
type Person struct {
	Name string
	Age  int
}
type Animal struct {
	Name string
	Age  int
}

// TODO: 為內層 struct 加上一個方法
func (_p *Person) SayHello() {
	fmt.Println("Hello, my name is", _p.Name)
}

// TODO: 定義外層 struct，內嵌上面那個 struct
type Student struct {
	Person
	School string
}

// TODO: （可選）為外層再加自己的方法
func (_s *Student) SayHello() {
	_s.Person.SayHello()
	fmt.Println("I am a student of", _s.School)
}

func runEmbeddingDemo() {
	fmt.Println("=== Embedding Demo ===")
	student := Student{}
	student.Name = "Johnny"
	student.Age = 20
	student.School = "Finger School"
	student.SayHello()
	// TODO: 建立外層實例，呼叫「內層提升上來」的方法並印出結果
}
