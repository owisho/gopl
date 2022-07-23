package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func BytePopCount(x byte) int {
	return int(pc[x])
}

func DiffCount(sum1, sum2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		dif := sum1[i] ^ sum2[i]
		count += BytePopCount(dif)
	}
	return count
}
