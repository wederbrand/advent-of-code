package se.wederbrand.advent_2017;

public class Day19 {
	String[] lines;
	StringBuilder seenLetters;

	int y;
	int x;
	String direction;

	public String part1(String input) {
		lines = input.split(System.lineSeparator());
		seenLetters = new StringBuilder();

		y = 0;
		x = lines[y].indexOf("|");
		direction = "d";
		int steps = 0;

		while (true) {
			char c;
			try {
				c = lines[y].charAt(x);
				if (c == ' ') {
					break;
				}
			}
			catch (ArrayIndexOutOfBoundsException e) {
				break;
			}
			switch (c) {
				case '|':
				case '-':
					// simple path, doesn't change directions, just moves on
					break;
				case '+':
					// always a corner, turn
					turn();
					break;
				default:
					// letter
					seenLetters.append(c);
			}
			moveOn();
			steps++;
		}

		return seenLetters.toString() + " " + steps;
	}

	private void moveOn() {
		switch (direction) {
			case "d":
				y++;
				break;
			case "u":
				y--;
				break;
			case "l":
				x--;
				break;
			case "r":
				x++;
				break;
		}
	}

	private void turn() {
		switch (direction) {
			case "d":
			case "u":
				if (lines[y].charAt(x-1) != ' ') direction = "l";
				else if (lines[y].charAt(x+1) != ' ') direction = "r";
				break;
			case "l":
			case "r":
				if (lines[y-1].charAt(x) != ' ') direction = "u";
				else if (lines[y+1].charAt(x) != ' ') direction = "d";
				break;
		}
	}

}

