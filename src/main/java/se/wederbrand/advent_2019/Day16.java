package se.wederbrand.advent_2019;

import java.util.LinkedList;

public class Day16 {

	public static String part1(String input, int iterations) {
		for (int i = 1; i <= iterations; i++) {
			input = phase(input);
		}

		return input.substring(0, 8);
	}

	private static String phase(String input) {
		String output = "";
		for (int i = 0; i < input.length(); i++) {
			output += getOne(input, i+1);
		}
		return output;
	}

	private static String getOne(String input, int phaseOrder) {
		LinkedList<Integer> basePattern = new LinkedList<>();
		basePattern.add(0);
		basePattern.add(1);
		basePattern.add(0);
		basePattern.add(-1);

		int[] pattern = new int[input.length()+1];
		int i = 0;
		while (i < pattern.length) {
			Integer pop = basePattern.pop();
			basePattern.add(pop);
			for (int j = 0; j < phaseOrder && i < pattern.length; j++) {
				pattern[i] = pop;
				i++;
			}
		}

		int result = 0;
		String[] split = input.split("");
		for (int i1 = 0; i1 < split.length; i1++) {
			String number = split[i1];
			int value = Integer.parseInt(number);
			value *= pattern[i1+1];
			result += value;
		}

		String s = String.valueOf(result);
		return s.substring(s.length()-1);
	}

	public static String part2(String input) {
		String actualInput = "";
		for (int i = 0; i < 10000; i++) {
			actualInput += input;
		}

		String output = part1(actualInput, 100);
		String offsetString = input.substring(0, 7);
		int offset = Integer.parseInt(offsetString);

		return output.substring(offset, offset+8);
	}
}
