package se.wederbrand.advent_2017;

public class Day09 {
	public int part1(String input) {
		// remove the ! and characters after
		String noNegotations = input.replaceAll("!.", "");

		// remove garbage
		String noGarbage = noNegotations.replaceAll("<.*?>", "");

		// count starting blocks
		int score = 0;
		int currentValue = 1;
		for (char c : noGarbage.toCharArray()) {
			switch (c) {
				case '{':
					score += currentValue;
					currentValue++;
					break;
				case '}':
					currentValue--;
					break;
			}

		}

		return score;
	}

	public int part2(String input) {
		// remove the ! and characters after
		String noNegotations = input.replaceAll("!.", "");

		// remove garbage but keep the <> for counting
		String noGarbage = noNegotations.replaceAll("<.*?>", "<>");

		return noNegotations.length() - noGarbage.length();
	}

}

