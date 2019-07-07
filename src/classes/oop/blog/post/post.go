// Package post represents post class
package post

import (
	"classes/oop/blog/author"
	"fmt"
)

type post struct {
	title   string
	content string
	author.Author
}

//Post is exported "post" class
// title string
// 	content string
// 	author.Author
// firstName string
// lastName  string
// bio       string
type Post post

//New creates post object
//Set titile, content, author first name, last name and bio
func New(initData ...string) Post {
	var p Post
	authorSlice := make([]string, 1)
outLabel:
	for i, v := range initData {
		switch i {
		case 0:
			p.title = v
		case 1:
			p.content = v
		case 2, 3, 4:
			authorSlice = append(authorSlice, v)
		default:
			break outLabel
		}
	}
	p.Author = author.New(authorSlice...)
	return p
}

//Details prints post details:
// title
// content
// author:
// 	firstName
// 	lastName
// 	bio
func (p Post) Details() {
	fmt.Println("Title :", p.title)
	fmt.Println("Content :", p.content)
	fmt.Println("Author :", p.FullName())
	fmt.Println("Bio :", p.GetBio())
}
