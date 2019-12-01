package se.wederbrand.advent_2019;

import java.util.Scanner;

public class Day01 {
	public long part1(String input) {
		Scanner scanner = new Scanner(input);
		long sum = 0;
		while (scanner.hasNextInt()) {
			sum+=getFuel(scanner.nextInt());
		}
		return sum;
	}

	private int getFuel(int mass) {
		return mass / 3 - 2;
	}

	public long part2(String input) {
		Scanner scanner = new Scanner(input);
		long sum = 0;
		while (scanner.hasNextInt()) {
			int fuel = getFuel(scanner.nextInt());
			while (fuel > 0) {
				sum+=fuel;
				fuel = getFuel(fuel);
			}
		}
		return sum;

	}

}
