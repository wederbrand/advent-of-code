package se.wederbrand.advent_2017;

import java.util.*;

public class Day12 {
	TreeMap<Integer, List<Integer>> map = new TreeMap<>();

	public int part1(String input) {
		createMap(input);
		return extractOne(0).size();
	}

	public int part2(String input) {
		createMap(input);

		int counter = 0;
		while (!map.isEmpty()) {
			counter++;
			Integer key = map.firstEntry().getKey();
			extractOne(key);
		}

		return counter;
	}

	private HashSet<Integer> extractOne(int start) {
		HashSet<Integer> group = new HashSet<>();
		group.add(start);

		outerloop:
		while (true) {
			for (Map.Entry<Integer, List<Integer>> entry : map.entrySet()) {
				if (group.contains(entry.getKey())) {
					group.addAll(entry.getValue());
					map.remove(entry.getKey());
					continue outerloop;
				}
			}

			break;
		}
		return group;
	}

	private void createMap(String input) {
		Scanner scanner = new Scanner(input);

		while (scanner.hasNextLine()) {
			String s = scanner.nextLine();
			String[] split = s.split("<->");
			String[] split1 = split[1].split(",");
			List<Integer> target = new ArrayList<>();
			for (String s1 : split1) {
				target.add(Integer.parseInt(s1.trim()));
			}
			map.put(Integer.parseInt(split[0].trim()), target);
		}
	}


}

