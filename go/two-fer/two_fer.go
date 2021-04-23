// Package twofer is all about sharing and caring.
package twofer

import "fmt"

// ShareWith returns a fair split between given name and me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
