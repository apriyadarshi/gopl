package popcount

func PopCount3(x uint64) int {
	count := 0
	var i uint64
	for i = 0; i < 65; i++ {
		count += int((byte(x >> i)) & 1)
	}
	return count
}
