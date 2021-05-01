// Package protein has methods for indentifying proteins from RNA and codons.
package protein

import "errors"

// Define custom errors
var (
	ErrStop        error = errors.New("stop")
	ErrInvalidBase error = errors.New("please dont")
)

// FromCodon returns protein matching given codon.
// Returns ErrStop if terminating codon met.
// Returns ErrInvalidBase if codon can't be identified.
func FromCodon(c string) (string, error) {
	switch c {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA", "STOP":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA returns proteins parsed from RNA strand.
// If unidentified codons met, returns ErrInvalidBase.
func FromRNA(c string) ([]string, error) {
	proteins := make([]string, 0, len(c)/3)
	for i := 0; i < len(c)-2; i += 3 {
		protein, err := FromCodon(c[i : i+3])
		if err == ErrStop {
			return proteins, nil
		}
		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
