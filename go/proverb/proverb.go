// Package proverb has cool text generator.
package proverb

var line string = "For want of a %s the %s was lost."
var outro string = "And all for the want of a %s."

// Proverb returns some old wisdom.
func Proverb(rhyme []string) []string {
	proverb := make([]string, len(rhyme))

	if len(rhyme) > 0 {
		for i, word := range rhyme {
			if i == 0 {
				continue
			}
			proverb[i-1] = "For want of a " + rhyme[i-1] + " the " + word + " was lost."
		}

		proverb[len(rhyme)-1] = "And all for the want of a " + rhyme[0] + "."
	}
	return proverb
}
