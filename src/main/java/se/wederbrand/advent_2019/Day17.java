package se.wederbrand.advent_2019;

import java.util.concurrent.ArrayBlockingQueue;

public class Day17 {
	IntcodeComputer computer;
	ArrayBlockingQueue<Long> input = new ArrayBlockingQueue<>(100);
	ArrayBlockingQueue<Long> output = new ArrayBlockingQueue<>(100);

	public Day17(String code) {
		computer = new IntcodeComputer("day 17", code, input, output);
		Thread thread = new Thread(computer);
		thread.start();
	}

	public int part1() throws InterruptedException {
		while (true) {
			int take = Math.toIntExact(output.take());
			String row = "";
			switch (take) {
				case 35: row +="#"; break;
				case 46: row +="."; break;
				case 10:
					System.out.printf(row);
					row = "";
					break;
			}
		}
	}



}
