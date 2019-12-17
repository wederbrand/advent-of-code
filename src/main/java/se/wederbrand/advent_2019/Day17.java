package se.wederbrand.advent_2019;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.TimeUnit;

public class Day17 {
	IntcodeComputer computer;
	ArrayBlockingQueue<Long> input = new ArrayBlockingQueue<>(100);
	ArrayBlockingQueue<Long> output = new ArrayBlockingQueue<>(100);
	Thread thread;
	long uglyLastVal;

	HashMap<String, Character> map = new HashMap<>();

	public Day17(String code, boolean wakeUp) {
		computer = new IntcodeComputer("day 17", code, input, output);
		if (wakeUp) {
			computer.setMemory(0, 2);
		}
		thread = new Thread(computer);
		thread.setName("computer");
		thread.start();
	}

	public int part1() throws InterruptedException {
		String row = "";
		int x=-1;
		int y=0;
		while (thread.getState() != Thread.State.TERMINATED || !output.isEmpty()) {
			Long poll = output.poll(10, TimeUnit.MILLISECONDS);
			if (poll == null) {
				continue;
			}
			int output = Math.toIntExact(poll);
			x++;
			map.put(getKey(x, y), (char) (int) output);
			switch (output) {
				case 35:
					row += "#";
					break;
				case 46:
					row += ".";
					break;
				case 94:
					row += "^";
					break;
				case 60:
					row += "<";
					break;
				case 62:
					row += ">";
					break;
				case 118:
					row += "v";
					break;
				case 10:
					System.out.println(row);
					row = "";
					y++;
					x=-1;
					break;
				default:
					throw new RuntimeException("bad output");
			}
		}

		int count = 0;
		for (Map.Entry<String, Character> entry : map.entrySet()) {
			x = Integer.parseInt(entry.getKey().split(",")[0]);
			y = Integer.parseInt(entry.getKey().split(",")[1]);
			if (entry.getValue() != '#') {
				continue;
			}
			if (checkSurrounding(x, y)) {
				map.put(getKey(x, y), '0');
				count+=x*y;
			}
		}

		drawIt();
		return count;
	}

	private void drawIt() {
		for (int y = 0; y < 50; y++) {
			StringBuilder row = new StringBuilder();
			for (int x = 0; x < 50; x++) {
				Character orDefault = map.getOrDefault(getKey(x, y), ' ');
				if (orDefault == 10) {
					continue;
				}
				row.append(orDefault);
			}
			System.out.println(row.toString());
		}
	}

	private boolean checkSurrounding(int x, int y) {
		if (map.getOrDefault(getKey(x-1, y), ' ') != '#') return false;
		if (map.getOrDefault(getKey(x+1, y), ' ') != '#') return false;
		if (map.getOrDefault(getKey(x, y-1), ' ') != '#') return false;
		if (map.getOrDefault(getKey(x, y+1), ' ') != '#') return false;

		return true;
	}

	private String getKey(int x, int y) {
		return x + "," + y;
	}

	public Long part2(String main, String a, String b, String c) throws InterruptedException {
		waitAndPrint();

		for (char c1 : main.toCharArray()) {
			long asciiCode = c1;
			input.put(asciiCode);
		}
		input.put(10L);
		waitAndPrint();


		for (char c1 : a.toCharArray()) {
			long asciiCode = c1;
			input.put(asciiCode);
		}
		input.put(10L);
		waitAndPrint();

		for (char c1 : b.toCharArray()) {
			long asciiCode = c1;
			input.put(asciiCode);
		}
		input.put(10L);
		waitAndPrint();

		for (char c1 : c.toCharArray()) {
			long asciiCode = c1;
			input.put(asciiCode);
		}
		input.put(10L);
		waitAndPrint();
		input.put((long) 'n');
		input.put(10L);
		waitAndPrint();

		return uglyLastVal;
	}

	private void waitAndPrint() throws InterruptedException {
		while (true) {
			Long poll = output.poll(100, TimeUnit.MILLISECONDS);
			if (poll == null) break;
			uglyLastVal = poll;
			System.out.print((char) poll.longValue());
		}
	}
}
