package popcount

func PopCount4(x uint64) int {
	var count int
	for {
		if x == 0 {
			break
		} else {
			x = x & (x - 1) //This takes x to the next non-zero bit
			count++
		}
	}
	return count
}

func PopCountInt8(x uint8) int {
	var count int
	for {
		if x == 0 {
			break
		} else {
			x = x & (x - 1) //This takes x to the next non-zero bit
			count++
		}
	}
	return count
}

func PopCountUInt(x uint) int {
	var count int
	for {
		if x == 0 {
			break
		} else {
			x = x & (x - 1) //This takes x to the next non-zero bit
			count++
		}
	}
	return count
}
