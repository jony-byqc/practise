package main

import (
	"fmt"
)

type stu struct {
	Name string
	Age  int
}

func (p *stu) setName(name string) *stu {
	p.Name = name
	return p
}

func (p *stu) setAge(age int) *stu {
	p.Age = age
	return p
}

func (p *stu) print() {
	fmt.Printf("name:%s age:%d\n", p.Name, p.Age)
}

func main() {
	stu := stu{}
	stu.setName("stu01").setAge(18).print()
}
