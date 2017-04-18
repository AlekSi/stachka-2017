package popcnt

func popcnt(x uint64) int {
	var res uint64
	for ; x > 0; x >>= 1 {
		res += x & 1
	}
	return int(res)
}

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt2(x uint64) int {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return int((x * h01) >> 56)
}
