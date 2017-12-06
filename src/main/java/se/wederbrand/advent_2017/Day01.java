package se.wederbrand.advent_2017;

import java.util.Scanner;

public class Day01 {
	public int part1(String input) {
		Scanner scanner = new Scanner(input).useDelimiter("");
		int sum = 0;
		int lastVal = ((int) input.charAt(input.length()-1))-48;
		while (scanner.hasNextInt()) {
			int value = scanner.nextInt();
			if (value == lastVal) {
				sum += value;
			}
			lastVal = value;
		}

		return sum;
	}

	public int part2(String input) {
		String firstHalf = input.substring(0, input.length() / 2);
		String secondHalf = input.substring(input.length() / 2);
		Scanner scanner1 = new Scanner(firstHalf).useDelimiter("");
		Scanner scanner2 = new Scanner(secondHalf).useDelimiter("");
		int sum = 0;

		while (scanner1.hasNextInt()) {
			int i = scanner1.nextInt();
			int j = scanner2.nextInt();
			if (i == j) {
				sum += i+j;
			}
		}

		return sum;
	}

}
