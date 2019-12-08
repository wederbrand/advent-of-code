package se.wederbrand.advent_2019;

import java.util.Arrays;
import java.util.concurrent.ArrayBlockingQueue;

public class Day05 {
	public static int refactoredMachine(String input, int inputTo3) throws InterruptedException {
		ArrayBlockingQueue<Integer> aInput = new ArrayBlockingQueue<>(100);
		ArrayBlockingQueue<Integer> bInput = new ArrayBlockingQueue<>(100);

		aInput.put(inputTo3);

		IntcodeComputer intcodeComputer = new IntcodeComputer("day 5", input, aInput, bInput);
		intcodeComputer.run();
		int last = 0;
		while (!bInput.isEmpty()) {
			last = bInput.poll();
		}
		return last;
	}

	public void machine(String input, int inputTo3) {
		int[] ints = Arrays.stream(input.split(",")).mapToInt(Integer::parseInt).toArray();
		int i = 0;
		outer:
		while (true) {
			int opCode = ints[i] % 100;
			int c = (ints[i]) / 100 % 10;
			int b = (ints[i]) / 1000 % 10;
			int a = (ints[i]) / 10000 % 10;

			int param1 = 0;
			int param2 = 0;
			try {
				param1 = c == 0 ? ints[ints[i + 1]] : ints[i + 1];
				param2 = b == 0 ? ints[ints[i + 2]] : ints[i + 2];
			}
			catch (ArrayIndexOutOfBoundsException e) {
				// ignore, it happens on some instructions
			}

			switch (opCode) {
				case 1: // +
					ints[ints[i + 3]] = param1 + param2;
					i += 4;
				break;
				case 2: // *
					ints[ints[i + 3]] = param1 * param2;
					i += 4;
				break;
				case 3: // input
					ints[ints[i + 1]] = inputTo3;
					i += 2;
					break;
				case 4: // output
					if (c == 0) {
						System.out.println(ints[ints[i + 1]]);
					} else {
						System.out.println(ints[i + 1]);
					}
					i += 2;
					break;
				case 5: // jump if true
					if (param1 != 0) {
						i = param2;
					}
					else {
						i+=3;
					}
					break;
				case 6: // jump if false
					if (param1 == 0) {
						i = param2;
					}
					else {
						i+=3;
					}
					break;
				case 7: // less than
					if (param1 < param2) {
						ints[ints[i + 3]] = 1;
					}
					else {
						ints[ints[i + 3]] = 0;
					}
					i+=4;
					break;
				case 8: // equals
					if (param1 == param2) {
						ints[ints[i + 3]] = 1;
					}
					else {
						ints[ints[i + 3]] = 0;
					}
					i+=4;
					break;
				case 99:
					break outer;
			}
		}
	}
}
