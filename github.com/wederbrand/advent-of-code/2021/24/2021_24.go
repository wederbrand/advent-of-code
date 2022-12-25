package main

func main() {
}

func doIt() int {
	z := 0

	i1 := 0
	i2 := 0
	i3 := 0
	i4 := 0
	i5 := 0
	i6 := 0
	i7 := 0
	i8 := 0
	i9 := 0
	i10 := 0
	i11 := 0
	i12 := 0
	i13 := 0
	i14 := 0

	z = push(z, i1, 10, 12)
	z = push(z, i2, 12, 7)
	z = push(z, i3, 10, 8)
	z = push(z, i4, 12, 8)
	z = push(z, i5, 11, 15)
	z = pop(z, i6, -16, 12) // i15 + 15 == i6 + 16

	z = push(z, i7, 10, 8)
	z = pop(z, i8, -11, 13) // i7 + 8 == i8 + 11
	z = pop(z, i9, -13, 3)  // i4 + 8 == i9 + 13
	z = push(z, i10, 13, 13)
	z = pop(z, i11, -8, 3)   // i10 + 13 == i11 + 8
	z = pop(z, i12, -1, 9)   // i3 +  8 == i12 + 1
	z = pop(z, i13, -4, 4)   // i2 +  7 == u13 + 4
	z = pop(z, i14, -14, 13) // i1 + 12 == i14 + 14

	// manually solve part 1 and part 2 using the above rules

	return z
}

func push(z, w, xPlus, yPlus int) int {
	// push is never called with xPlus < 10 so w-xPlus is always negative
	// % is always positive
	//	if z%26 == w-xPlus {
	//	} else {
	z *= 26
	z += w + yPlus
	//	}

	return z
}

// this breaks if last "26" of z == w-xPlus
func pop(z, w, xPlus, yPlus int) int {
	z %= 26
	if z%26 == w-xPlus {
	} else {
		z *= 26
		z += w + yPlus
	}

	return z
}
