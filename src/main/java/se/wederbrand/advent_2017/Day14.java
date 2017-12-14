package se.wederbrand.advent_2017;

public class Day14 {
	public int part1(String input) {
		Day10 day10 = new Day10();

		int count = 0;
		for (int i = 0; i < 128; i++) {
			String knotHash = day10.part2(256, input + "-" + i);
			for (char c : knotHash.toCharArray()) {
				int digit = Character.digit(c, 16);
				for (char binaryCharacter : Integer.toString(digit, 2).toCharArray()) {
					if (binaryCharacter == '1') {
						count++;
					}
				}

			}

		}
		return count;

	}
	public int part2(String input) {
		Day10 day10 = new Day10();

		int[][] map = new int[128][128];
		for (int i = 0; i < 128; i++) {
			String knotHash = day10.part2(256, input + "-" + i);
			int j=0;
			for (char c : knotHash.toCharArray()) {
				int digit = Character.digit(c, 16);
				for (char binaryCharacter : String.format("%4s", Integer.toBinaryString(digit)).replace(' ', '0').toCharArray()) {
					if (binaryCharacter == '1') {
						map[i][j] = 1;
					}
					j++;
				}
			}
		}

		// find each region
		int count = 0;

		for (int i = 0; i < map.length; i++) {
			int[] ints = map[i];
			for (int j = 0; j < ints.length; j++) {
				int square = ints[j];
				if (square == 1) {
					clearRegionRecursive(i,j,map);
					count++;
				}
			}
		}

		return count;

	}

	private void clearRegionRecursive(int i, int j, int[][] map) {
		if (i < 0 || i > 127 || j < 0 || j > 127) {
			return;
		}

		if (map[i][j] == 1) {
			map[i][j] = 0;
			clearRegionRecursive(i-1, j, map);
			clearRegionRecursive(i+1, j, map);
			clearRegionRecursive(i, j-1, map);
			clearRegionRecursive(i, j+1, map);
		}

	}
}

