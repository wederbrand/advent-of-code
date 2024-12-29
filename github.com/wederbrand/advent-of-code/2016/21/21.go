package main

import (
	"fmt"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inFile := GetFileContents("2016/21/input.txt", "\n")

	input := "abcdefgh"
	password := doIt([]byte(input), inFile)
	fmt.Println("Part 1: ", password, "in", time.Since(start))

	permutations := Permutations([]byte(input))
	for _, permutation := range permutations {
		password = doIt(permutation, inFile)
		if password == "fbgdceah" {
			fmt.Println("Part 2: ", string(permutation), "in", time.Since(start))
			break
		}
	}

}

func doIt(input []byte, inFile []string) string {
	password := slices.Clone(input)

	for _, line := range inFile {
		if strings.HasPrefix(line, "swap position") {
			var x, y int
			fmt.Sscanf(line, "swap position %d with position %d", &x, &y)
			password = swapPos(password, x, y)
		} else if strings.HasPrefix(line, "swap letter") {
			var x, y string
			fmt.Sscanf(line, "swap letter %s with letter %s", &x, &y)
			password = swapLetter(password, x, y)
		} else if strings.HasPrefix(line, "rotate left") {
			var x int
			fmt.Sscanf(line, "rotate left %d step", &x)
			password = rotateLeftRight(password, "left", x)
		} else if strings.HasPrefix(line, "rotate right") {
			var x int
			fmt.Sscanf(line, "rotate right %d step", &x)
			password = rotateLeftRight(password, "right", x)
		} else if strings.HasPrefix(line, "rotate based on position") {
			var x string
			fmt.Sscanf(line, "rotate based on position of letter %s", &x)
			password = rotateBasedOn(password, x)
		} else if strings.HasPrefix(line, "reverse") {
			var x, y int
			fmt.Sscanf(line, "reverse positions %d through %d", &x, &y)
			password = reverse(password, x, y)
		} else if strings.HasPrefix(line, "move") {
			var x, y int
			fmt.Sscanf(line, "move position %d to position %d", &x, &y)
			password = move(password, x, y)
		}
	}

	return string(password)
}

func swapPos(password []byte, x, y int) []byte {
	// swap position X with position Y means that the letters at indexes X and Y (counting from 0) should be swapped.
	password[x], password[y] = password[y], password[x]
	return password
}

func swapLetter(password []byte, x, y string) []byte {
	// swap letter X with letter Y means that the letters X and Y should be swapped (regardless of where they appear in the string).
	return swapPos(password, slices.Index(password, x[0]), slices.Index(password, y[0]))
}

func rotateLeftRight(password []byte, direction string, steps int) []byte {
	// rotate left/right X steps means that the whole string should be rotated; for example, one right rotation would turn abcd into dabc.
	if direction == "left" {
		steps *= -1
	}
	steps = ((steps % len(password)) + len(password)) % len(password)

	newPassword := make([]byte, len(password))
	for i, b := range password {
		newPassword[(i+steps+len(password))%len(password)] = b
	}

	return newPassword
}

func rotateBasedOn(password []byte, x string) []byte {
	// rotate based on position of letter X means that the whole string should be
	// rotated to the right based on the index of letter X (counting from 0) as
	// determined before this instruction does any rotations. Once the index is
	// determined, rotate the string to the right one time, plus a number of times
	// equal to that index, plus one additional time if the index was at least 4.
	steps := slices.Index(password, x[0])
	if steps >= 4 {
		steps += 1
	}
	steps += 1

	return rotateLeftRight(password, "right", steps)
}

func reverse(password []byte, x, y int) []byte {
	// reverse positions X through Y means that the span of letters at indexes X through Y (including the letters at X and Y) should be reversed in order.
	newPassword := make([]byte, 0)
	newPassword = append(newPassword, password[:x]...)
	bytes := slices.Clone(password[x : y+1])
	slices.Reverse(bytes)
	newPassword = append(newPassword, bytes...)
	newPassword = append(newPassword, password[y+1:]...)
	return newPassword
}

func move(password []byte, x, y int) []byte {
	// move position X to position Y means that the letter which is at index X should be removed from the string, then inserted such that it ends up at index Y.
	xValue := password[x]
	newPassword := slices.Clone(password)
	newPassword = slices.Delete(newPassword, x, x+1)
	newPassword = slices.Insert(newPassword, y, xValue)
	return newPassword
}
