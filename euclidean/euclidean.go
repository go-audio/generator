package euclidean

// Rhythm returns a rhythmical pattern of equally distributed accents throughout
// the total steps.
// The Euclidean rhythms are explained in this white paper:
// http://cgm.cs.mcgill.ca/~godfried/publications/banff.pdf
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
