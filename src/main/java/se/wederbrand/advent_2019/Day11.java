package se.wederbrand.advent_2019;

import javax.sound.midi.Soundbank;
import java.util.HashMap;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.TimeUnit;

public class Day11 {

	public static final int BLACK = 0;
	public static final int WHITE = 1;

	private static String getKey(int x, int y) {
		return x + "," + y;
	}

	public static int part1(String code) throws InterruptedException {
		ArrayBlockingQueue<Long> input = new ArrayBlockingQueue<>(100);
		ArrayBlockingQueue<Long> output = new ArrayBlockingQueue<>(100);

		IntcodeComputer computer = new IntcodeComputer("day 11", code, input, output);
		Thread thread = new Thread(computer);
		thread.start();
		int x = 0;
		int y = 0;
		int direction = 0; // 0 is up, 1 is right, 2 is downs and 3 is left
		HashMap<String, Integer> map = new HashMap<>();

		outer:
		while (thread.getState() != Thread.State.TERMINATED) {
			String key = getKey(x, y);
			int currentColour = map.getOrDefault(key, BLACK);
			input.put((long) currentColour);
			Long newColour = null;
			Long turn = null;
			while (newColour == null || turn == null) {
				if (thread.getState() == Thread.State.TERMINATED) {
					break outer;
				}
				// try with timeout until the thread terminates
					newColour = output.poll(10, TimeUnit.MILLISECONDS);
					turn = output.poll(10, TimeUnit.MILLISECONDS);
			}

			// paint
			map.put(key, Math.toIntExact(newColour));

			// turn
			if (turn == 0L) {
				direction--;
			}
			else {
				direction++;
			}

			direction = (direction + 4) % 4;

			// move
			switch (direction) {
				case 0: y--; break;
				case 1: x++; break;
				case 2: y++; break;
				case 3: x--; break;
			}
		}

		return map.size();
	}

	public static void part2(String code) throws InterruptedException {
		ArrayBlockingQueue<Long> input = new ArrayBlockingQueue<>(100);
		ArrayBlockingQueue<Long> output = new ArrayBlockingQueue<>(100);

		IntcodeComputer computer = new IntcodeComputer("day 11", code, input, output);
		Thread thread = new Thread(computer);
		thread.start();
		int x = 0;
		int y = 0;
		int direction = 0; // 0 is up, 1 is right, 2 is downs and 3 is left
		HashMap<String, Integer> map = new HashMap<>();
		map.put(getKey(0, 0), WHITE);

		outer:
		while (thread.getState() != Thread.State.TERMINATED) {
			String key = getKey(x, y);
			int currentColour = map.getOrDefault(key, BLACK);
			input.put((long) currentColour);
			Long newColour = null;
			Long turn = null;
			while (newColour == null || turn == null) {
				if (thread.getState() == Thread.State.TERMINATED) {
					break outer;
				}
				// try with timeout until the thread terminates
					newColour = output.poll(10, TimeUnit.MILLISECONDS);
					turn = output.poll(10, TimeUnit.MILLISECONDS);
			}

			// paint
			map.put(key, Math.toIntExact(newColour));

			// turn
			if (turn == 0L) {
				direction--;
			}
			else {
				direction++;
			}

			direction = (direction + 4) % 4;

			// move
			switch (direction) {
				case 0: y--; break;
				case 1: x++; break;
				case 2: y++; break;
				case 3: x--; break;
			}
		}

		for (int plotY = -1; plotY < 10; plotY++) {
			String row = "";
			for (int plotX = -1; plotX < 50; plotX++) {
				if (map.getOrDefault(getKey(plotX, plotY), -1) == WHITE) {
					row += "#";
				}
				else if (map.getOrDefault(getKey(plotX, plotY), -1) == BLACK) {
					row += ".";
				}
				else {
					row += " ";
				}
				map.remove(getKey(plotX, plotY));
			}
			System.out.println(row);
		}

		System.out.println(map.size());
	}
}
