package se.wederbrand.advent_2019;

import java.net.InetAddress;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;
import java.util.concurrent.atomic.AtomicInteger;

public class Day08 {
	int WIDTH = 25;
	int HEIGHT = 6;

	public int part1(String input) {
		List<HashMap<Integer, AtomicInteger>> counters = new ArrayList<>();
		Scanner scanner = new Scanner(input).useDelimiter("");
		while (scanner.hasNextInt()) {
			HashMap<Integer, AtomicInteger> hashMap = new HashMap<>();
			counters.add(hashMap);
			for (int i = 0; i < WIDTH * HEIGHT; i++) {
				int value = scanner.nextInt();
				AtomicInteger atomicInteger = hashMap.computeIfAbsent(value, (v) -> new AtomicInteger(0));
				atomicInteger.incrementAndGet();
			}
		}

		int minZeroes = Integer.MAX_VALUE;
		int product = 0;
		for (HashMap<Integer, AtomicInteger> counter : counters) {
			if (counter.get(0).get() < minZeroes) {
				minZeroes = counter.get(0).get();
				product = counter.get(1).get() * counter.get(2).get();
			}
		}

		return product;
	}

	public int part2(String input) {
		return 0;
	}


}
