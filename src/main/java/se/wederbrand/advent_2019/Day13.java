package se.wederbrand.advent_2019;

import java.util.HashMap;
import java.util.concurrent.ArrayBlockingQueue;

import static java.lang.Math.abs;

public class Day13 {
	HashMap<String, Long> map = new HashMap<>();
	IntcodeComputer computer;

	public Day13(String code) throws InterruptedException {
		ArrayBlockingQueue<Long> input = new ArrayBlockingQueue<>(100);
		ArrayBlockingQueue<Long> output = new ArrayBlockingQueue<>(100);

		this.computer = new IntcodeComputer("day 13", code, input, output);
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
	}


	private String getKey(long x, long y) {
		return x + "," + y;
	}

	public long numberOf(int i) {
		long count = 0;
		for (Long value : map.values()) {
			if (value == i) {
				count++;
			}
		}

		return count;
	}
}
