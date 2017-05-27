package ch4

//Eliminates/deletes adjacent duplicates in-place from a slice
//Assumption: [10 10 10] => [10]
//Done in 2 passes. Ist pass. Empty. second, remove empty
func remAdjDups(s []string) []string {
	for i := 0; i+1 < len(s); {
		if s[i] == s[i+1] {
			s[i], s[i+1] = "", ""
			i = i + 2
		} else {
			i = i + 1
		}
	}
	return nonEmpty(s)
}

//returns nonempty strings preserving order
func nonEmpty(s []string) []string {
	var i, j int
	for i, j = 0, 0; j < len(s); j++ {
		if s[j] != "" {
			s[i] = s[j]
			i++
		}
	}
	return s[:i]
}
