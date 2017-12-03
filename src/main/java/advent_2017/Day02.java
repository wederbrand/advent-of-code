package advent_2017;

import java.util.Scanner;

public class Day02 {
	public int part1(String input) {
		Scanner scanner = new Scanner(input).useDelimiter("\t");
		int checkSum = 0;
		while (scanner.hasNextLine()) {
			String line = scanner.nextLine();
			String[] split = line.split("\\t");
			int min = Integer.MAX_VALUE;
			int max = Integer.MIN_VALUE;
			for (String value : split) {
				Integer intValue = Integer.valueOf(value);
				min = Integer.min(min, intValue);
				max = Integer.max(max, intValue);
			}

			checkSum += max - min;
		}

		return checkSum;
	}

	public int part2(String input) {
		Scanner scanner = new Scanner(input).useDelimiter("\t");
		int checkSum = 0;
		outerLoop:
		while (scanner.hasNextLine()) {
			String line = scanner.nextLine();
			String[] split = line.split("\\t");
			int[] splitInts = new int[split.length];
			for (int i = 0; i < split.length; i++) {
				splitInts[i] = Integer.valueOf(split[i]);
			}


			for (int i : splitInts) {
				for (int j : splitInts) {
					if (i == j) {
						continue;
					}
					if (i % j == 0) {
						checkSum += i / j;
						continue outerLoop;
					}
					if (j % i == 0) {
						checkSum += j / i;
						continue outerLoop;
					}
				}
			}
		}

		return checkSum;
	}


}
