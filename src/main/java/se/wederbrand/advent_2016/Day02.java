package se.wederbrand.advent_2016;

import java.util.Scanner;

public class Day02 {

	public String part1(String input) {
		Scanner scanner = new Scanner(input);

		int digit = 5;
		String code = "";

		while (scanner.hasNextLine()) {
			String line = scanner.nextLine();
			for (int i = 0; i < line.length(); i++) {
				char c = line.charAt(i);
				switch (c) {
					case 'U':
						if (digit > 3) {
							digit -= 3;
						}
						break;
					case 'D':
						if (digit < 7) {
							digit += 3;
						}
						break;
					case 'L':
						if (digit != 1 && digit != 4 && digit != 7) {
							digit -= 1;
						}
						break;
					case 'R':
						if (digit != 3 && digit != 6 && digit != 9) {
							digit += 1;
						}
						break;
				}
			}
			code += digit;
		}
		return code;
	}

	public String part2(String input) {
		Scanner scanner = new Scanner(input);

		int digit = 5;
		String code = "";

		while (scanner.hasNextLine()) {
			String line = scanner.nextLine();
			for (int i = 0; i < line.length(); i++) {
				char c = line.charAt(i);
				if (c == 'U') {
					if (digit == 3) digit = 1;
					else if (digit > 5 && digit < 9) digit -= 4;
					else if (digit > 9 && digit < 13) digit -= 4;
					else if (digit == 13) digit = 11;
				}
				else if (c == 'D') {
					if (digit == 1) digit = 3;
					else if (digit > 1 && digit < 5) digit += 4;
					else if (digit > 5 && digit < 9) digit += 4;
					else if (digit == 11) digit = 13;
				}
				else if (c == 'L') {
					if (digit > 2 && digit < 5) digit--;
					else if (digit > 5 && digit < 10) digit--;
					else if (digit > 10 && digit < 13) digit--;
				}
				else if (c == 'R') {
					if (digit > 1 && digit < 4) digit++;
					else if (digit > 4 && digit < 9) digit++;
					else if (digit > 9 && digit < 12) digit++;
				}
			}

			if (digit < 10) {
				code += digit;
			}
			else if (digit == 10) {
				code += "A";
			}
			else if (digit == 11) {
				code += "B";
			}
			else if (digit == 12) {
				code += "C";
			}
			else if (digit == 13) {
				code += "D";
			}

		}
		return code;
	}

}
