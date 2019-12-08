package se.wederbrand.advent_2019;

import java.util.concurrent.ArrayBlockingQueue;

public class Day05 {
	public static int refactoredMachine(String input, int inputTo3) throws InterruptedException {
		ArrayBlockingQueue<Integer> machineInput = new ArrayBlockingQueue<>(100);
		ArrayBlockingQueue<Integer> machineOutput = new ArrayBlockingQueue<>(100);

		machineInput.put(inputTo3);

		IntcodeComputer intcodeComputer = new IntcodeComputer("day 5", input, machineInput, machineOutput);
		intcodeComputer.run();
		int last = 0;
		while (!machineOutput.isEmpty()) {
			last = machineOutput.poll();
		}
		return last;
	}
}
