package nestedvisistor

// print result

//start Visitor1
//start VisitorList
//start Visitor2
//start Visitor3
//handle Visit
//end Visitor3
//end Visitor2
//end VisitorList
//end Visitor1
func ExampleNestedVisitor (){
	var visitor Visitor
	var visitors []Visitor
	visitor = Visitor1{}

	visitors = append(visitors, visitor)
	visitor = Visitor2{VisitorList(visitors)}
	visitor = Visitor3{visitor}

	visitor.Visit(func() error {
		fmt.Println("handle Visit")
		return nil
	})
}