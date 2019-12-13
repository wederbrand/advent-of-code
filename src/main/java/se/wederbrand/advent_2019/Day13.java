package se.wederbrand.advent_2019;

import java.util.HashMap;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.TimeUnit;

import static java.lang.Math.abs;

public class Day13 {

	public static long part1(String code) throws InterruptedException {
		HashMap<String, Long> map = new HashMap<>();
		IntcodeComputer computer;

		ArrayBlockingQueue<Long> input = new ArrayBlockingQueue<>(1);
		ArrayBlockingQueue<Long> output = new ArrayBlockingQueue<>(100);

		computer = new IntcodeComputer("day 13", code, input, output);
		Thread thread = new Thread(computer);
		thread.start();

		while (thread.getState() != Thread.State.TERMINATED || !output.isEmpty()) {
			if (output.isEmpty()) {
				continue;
			}
			long x = output.take();
			long y = output.take();
			long tile = output.take();
			map.put(getKey(x, y), tile);
		}

		return numberOf(map, 2);
	}

	public static long part2(String code) throws InterruptedException {
		HashMap<String, Long> map = new HashMap<>();
		IntcodeComputer computer;

		ArrayBlockingQueue<Long> input = new ArrayBlockingQueue<>(100);
		ArrayBlockingQueue<Long> output = new ArrayBlockingQueue<>(100);

		computer = new IntcodeComputer("day 13", code, input, output);
		computer.setMemory(0, 2);
		Thread thread = new Thread(computer);
		thread.setName("computer");
		thread.start();

		long score = 0;
		long ball = -1;
		long paddle = -1;
		while (true) {
			Long x = output.poll(10, TimeUnit.MILLISECONDS);
			if (x == null && input.isEmpty() && ball != -1 && paddle != -1) {
				if (ball == paddle) {
					input.put(0L);
				}
				else if (ball < paddle) {
					input.put(-1L);
				}
				else {
					input.put(1L);
				}
				x = output.take();
			}
			long y = output.take();
			long tile = output.take();
			if (x == -1 && y == 0) {
				score = tile;
				long bricks = numberOf(map, 2);
				System.out.println(bricks + " " + score);
				if (bricks == 0) {
					return score;
				}
			}
			else {
				map.put(getKey(x, y), tile);
				if (tile == 4) {
					ball = x;
				}
				if (tile == 3) {
					paddle = x;
				}
			}
		}
	}

	private static String getKey(long x, long y) {
		return x + "," + y;
	}

	public static long numberOf(HashMap<String, Long> map, int i) {
		long count = 0;
		for (Long value : map.values()) {
			if (value == i) {
				count++;
			}
		}

		return count;
	}
}
