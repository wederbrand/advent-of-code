package se.wederbrand.advent_2019;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;
import java.util.concurrent.atomic.AtomicInteger;

public class Day08 {

	public int part1(String input, int width, int height) {
		List<HashMap<Integer, AtomicInteger>> counters = new ArrayList<>();
		Scanner scanner = new Scanner(input).useDelimiter("");
		while (scanner.hasNextInt()) {
			HashMap<Integer, AtomicInteger> hashMap = new HashMap<>();
			counters.add(hashMap);
			for (int i = 0; i < width * height; i++) {
				int value = scanner.nextInt();
				AtomicInteger atomicInteger = hashMap.computeIfAbsent(value, (v) -> new AtomicInteger(0));
				atomicInteger.incrementAndGet();
			}
		}

		int minZeroes = Integer.MAX_VALUE;
		int product = 0;
		for (HashMap<Integer, AtomicInteger> counter : counters) {
			AtomicInteger zeroes = counter.getOrDefault(0, new AtomicInteger(0));
			AtomicInteger ones = counter.getOrDefault(1, new AtomicInteger(0));
			AtomicInteger twos = counter.getOrDefault(2, new AtomicInteger(0));
			if (zeroes.get() < minZeroes) {
				minZeroes = zeroes.get();
				product = ones.get() * twos.get();
			}
		}

		return product;
	}

	public void part2(String input, int width, int height) {
		Integer[][] image = new Integer[height][width];

		Scanner scanner = new Scanner(input).useDelimiter("");
		while (scanner.hasNextInt()) {

			for (int i = 0; i < height; i++) {
				for (int j = 0; j < width; j++) {
					int value = scanner.nextInt();
					if (image[i][j] == null) {
						if (value == 0) {
							image[i][j] = 0;
						}
						if (value == 1) {
							image[i][j] = 1;
						}
					}
				}
			}
		}

		for (int h = 0; h < height; h++) {
			StringBuffer sb = new StringBuffer();
			for (int w = 0; w < width; w++) {
				if (image[h][w] == 0) {
					// don't print the black ones
					sb.append(" ");
				}
				else {
					sb.append("X");
				}
			}
			System.out.println(sb);
		}

	}


}
