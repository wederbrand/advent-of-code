package main

func main() {
	// r2 is the instruction pointer

	r0 := 0
	r1 := 0
	r2 := 0
	r3 := 0
	r4 := 0
	r5 := 0

	// 0 - seti 123 0 3
	r3 = 123

	// 1 - bani 3 456 3
	r3 = r3 & 456

	// 2 - eqri 3 72 3
	if r3 == 72 {
		r3 = 1
	} else {
		r3 = 0
	}

	// 3 - addr 3 2 2
	r2 = r3 + r2 // would have restarted if bani didn't work

	// 4 - seti 0 0 2
	r2 = 0

	// 5 - seti 0 4 3
	r3 = 0

	// 6 - bori 3 65536 4
	r4 = r3 | 65536

	// 7 - seti 1107552 3 3
	r3 = 1107552

	// 8 - bani 4 255 5
	r5 = r4 & 255

	// 9 - addr 3 5 3
	r3 = r3 + r5

	// 10 - bani 3 16777215 3
	r3 = r3 & 16777215

	// 11 - muli 3 65899 3
	r3 = r3 * 65899

	// 12 - bani 3 16777215 3
	r3 = r3 & 16777215

	// 13 - gtir 256 4 5
	if 256 > r4 {
		r5 = 1
	} else {
		r5 = 0
	}

	// 14 - addr 5 2 2
	r2 = r5 + r2

	// 15 - addi 2 1 2
	r2 = r2 + 1

	// 16 - seti 27 0 2
	r2 = 27

	// 17 - seti 0 2 5
	r5 = 0

	// 18 - addi 5 1 1
	r1 = r5 + 1

	// 19 - muli 1 256 1
	r1 = r1 * 256

	// 20 - gtrr 1 4 1
	if r1 > r4 {
		r1 = 1
	} else {
		r1 = 0
	}

	// 21 - addr 1 2 2
	r2 = r1 + r2

	// 22 - addi 2 1 2
	r2 = r2 + 1

	// 23 - seti 25 3 2
	r2 = 25

	// 24 - addi 5 1 5
	r5 = r5 + 1

	// 25 - seti 17 3 2
	r2 = 17

	// 26 - setr 5 3 4
	r4 = r5

	// 27 - seti 7 4 2
	r2 = 7

	// 28 - eqrr 3 0 5
	if r3 == r0 {
		r5 = 1
	} else {
		r5 = 0
	}

	// 29 - addr 5 2 2
	r2 = r5 + r2

	// 30 - seti 5 8 2
	r2 = 5
}
