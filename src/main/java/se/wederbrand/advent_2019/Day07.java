package se.wederbrand.advent_2019;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.*;

public class Day07 {

	public static List<int[]> getPermutations(int min, int max) {
		List<int[]> permutations = new ArrayList<>();

		for (int a = min; a <= max; a++) {
			for (int b = min; b <= max; b++) {
				if (b == a) {
					continue;
				}
				for (int c = min; c <= max; c++) {
					if (c == a || c == b) {
						continue;
					}
					for (int d = min; d <= max; d++) {
						if (d == a || d == b || d == c) {
							continue;
						}
						for (int e = min; e <= max; e++) {
							if (e == a || e == b || e == c || e == d) {
								continue;
							}
							int[] permutation = {a, b, c, d, e};
							permutations.add(permutation);
						}
					}
				}
			}
		}

		return permutations;
	}

	public static int machineOfLoopingMachines(int a, int b, int c, int d, int e, String input) throws InterruptedException {
		return 0;
	}

	public static int bestOfLoopingMachines(String input, int min, int max) throws InterruptedException {
		List<int[]> permutations = getPermutations(min, max);
		int maxResult = 0;
		for (int[] permutation : permutations) {
			int i = machineOfLoopingMachines(permutation[0], permutation[1], permutation[2], permutation[3], permutation[4], input);
			if (i > maxResult) {
				maxResult = i;
			}
		}

		return maxResult;
	}
}
