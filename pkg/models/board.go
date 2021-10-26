package models

import "fmt"

type Board struct{
	Id int64
}

func (b *Board) foo() {
	fmt.Println("foo")
}
