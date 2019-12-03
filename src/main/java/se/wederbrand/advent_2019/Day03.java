package se.wederbrand.advent_2019;

import java.util.HashMap;

public class Day03 {

	public long part1(String input) {
		HashMap<String, String> map = new HashMap<>();

		String[] wires = input.split(System.lineSeparator());
		String[] wire0 = wires[0].split((","));
		String[] wire1 = wires[1].split((","));

		int x = 0;
		int y = 0;

		map.put(getKey(x, y), ".");

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
				map.put(getKey(x, y), "1");
			}
		}

		x = 0;
		y = 0;
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

				if (map.getOrDefault(getKey(x, y), "").equals("1")) {
					map.put(getKey(x, y), "3");
					int distance = Math.abs(x) + Math.abs(y);
					if (distance < minDistance) {
						minDistance = distance;
					}
				} else {
					map.put(getKey(x, y), "3");
				}
			}

		}

		return minDistance;
	}

	private String getKey(int x, int y) {
		return x + "," + y;
	}

	public long part2(String input) {
		return 0;
	}

}
