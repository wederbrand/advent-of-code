package se.wederbrand.advent_2017;

import java.util.Dictionary;
import java.util.HashMap;

public class Day11 {
	HashMap<String, Integer> map = new HashMap<>();

	public int part1(String input) {
		String[] steps = input.split(",");

		map.put("n", 0);
		map.put("ne", 0);
		map.put("se", 0);
		map.put("s", 0);
		map.put("sw", 0);
		map.put("nw", 0);

		for (String step : steps) {
			switch (step) {
				case "n" : map.put("n", map.get("n")+1); break;
				case "ne": map.put("ne", map.get("ne")+1); break;
				case "se": map.put("se", map.get("se")+1); break;
				case "s" : map.put("s", map.get("s")+1); break;
				case "sw": map.put("sw", map.get("sw")+1); break;
				case "nw": map.put("nw", map.get("nw")+1); break;
			}
		}

		return step1();
	}

	public int part2(String input) {
		int max = 0;
		String[] steps = input.split(",");

		map.put("n", 0);
		map.put("ne", 0);
		map.put("se", 0);
		map.put("s", 0);
		map.put("sw", 0);
		map.put("nw", 0);

		for (String step : steps) {
			switch (step) {
				case "n" : map.put("n", map.get("n")+1); break;
				case "ne": map.put("ne", map.get("ne")+1); break;
				case "se": map.put("se", map.get("se")+1); break;
				case "s" : map.put("s", map.get("s")+1); break;
				case "sw": map.put("sw", map.get("sw")+1); break;
				case "nw": map.put("nw", map.get("nw")+1); break;
			}

			int distance = step1();
			if (distance > max) {
				max = distance;
			}
		}

		return max;
	}

	private int step1() {
		opposites("n", "s");
		opposites("ne", "sw");
		opposites("nw", "se");

		triplets("n", "ne", "se");
		triplets("ne", "se", "s");
		triplets("se", "s", "sw");
		triplets("s", "sw", "nw");
		triplets("sw", "nw", "n");
		triplets("nw", "n", "ne");

		int value = 0;
		for (Integer integer : map.values()) {
			value += integer;
		}
		return value;
	}

	private void triplets(String one, String two, String three) {
		while (map.get(one) > 0 && map.get(three) > 0) {
			map.put(one, map.get(one)-1);
			map.put(three, map.get(three)-1);
			map.put(two, map.get(two)+1);
		}
	}

	private void opposites(String one, String two) {
		if (map.get(one) > map.get(two)) {
			map.put(one, map.get(one) - map.get(two));
			map.put(two, 0);
		}
		else if (map.get(one) < map.get(two)) {
			map.put(two, map.get(two) - map.get(one));
			map.put(one, 0);
		}
		else {
			map.put(one, 0);
			map.put(two, 0);
		}
	}

}

