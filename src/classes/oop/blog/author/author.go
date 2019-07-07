//Package author represents author blog class
package author

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

//Author exported type of private struct "author"
// firstName string
// 	lastName  string
// 	bio       string
type Author author

//New func creates author object
//firstName, lastName, bio
func New(initData ...string) Author {
	var a Author
outLabel:
	for _, v := range initData {
		switch {
		case len(a.firstName) == 0:
			a.firstName = v
		case len(a.lastName) == 0:
			a.lastName = v
		case len(a.bio) == 0:
			a.bio = v
		default:
			break outLabel
		}
	}

	return Author(a)
}

//FullName returns full name of an author in format %s %s
func (a Author) FullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

//GetBio returns bio of an author
func (a Author) GetBio() string {
	return a.bio
}

//GetAttributes returna firstName, secondName and bio
func (a Author) GetAttributes() []string {
	return append(make([]string, 0), a.firstName, a.lastName, a.bio)
}
