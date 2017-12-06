package se.wederbrand.advent_2017;

public class Day05 {
	public int part1(String input) {
		String[] instructionsString = input.split("\\n");
		int[] instructions = new int[instructionsString.length];

		for (int i = 0, instructionsStringLength = instructionsString.length; i < instructionsStringLength; i++) {
			String s = instructionsString[i];
			instructions[i] = Integer.valueOf(s);
		}

		int i = 0;
		int steps = 0;

		while (i < instructions.length) {
			steps++;
			int j = i;
			i += instructions[i];
			instructions[j]++;
		}

		return steps;
	}

	public int part2(String input) {
		String[] instructionsString = input.split("\\n");
		int[] instructions = new int[instructionsString.length];

		for (int i = 0, instructionsStringLength = instructionsString.length; i < instructionsStringLength; i++) {
			String s = instructionsString[i];
			instructions[i] = Integer.valueOf(s);
		}

		int i = 0;
		int steps = 0;

		while (i < instructions.length) {
			steps++;
			int j = i;
			i += instructions[i];
			if (instructions[j] >= 3) {
				instructions[j]--;
			}
			else {
				instructions[j]++;
			}
		}

		return steps;
	}
}

