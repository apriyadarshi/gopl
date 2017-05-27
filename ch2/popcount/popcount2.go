package popcount

//PopCount2 uses loops
func PopCount2(x uint64) int {
	var count int
	var i uint64
	for i = 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
