package se.wederbrand.advent_2019;

import java.util.Arrays;

public class Day02 {
	public long part1(String input, int noun, int verb) {
		int[] ints = Arrays.stream(input.split(",")).mapToInt(Integer::parseInt).toArray();
		int i = 0;
		ints[1] = noun;
		ints[2] = verb;
		outer:
		while (true) {
			switch (ints[i]) {
				case 1:
					ints[ints[i + 3]] = ints[ints[i + 1]] + ints[ints[i + 2]];
					break;
				case 2:
					ints[ints[i + 3]] = ints[ints[i + 1]] * ints[ints[i + 2]];
					break;
				case 99:
					break outer;
			}
			i += 4;
		}
		return ints[0];
	}

	public long part2(String input) {
		for (int i = 0; i < 99; i++) {
			for (int j = 0; j < 99; j++) {
				if (part1(input, i, j) == 19690720) {
					return i * 100 + j;
				}
			}
		}
		return 0;
	}
}
