package se.wederbrand.advent_2019;

import java.util.Arrays;
import java.util.concurrent.ArrayBlockingQueue;

public class Day02 {
	public static int refactoredMachine(String code, int noun, int verb) throws InterruptedException {
		ArrayBlockingQueue<Integer> input = new ArrayBlockingQueue<>(100);
		ArrayBlockingQueue<Integer> output = new ArrayBlockingQueue<>(100);

		IntcodeComputer intcodeComputer = new IntcodeComputer("day 2", code, input, output);
		intcodeComputer.setMemory(1, noun);
		intcodeComputer.setMemory(2, verb);

		intcodeComputer.run();

		return intcodeComputer.getMemory(0);
	}

	public long part2(String input) throws InterruptedException {
		for (int i = 0; i < 99; i++) {
			for (int j = 0; j < 99; j++) {
				if (refactoredMachine(input, i, j) == 19690720) {
					return i * 100 + j;
				}
			}
		}
		return 0;
	}
}
