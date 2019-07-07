package website

import (
	"classes/oop/blog/post"
	"fmt"
)

type website struct {
	p []post.Post
}

//Website exports website
type Website website

//New creates website object
func New(args ...interface{}) Website {
	var w Website
	wArgs := make([]string, 0)
	for i, v := range args {
		wArgs = append(wArgs, fmt.Sprintf("%s", v))
		if (i != 0) && ((i % 4) == 0) {
			w.p = append(w.p, post.New(wArgs...))
			wArgs = make([]string, 0)
		}
	}
	return w
}

//NewFromPosts creates Website object from post.Post objects
func NewFromPosts(p ...post.Post) Website {
	var w Website
	for _, v := range p {
		w.p = append(w.p, v)
	}
	return w
}


//Contents prints blogs
func (w Website) Contents() {
	fmt.Println("Content of Website")
	fmt.Println()
	for _, v := range w.p {
		v.Details()
		fmt.Println()
	}
}
