// Euclidean package implements the Euclidean Algorithm to generate traditional
// musical rhythms. This popular rhythm approach was defined by Godfried
// Toussaint in 2005
// https://en.wikipedia.org/wiki/Euclidean_rhythm
package euclidean

// Rhythm returns a rhythmical pattern of equally distributed accents throughout
// the total steps. The Euclidean rhythms are explained in this white paper:
// http://cgm.cs.mcgill.ca/~godfried/publications/banff.pdf
// The total steps are the steps in a grid (or a circle) and the accents is the
// number of those steps you want to be triggered. The algorithm will position
// the accents (aka pulses) equally distributed across the available steps.
func Rhythm(accents, totalSteps int) []bool {
	var pattern []bool
	if totalSteps <= 0 {
		return []bool{}
	}
	if accents <= 0 {
		return make([]bool, totalSteps)
	}
	if accents > totalSteps {
		pattern = make([]bool, totalSteps)
		// we can't have more accent steps than total steps
		// so we just set all steps as accented
		for i := range pattern {
			pattern[i] = true
		}
		return pattern
	}

	a := make([][]bool, accents)
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		a[i] = []bool{true}
	}
	b := make([][]bool, totalSteps-accents)
	for i := 0; i < len(b); i++ {
		b[i] = []bool{false}
	}

	minLen := min(aLen, len(b))
	thresh := 0
	// Loop until len(a or b) > 2
	for minLen > thresh {
		// set the threshold to 1 after the first pass
		if thresh == 0 {
			thresh = 1
		}

		for i := 0; i < minLen; i++ {
			a[i] = append(a[i], b[i]...)
		}

		// if the b was the bigger array, only keep what we need
		if minLen == aLen {
			b = b[minLen:]
		} else {
			// update the smallest array with the remainders of a
			// and update a to include only the extended sub-arrays
			b = a[minLen:]
			a = a[:minLen]
		}
		aLen = len(a)
		minLen = min(aLen, len(b))
	}

	pattern = make([]bool, 0, totalSteps)
	for i := 0; i < len(a); i++ {
		pattern = append(pattern, a[i]...)
	}
	for i := 0; i < len(b); i++ {
		pattern = append(pattern, b[i]...)
	}

	return pattern
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
