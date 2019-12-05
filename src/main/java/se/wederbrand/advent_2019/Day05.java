package se.wederbrand.advent_2019;

import java.util.Arrays;

public class Day05 {
	public long part1(String input, int inputTo3) {
		int[] ints = Arrays.stream(input.split(",")).mapToInt(Integer::parseInt).toArray();
		int i = 0;
		outer:
		while (true) {
			int opCode = ints[i] % 100;
			int c = (ints[i]) / 100 % 10;
			int b = (ints[i]) / 1000 % 10;
			int a = (ints[i]) / 10000 % 10;

			switch (opCode) {
				case 1: // +
				{
					int param1 = c == 0 ? ints[ints[i + 1]] : ints[i + 1];
					int param2 = b == 0 ? ints[ints[i + 2]] : ints[i + 2];
					ints[ints[i + 3]] = param1 + param2;
					i += 4;
				}

				break;
				case 2: // *
				{
					int param1 = c == 0 ? ints[ints[i + 1]] : ints[i + 1];
					int param2 = b == 0 ? ints[ints[i + 2]] : ints[i + 2];
					ints[ints[i + 3]] = param1 * param2;
					i += 4;
				}
				break;
				case 3: // input
					ints[ints[i + 1]] = inputTo3;
					i += 2;
					break;
				case 4: // output
					if (c == 0) {
						System.out.println(ints[ints[i + 1]]);
					} else {
						System.out.println(ints[i + 1]);
					}
					i += 2;
					break;
				case 99:
					break outer;
			}
		}
		return ints[0];
	}

	public long part2(String input) {
		return 0;
	}
}
