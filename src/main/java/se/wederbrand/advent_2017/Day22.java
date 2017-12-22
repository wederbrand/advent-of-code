package se.wederbrand.advent_2017;

import java.util.HashMap;

public class Day22 {
	private HashMap<String, Character> map = new HashMap<>();
	private int x = 0;
	private int y = 0;
	private int dir = 0; // 0 is up, 1 is right, 2 is down and 3 is left

	public int part1(int iterations, String input) {
		createMap(input);

		int infectionCount = 0;
		for (int i = 0; i < iterations; i++) {
			String pos = x + ":" + y;
			switch (map.getOrDefault(pos, '.')) {
				case '.':
					dir = (dir + 3) % 4; // turn left
					map.put(pos, '#');
					infectionCount++;
					break;
				case '#':
					dir = (dir + 5) % 4; // turn right
					map.remove(pos);
					break;
			}

			moveOne();
		}

		return infectionCount;
	}

	private void moveOne() {
		switch (dir) {
			case 0:
				y--;
				break; // up
			case 1:
				x++;
				break; // right
			case 2:
				y++;
				break; // down
			case 3:
				x--;
				break; // left
		}
	}

	public int part2(int iterations, String input) {
		createMap(input);

		int infectionCount = 0;
		for (int i = 0; i < iterations; i++) {
			String pos = x + ":" + y;
			switch (map.getOrDefault(pos, '.')) {
				case '.':
					dir = (dir + 3) % 4; // turn left
					map.put(pos, 'W');
					break;
				case 'W':
					map.put(pos, '#');
					infectionCount++;
					break;
				case '#':
					dir = (dir + 5) % 4; // turn right
					map.put(pos, 'F');
					break;
				case 'F':
					dir = (dir + 2) % 4; // reverse
					map.remove(pos);
					break;
			}
			moveOne();
		}

		return infectionCount;
	}

	private void createMap(String input) {
		String[] split = input.split(System.lineSeparator());
		for (int iy = 0; iy < split.length; iy++) {
			String line = split[iy];
			x = line.length() / 2;
			y = iy / 2;
			char[] charArray = line.toCharArray();
			for (int ix = 0; ix < charArray.length; ix++) {
				char c = charArray[ix];
				if (c == '#') {
					map.put(ix + ":" + iy, '#');
				}
			}
		}
	}

}

