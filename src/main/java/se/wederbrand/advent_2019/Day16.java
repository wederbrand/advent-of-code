package se.wederbrand.advent_2019;

import java.util.Arrays;
import java.util.LinkedList;

public class Day16 {

	public static String part1(String input, int iterations) {
		char[] chars = input.toCharArray();

		for (int i = 1; i <= iterations; i++) {
			System.out.println("iteration: " + i);
			chars = phase(chars);
		}

		return new String(chars).substring(0, 8);
	}

	private static char[] phase(char[] chars) {
		char[] output = new char[chars.length];
		for (int i = 0; i < chars.length; i++) {
			output[i] = getOne(chars, i+1);
		}
		return output;
	}

	private static char getOne(char[] input, int phaseOrder) {
		LinkedList<Integer> basePattern = new LinkedList<>();
		basePattern.add(0);
		basePattern.add(1);
		basePattern.add(0);
		basePattern.add(-1);

		int[] pattern = new int[input.length+1];
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
		for (int i1 = 0; i1 < input.length; i1++) {
			char c = input[i1];
			int value = c - 48;
			value *= pattern[i1+1];
			result += value;
		}

		return (char) ((Math.abs(result) % 10) + 48);
	}

	public static String part2(String input, int multiplier) {
		char[] inputChars = input.toCharArray();
		char[] actualInput = new char[input.length()* multiplier];
		for (int i = 0; i < multiplier; i++) {
			System.arraycopy(inputChars, 0, actualInput, i*input.length(), input.length());
		}

		String offsetString = input.substring(0, 7);
		int offset = Integer.parseInt(offsetString);
		char[] actualInputTruncated = Arrays.copyOfRange(actualInput, offset, actualInput.length);

		int[] result = new int[actualInputTruncated.length];
		for (int i = actualInputTruncated.length-1; i >= 0; i--) {
			char c = actualInputTruncated[i];
			int value = c - 48;
			result[result.length-1] = (result[result.length-1] + value) % 10;
		}

		for (int i = actualInputTruncated.length-2; i >= 0; i--) {
			char c = actualInputTruncated[i];
			int value = c - 48;
			result[i] = (result[i+1] - value + 10) % 10;
		}

		return new String(Arrays.copyOfRange(actualInput, 0, 8));
	}
}

