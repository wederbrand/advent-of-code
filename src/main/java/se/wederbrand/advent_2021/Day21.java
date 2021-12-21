package se.wederbrand.advent_2021;

import java.util.*;

public class Day21 {
	HashMap<String, Long[]> cache;

	public Day21() {
		cache = new HashMap<>();
	}

	public static void main(String[] args) {
		Day21 day21 = new Day21();
		Long[] wins = day21.part2(new int[]{4, 9, 0, 0, 0});

		if (wins[0] > wins[1]) {
			System.out.println(wins[0]);
		}
		else {
			System.out.println(wins[1]);
		}
	}

	private Long[] part2(int[] state)  {
		Long[] wins = cache.get(key(state));
		if (wins != null) {
			return wins;
		} else {
			wins = new Long[2];
			wins[0] = 0L;
			wins[1] = 0L;
		}

		// do one player turn, then send it on recursively
		// for each turn spawn 27 universe (3^3 dice results)
		for (int d1 = 1; d1 <= 3; d1++) {
			for (int d2 = 1; d2 <= 3; d2++) {
				for (int d3 = 1; d3 <= 3; d3++) {
					int[] newState = Arrays.copyOf(state, 5);
					int totalDistance = d1 + d2 + d3;

					newState[state[4]] += totalDistance;
					if (newState[state[4]] > 10) {
						newState[state[4]] %= 10;
					}

					newState[state[4]+2] += newState[state[4]];
					if (newState[state[4]+2] >= 21) {
						wins[state[4]]++;
					} else {
						newState[4] += 1;
						newState[4] %= 2;
						Long[] newWins = part2(newState);
						wins[0] += newWins[0];
						wins[1] += newWins[1];
					}
				}
			}
		}

		cache.put(key(state), wins);
		return wins;
	}

	private String key(int[] state) {
		return state[0] + "," + state[1] + "," + state[2] + "," + state[3] + "," + state[4];
	}
}
