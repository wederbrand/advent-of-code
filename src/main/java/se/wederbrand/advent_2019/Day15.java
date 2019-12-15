package se.wederbrand.advent_2019;

import java.util.HashMap;
import java.util.concurrent.ArrayBlockingQueue;

public class Day15 {
	public static final int WALL = -1;
	private static final long N = 1;
	private static final long S = 2;
	private static final long W = 3;
	private static final long E = 4;

	IntcodeComputer computer;
	ArrayBlockingQueue<Long> input = new ArrayBlockingQueue<>(100);
	ArrayBlockingQueue<Long> output = new ArrayBlockingQueue<>(100);

	HashMap<String, Integer> map = new HashMap<>();
	Integer oxygenX = null;
	Integer oxygenY = null;

	public Day15(String code) throws InterruptedException {
		computer = new IntcodeComputer("day 15", code, input, output);
		Thread thread = new Thread(computer);
		thread.start();
	}

	public long part1() throws InterruptedException {
		// 0 distance travelled
		recursive(21, 21, 0, N, false);
		System.out.println(oxygenX + " " + oxygenY);
		return map.get(getKey(oxygenX, oxygenY));
	}

	public long part2() throws InterruptedException {
		// drive the bot to the oxygen (37, 33)
		try {
			recursive(21, 21, 0, N, true);
		}
		catch (RuntimeException e) {
			// all good
		}

		// 0 distance travelled
		map.clear();
		recursive(oxygenX, oxygenY, 0, N, false);
		int max = 0;
		for (Integer value : map.values()) {
			if (value > max) {
				max = value;
			}
		}
		return max;
	}

	private void recursive(int x, int y, int distance, long lastDirection, boolean exitAtOxygen) throws InterruptedException {
		boolean alreadyVisited = map.get(getKey(x, y)) != null;
		Integer bestDistance = map.getOrDefault(getKey(x, y), Integer.MAX_VALUE);
		if (distance < bestDistance) {
			map.put(getKey(x, y), distance);
//			System.out.println(x + " " + y + " " + distance);
		}

		if (alreadyVisited) {
			backtrack(lastDirection);
			return;
		}

//		drawMap();

		for (long i = 1; i <= 4; i++) {
			input.offer(i);
			long take = output.take();
			if (take == 0) {
				map.put(getKey(getX(x, i), getY(y, i)), WALL);
			}
			else if (take == 1) {
				// moved
				recursive(getX(x, i), getY(y, i), distance+1, i, exitAtOxygen);
			}
			else if (take == 2) {
				// found oxygen
				oxygenX = getX(x, i);
				oxygenY = getY(y, i);
				if (exitAtOxygen) {
					throw new RuntimeException("found");
				}
				// moved
				recursive(getX(x, i), getY(y, i), distance+1, i, exitAtOxygen);
			}
		}
		if (distance == 0) {
			return;
		}
		backtrack(lastDirection);
	}


	private void backtrack(long lastDirection) throws InterruptedException {
		// backtrack
		input.offer(getBackTrack(lastDirection));
		if (output.take() == 0) {
			throw new RuntimeException("illegal move back");
		}
	}

	private long getBackTrack(long lastDirection) {
		long moveBack;
		if (lastDirection == N) {
			moveBack = S;
		}
		else if (lastDirection == S) {
			moveBack = N;
		}
		else if (lastDirection == W) {
			moveBack = E;
		}
		else{
			moveBack = W;
		}
		return moveBack;
	}

	private void drawMap() {
		for (int y = 0; y < 40; y++) {
			String row = "";
			for (int x = 0; x < 40; x++) {
				String key = getKey(x, y);
				Integer value = map.get(key);
				if (value == null) {
					row += " ";
				}
				else if (value == WALL) {
					row += "X";
				}
				else {
					row += value<=9?value:9;
				}
			}
			System.out.println(row);
		}
	}

	private int getX(int x, long i) {
		switch ((int) i) {
			case (int) N: return x;
			case (int) S: return x;
			case (int) W: return x+1;
			case (int) E: return x-1;
		}
		throw new RuntimeException("no");
	}

	private int getY(int y, long i) {
		switch ((int) i) {
			case (int) N: return y-1;
			case (int) S: return y+1;
			case (int) W: return y;
			case (int) E: return y;
		}
		throw new RuntimeException("no");
	}

	private String getKey(int x, int y) {
		return x + "," + y;
	}

}
