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
		int noun=0;
		int verb=0;
		int target = 19690720;
		while (part1(input, noun, verb) != target) {

		}
		return noun*100 + verb;
	}

}
