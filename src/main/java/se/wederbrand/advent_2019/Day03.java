package se.wederbrand.advent_2019;

public class Day03 {

	public long part1(String input) {
		String[][] map = new String[1000][1000];
		int offset = 500;

		String[] wires = input.split(System.lineSeparator());
		String[] wire0 = wires[0].split((","));
		String[] wire1 = wires[1].split((","));

		int x = offset;
		int y = offset;

		map[x][y] = ".";

		for (String line : wire0) {
			String direction = line.substring(0, 1);
			int length = Integer.parseInt(line.substring(1));

			for (int i = 0; i < length; i++) {
				switch (direction) {
					case "U":
						y++;
						break;
					case "D":
						y--;
						break;
					case "R":
						x++;
						break;
					case "L":
						x--;
						break;
				}
				map[x][y] = "1";
			}
		}

		x = offset;
		y = offset;
		int minDistance = Integer.MAX_VALUE;

		for (String line : wire1) {
			String direction = line.substring(0, 1);
			int length = Integer.parseInt(line.substring(1));

			for (int i = 0; i < length; i++) {
				switch (direction) {
					case "U":
						y++;
						break;
					case "D":
						y--;
						break;
					case "R":
						x++;
						break;
					case "L":
						x--;
						break;
				}

				if (map[x][y] !=  null && map[x][y].equals("1")) {
					map[x][y] = "3";
					int distance = x-offset + y - offset;
					if (distance < minDistance) {
						minDistance = distance;
					}
				} else {
					map[x][y] = "2";
				}
			}

		}

		return minDistance;
	}

	public long part2(String input) {
		return 0;
	}

}
