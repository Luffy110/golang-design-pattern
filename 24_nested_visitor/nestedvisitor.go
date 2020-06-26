package nestedvisistor

import "fmt"

type Visitor interface {
	Visit(VisitorFunc) error
}

type VisitorFunc func() error

type Visitor1 struct {
}

func (v Visitor1) Visit(fn VisitorFunc) error {
	fmt.Println("start Visitor1")
	fn()
	fmt.Println("end Visitor1")

	return nil
}

type VisitorList []Visitor

func (vl VisitorList) Visit(fn VisitorFunc) error {
	for _, v := range vl {
		v.Visit(func() error {
			fmt.Println("start VisitorList")
			fn()
			fmt.Println("end VisitorList")
			return nil
		})
	}
	return nil
}

type Visitor2 struct {
	visitor Visitor
}

func (v Visitor2) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("start Visitor2")
		fn()
		fmt.Println("end Visitor2")
		return nil
	})
	return nil
}

type Visitor3 struct {
	visitor Visitor
}

func (v Visitor3) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("start Visitor3")
		fn()
		fmt.Println("end Visitor3")
		return nil
	})
	return nil
}
