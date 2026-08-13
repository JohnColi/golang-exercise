package main

import "fmt"

// 練習 2：interface 與鴨子型別
//
// 目標：
// 1. 定義一個 interface（例如 Speaker / Shape / Writer ...）
// 2. 讓兩個以上不同的 struct 都「實作」它（不必寫 implements）
// 3. 寫一個函式，參數是 interface，分別傳入不同 struct 觀察行為

// TODO: 定義 interface
type Speaker interface {
	SayHello()
}
type Flyer interface {
	Fly()
}

// TODO: 定義第一個 struct，並實作 interface 的方法
type Duck struct {
	Name string
}

func (d Duck) SayHello() {
	fmt.Println("Hello, i'm a duck , ", d.Name)
}
func (d Duck) Fly() {
	fmt.Println("I'm flying")
}

// TODO: 定義第二個 struct，並實作同一組方法
type Bird struct {
	Name string
}

func (b Bird) SayHello() {
	fmt.Println("Hello, i'm a bird , ", b.Name)
}
func (b Bird) Fly() {
	fmt.Println("I'm flying")
}

// TODO: 寫一個接受 interface 的函式
func SayHello(speaker Speaker) {
	speaker.SayHello()
}

func Fly(flyer Flyer) {
	flyer.Fly()
}

func runInterfaceDemo() {
	// TODO: 建立兩個不同 struct 的實例，傳給同一個 interface 函式
	fmt.Println("=== Interface Demo ===")

	duck := Duck{}
	duck.Name = "Duck"
	duck.SayHello()
	duck.Fly()

	bird := Bird{}
	bird.Name = "Bird"
	bird.SayHello()
	bird.Fly()
}
