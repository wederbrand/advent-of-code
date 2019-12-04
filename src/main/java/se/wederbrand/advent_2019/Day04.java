package se.wederbrand.advent_2019;

public class Day04 {
	public boolean part1internal(String s) {
		String[] split = s.split("");
		boolean doubleFound = false;
		int last = 0;
		for (String s1 : split) {
			int i = Integer.parseInt(s1);
			if (i == last) {
				doubleFound = true;
			}
			if (i >= last) {
				last = i;
			} else {
				return false;
			}
		}
		return doubleFound;
	}

	public boolean part2internal(String s) {
		String[] split = s.split("");
		boolean doubleFound = false;
		int last = 0;
		int inARow = 1;
		for (String s1 : split) {
			int i = Integer.parseInt(s1);
			if (i == last) {
				inARow++;
			}
			else if (i > last) {
				if (inARow == 2) {
					doubleFound = true;
				}
				last = i;
				inARow = 1;
			}
			else {
				return false;
			}
		}

		if (inARow == 2) {
			doubleFound = true;
		}

		return doubleFound;
	}

	public long part1(String input) {
		String[] split = input.split("-");
		int from = Integer.parseInt(split[0]);
		int to = Integer.parseInt(split[1]);
		int totalFound = 0;
		for (int i = from; i < to; i++) {
			if (part1internal(String.valueOf(i))) {
				totalFound++;
			}
		}

		return totalFound;
	}

	public long part2(String input) {
		String[] split = input.split("-");
		int from = Integer.parseInt(split[0]);
		int to = Integer.parseInt(split[1]);
		int totalFound = 0;
		for (int i = from; i < to; i++) {
			if (part2internal(String.valueOf(i))) {
				totalFound++;
			}
		}

		return totalFound;

	}

}
