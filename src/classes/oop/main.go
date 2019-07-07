package main

import (
	"classes/oop/blog/website"
	"classes/oop/blog/author"
	"classes/oop/blog/post"
	"classes/oop/employee"
	"fmt"
)

func slice(args ...interface{}) []interface{} {
	return args
}
func main() {

	e1 := employee.New("TuTTa", "LuTTa", 28, 5)
	e1.PrintLeavesRemaining()
	fmt.Printf("%T %v", e1, e1)
	fmt.Println()
	a := author.New(
		"Viktor",
		"Bushmin",
		"Go dummer")

	p := post.New(
		append(
			[]string{
				"Inheritance in Go",
		    	"Go supports copmosition instead of inheritance"},
			a.GetAttributes()...
			)...
		)

	p1 := post.New(
		append(
			[]string{
				"Struct instead of Classes in Go",
		    	"Go does not support classes but methods can be added to structs"},
			a.GetAttributes()...
			)...
		)

	p2 := post.New(
		append(
			[]string{
				"Concurrency",
		    	"Go is a concurrent language and not a parallel one"},
			a.GetAttributes()...
			)...
		)
	w := website.NewFromPosts(p, p1, p2)
	w.Contents()
}
