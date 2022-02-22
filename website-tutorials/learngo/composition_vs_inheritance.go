/*
 * Go does not support inheritance, however, it does support composition.
 */
package main

import "fmt"

// 1. Composition by embedding structs
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.fullName()) // access embedded structs method b.author.fullName
	fmt.Println("Bio: ", b.bio)
}

// 2. Embedding a slice of structs
type website struct {
	blogPosts []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of the Website")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}

func main() {
	// 1. Composition by embedding structs
	author1 := author{
		"Andere",
		"Brian",
		"Engineer and Teacher",
	}

	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost1.details()

	// 2. Embedding a slice of structs
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}

	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}

	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
}
