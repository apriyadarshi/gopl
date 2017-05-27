package ch4

//Shifts the item left by n positions
func rotL(s []int, n int) {
	rev(s[:n])
	rev(s[n:])
	rev(s[:])
}

//Rotate left in single pass
func rotL2(s []int, n int) []int {
	t := s[n:]
	for i := range s[0:n] {
		t = append(t, s[i])
	}
	return t
}
