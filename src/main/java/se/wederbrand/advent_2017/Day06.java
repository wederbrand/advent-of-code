package se.wederbrand.advent_2017;

import java.util.HashMap;
import java.util.HashSet;

public class Day06 {
	public int part1(String input) {
		String[] memoryString = input.split("\\t");
		int[] memory = new int[memoryString.length];

		for (int i = 0, instructionsStringLength = memoryString.length; i < instructionsStringLength; i++) {
			String s = memoryString[i];
			memory[i] = Integer.valueOf(s);
		}

		HashSet<String> history = new HashSet<>();

		int steps = 0;

		while (!history.contains(toHash(memory))) {
			steps++;
			history.add(toHash(memory));
			memory = transform(memory);
		}

		return steps;
	}

	public int part2(String input) {
		String[] memoryString = input.split("\\t");
		int[] memory = new int[memoryString.length];

		for (int i = 0, instructionsStringLength = memoryString.length; i < instructionsStringLength; i++) {
			String s = memoryString[i];
			memory[i] = Integer.valueOf(s);
		}

		HashMap<String, Integer> history = new HashMap<>();

		int steps = 0;

		while (!history.containsKey(toHash(memory))) {
			steps++;
			history.put(toHash(memory), steps);
			memory = transform(memory);
		}

		return steps-history.get(toHash(memory)) + 1;
	}

	private int[] transform(int[] memory) {
		int maxIndex = 0;
		int max = Integer.MIN_VALUE;
		for (int i = 0; i < memory.length; i++) {
			if (memory[i] > max) {
				max = memory[i];
				maxIndex = i;
			}
		}

		int blocks = max;
		memory[maxIndex] = 0;

		while (blocks > 0) {
			maxIndex++;
			if (maxIndex >= memory.length) {
				maxIndex = 0;
			}
			memory[maxIndex]++;
			blocks--;
		}

		return memory;
	}

	private String toHash(int[] memory) {
		StringBuilder hash = new StringBuilder();
		for (int i : memory) {
			hash.append(i).append(".");
		}

		return hash.toString();
	}

}

