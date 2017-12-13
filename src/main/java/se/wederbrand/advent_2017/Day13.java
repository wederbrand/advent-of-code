package se.wederbrand.advent_2017;

import java.util.Scanner;

public class Day13 {
	public int part1(String input) {
		return getPrice(input, 0);
	}

	public int part2(String input) {
		int delay = 0;
		while (true) {
			if (!isCaught(input, delay)) {
				return delay;
			}
			delay++;
			if(delay % 10000 == 0) {
				System.out.println(delay);
			}
		}
	}

	private int getPrice(String input, int delay) {
		int price = 0;

		Scanner scanner = new Scanner(input).useDelimiter("[\\s:\\n]+");
		while (scanner.hasNextInt()) {
			int depth = scanner.nextInt();
			int range = scanner.nextInt();

			if ((depth + delay) % ((range-1)*2) == 0) {
				price += depth*range;
			}
		}

		return price;
	}

	private boolean isCaught(String input, int delay) {
		Scanner scanner = new Scanner(input).useDelimiter("[\\s:\\n]+");
		while (scanner.hasNextInt()) {
			int depth = scanner.nextInt();
			int range = scanner.nextInt();

			if ((depth + delay) % ((range-1)*2) == 0) {
				return true;
			}
		}

		return false;
	}


}

