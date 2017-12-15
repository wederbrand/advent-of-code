package se.wederbrand.advent_2017;

public class Day15 {
	long A = 16807;
	long B = 48271;
	long DIV = 2147483647;

	public long part1(long a, long b) {
		long count = 0;
		for (int i = 0; i < 40000000; i++) {
			a = (a*A)%DIV;
			b = (b*B)%DIV;

			if ((a & 0xFFFF) == (b & 0xFFFF)) {
				count++;
			}

		}
		return count;
	}

	public long part2(long a, long b) {
		long count = 0;
		for (int i = 0; i < 5000000; i++) {
			do {
				a = (a*A)%DIV;
			} while (a % 4 != 0);
			do {
				b = (b*B)%DIV;
			} while (b % 8 != 0);

			if ((a & 0xFFFF) == (b & 0xFFFF)) {
				count++;
			}

		}
		return count;
	}

}

