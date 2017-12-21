package se.wederbrand.advent_2017;

import java.util.ArrayList;
import java.util.List;

public class Day21 {
	char[][] pattern = new char[3][3];

	public Day21() {
		for (int i = 0; i < 3; i++) {
			for (int j = 0; j < 3; j++) {
				pattern[i][j] = '.';
			}
		}
		pattern[0][1] = '#';
		pattern[1][2] = '#';
		pattern[2][0] = '#';
		pattern[2][1] = '#';
		pattern[2][2] = '#';
	}

	public int part1(int iterations, String input) {
		String[] inputPatterns = input.split(System.lineSeparator());
		List<Pattern> patterns = new ArrayList<>();
		for (String pattern : inputPatterns) {
			patterns.add(new Pattern(pattern));
		}

		for (int iteration = 0; iteration < iterations; iteration++) {
			if (pattern.length % 2 == 0) {
				// 2 -> 3
				char[][] newPattern = new char[3 * pattern.length / 2][3 * pattern.length / 2];
				for (int x = 0; x < pattern.length / 2; x++) {
					for (int y = 0; y < pattern.length / 2; y++) {
						for (Pattern matchPattern : patterns) {
							if (matchPattern.matches(extractPattern(pattern, x*2, y*2, 2))) {
								addPattern(newPattern, matchPattern, x*3, y*3);
								break;
							}
						}

					}
				}
				pattern = newPattern;
			}
			else if (pattern.length % 3 == 0) {
				// 3 -> 4
				char[][] newPattern = new char[4 * pattern.length / 3][4 * pattern.length / 3];
				for (int y = 0; y < pattern.length / 3; y++) {
					for (int x = 0; x < pattern.length / 3; x++) {
						for (Pattern matchPattern : patterns) {
							if (matchPattern.matches(extractPattern(pattern, x*3, y*3, 3))) {
								addPattern(newPattern, matchPattern, x*4, y*4);
								break;
							}
						}
					}
				}
				pattern = newPattern;
			}
			else {
				throw new RuntimeException("this should never happen");
			}
		}

		int count = 0;
		for (char[] chars : pattern) {
			for (char c : chars) {
				if (c == '#') {
					count++;
				}
			}
		}

		return count;
	}

	private void addPattern(char[][] newPattern, Pattern matchPattern, int x, int y) {
		String pattern = matchPattern.getReplacePattern();

		String[] split = pattern.split("/");
		for (int i = 0; i < split.length; i++) {
			String line = split[i];
			char[] charArray = line.toCharArray();
			System.arraycopy(charArray, 0, newPattern[x + i], y, charArray.length);
		}

	}

	private String extractPattern(char[][] pattern, int x, int y, int size) {
		StringBuilder result = new StringBuilder();
		for (int i = 0; i < size; i++) {
			for (int j = 0; j < size; j++) {
				result.append(pattern[i + x][j + y]);
			}
			result.append("/");
		}
		return result.substring(0,result.length()-1);
	}

	private class Pattern {
		final String replacePattern;
		final String[] matchPatterns = new String[8];

		public Pattern(String inputPattern) {
			replacePattern = inputPattern.split(" => ")[1];
			matchPatterns[0] = inputPattern.split(" => ")[0];

			if (matchPatterns[0].length() == 5) {
				// 2
				for (int i = 1; i < 4; i++) {
					rotate2(i);
				}
				for (int i = 0; i < 4; i++) {
					flip2(i);
				}

			}
			else {
				// 3
				for (int i = 1; i < 4; i++) {
					rotate3(i);
				}
				for (int i = 0; i < 4; i++) {
					flip3(i);
				}
			}

		}

		private void rotate2(int i) {
			String[] split = matchPatterns[i-1].split("/");

			matchPatterns[i] = "";

			matchPatterns[i] += split[0].charAt(1);
			matchPatterns[i] += split[1].charAt(1);
			matchPatterns[i] += "/";
			matchPatterns[i] += split[0].charAt(0);
			matchPatterns[i] += split[1].charAt(0);
		}

		private void flip2(int i) {
			matchPatterns[i+4] = "";

			matchPatterns[i+4] += matchPatterns[i].toCharArray()[1];
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[0];
			matchPatterns[i+4] += "/";
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[4];
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[3];
		}

		private void rotate3(int i) {
			String[] split = matchPatterns[i-1].split("/");

			matchPatterns[i] = "";

			matchPatterns[i] += split[0].charAt(2);
			matchPatterns[i] += split[1].charAt(2);
			matchPatterns[i] += split[2].charAt(2);
			matchPatterns[i] += "/";
			matchPatterns[i] += split[0].charAt(1);
			matchPatterns[i] += split[1].charAt(1);
			matchPatterns[i] += split[2].charAt(1);
			matchPatterns[i] += "/";
			matchPatterns[i] += split[0].charAt(0);
			matchPatterns[i] += split[1].charAt(0);
			matchPatterns[i] += split[2].charAt(0);
		}

		private void flip3(int i) {
			matchPatterns[i+4] = "";

			matchPatterns[i+4] += matchPatterns[i].toCharArray()[2];
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[1];
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[0];
			matchPatterns[i+4] += "/";
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[6];
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[5];
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[4];
			matchPatterns[i+4] += "/";
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[10];
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[9];
			matchPatterns[i+4] += matchPatterns[i].toCharArray()[8];
		}


		public boolean matches(String pattern) {
			for (String matchPattern : matchPatterns) {
				if (matchPattern.equals(pattern)) {
					return true;
				}
			}
			return false;
		}

		public String getReplacePattern() {
			return replacePattern;
		}
	}
}

