package se.wederbrand.advent_2019;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.*;

public class Day07 implements Runnable {

	private final String name;
	private String input;
	private BlockingQueue<Integer> inputQueue;
	private BlockingQueue<Integer> outputQueue;

	public Day07(String name, String input, BlockingQueue<Integer> inputQueue, BlockingQueue<Integer> outputQueue) {
		this.name = name;
		this.input = input;
		this.inputQueue = inputQueue;
		this.outputQueue = outputQueue;
	}

	public static List<int[]> getPermutations() {
		List<int[]> permutations = new ArrayList<>();

		for (int a = 0; a < 5; a++) {
			for (int b = 0; b < 5; b++) {
				if (b == a) {
					continue;
				}
				for (int c = 0; c < 5; c++) {
					if (c == a || c == b) {
						continue;
					}
					for (int d = 0; d < 5; d++) {
						if (d == a || d == b || d == c) {
							continue;
						}
						for (int e = 0; e < 5; e++) {
							if (e == a || e == b || e == c || e == d) {
								continue;
							}
							int[] permutation = {a, b, c, d, e};
							permutations.add(permutation);
						}
					}
				}
			}
		}

		return permutations;
	}

	public static int machineOfMachines(int a, int b, int c, int d, int e, String input) throws InterruptedException {
		ArrayBlockingQueue<Integer> aInput = new ArrayBlockingQueue<>(2);
		ArrayBlockingQueue<Integer> bInput = new ArrayBlockingQueue<>(2);
		ArrayBlockingQueue<Integer> cInput = new ArrayBlockingQueue<>(2);
		ArrayBlockingQueue<Integer> dInput = new ArrayBlockingQueue<>(2);
		ArrayBlockingQueue<Integer> eInput = new ArrayBlockingQueue<>(2);

		aInput.put(a);
		aInput.put(0);
		bInput.put(b);
		cInput.put(c);
		dInput.put(d);
		eInput.put(e);
		
		Day07 dayA = new Day07("a", input, aInput, bInput);
		Day07 dayB = new Day07("b", input, bInput, cInput);
		Day07 dayC = new Day07("c", input, cInput, dInput);
		Day07 dayD = new Day07("d", input, dInput, eInput);
		Day07 dayE = new Day07("e", input, eInput, aInput);

		new Thread(dayA).start();
		new Thread(dayB).start();
		new Thread(dayC).start();
		new Thread(dayD).start();
		Thread thread = new Thread(dayE);
		thread.start();
		thread.join();

		return aInput.take();
	}

	public int machineOfLoopingMachines(int a, int b, int c, int d, int e, String input) {
		return 0;
	}

	public static int bestOfMachines(String input) throws InterruptedException {
		List<int[]> permutations = getPermutations();
		int max = 0;
		for (int[] permutation : permutations) {
			int i = machineOfMachines(permutation[0], permutation[1], permutation[2], permutation[3], permutation[4], input);
			if (i > max) {
				max = i;
			}
		}

		return max;
	}

	@Override
	public void run() {
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
					try {
//						System.out.println(name + " taking from queue of size " + inputQueue.size());
						ints[ints[i + 1]] = inputQueue.take();
					}
					catch (InterruptedException e) {
						e.printStackTrace();
					}
					i += 2;
					break;
				case 4: // output
					if (c == 0) {
						try {
//							System.out.println(name + " posting in queue of size " + outputQueue.size());
							outputQueue.put(ints[ints[i + 1]]);
						}
						catch (InterruptedException e) {
							e.printStackTrace();
						}
					}
					else {
						try {
//							System.out.println(name + " posting in queue of size " + outputQueue.size());
							outputQueue.put(ints[i + 1]);
						}
						catch (InterruptedException e) {
							e.printStackTrace();
						}
					}
					i += 2;
					break;
				case 5: // jump if true
					if (param1 != 0) {
						i = param2;
					}
					else {
						i += 3;
					}
					break;
				case 6: // jump if false
					if (param1 == 0) {
						i = param2;
					}
					else {
						i += 3;
					}
					break;
				case 7: // less than
					if (param1 < param2) {
						ints[ints[i + 3]] = 1;
					}
					else {
						ints[ints[i + 3]] = 0;
					}
					i += 4;
					break;
				case 8: // equals
					if (param1 == param2) {
						ints[ints[i + 3]] = 1;
					}
					else {
						ints[ints[i + 3]] = 0;
					}
					i += 4;
					break;
				case 99:
					System.out.println(name +" EXIT");
					break outer;
			}
		}
	}
}
