package models

import "fmt"

type Board struct{
	Id int64
}

func (b *Board) Foo() {
	fmt.Println("foo")
}
