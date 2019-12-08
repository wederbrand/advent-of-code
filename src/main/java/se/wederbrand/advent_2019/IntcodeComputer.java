package se.wederbrand.advent_2019;

import java.util.Arrays;
import java.util.concurrent.BlockingQueue;

public class IntcodeComputer implements Runnable {
	private final String name;
	private int[] memory;
	private BlockingQueue<Integer> inputQueue;
	private BlockingQueue<Integer> outputQueue;

	public IntcodeComputer(String name, String code, BlockingQueue<Integer> inputQueue, BlockingQueue<Integer> outputQueue) {
		this.name = name;
		this.memory = Arrays.stream(code.split(",")).mapToInt(Integer::parseInt).toArray();
		this.inputQueue = inputQueue;
		this.outputQueue = outputQueue;
	}

	@Override
	public void run() {
		int i = 0;
		outer:
		while (true) {
			int opCode = memory[i] % 100;
			int c = (memory[i]) / 100 % 10;
			int b = (memory[i]) / 1000 % 10;
			int a = (memory[i]) / 10000 % 10;

			int param1 = 0;
			int param2 = 0;
			try {
				param1 = c == 0 ? memory[memory[i + 1]] : memory[i + 1];
				param2 = b == 0 ? memory[memory[i + 2]] : memory[i + 2];
			}
			catch (ArrayIndexOutOfBoundsException e) {
				// ignore, it happens on some instructions
			}

			switch (opCode) {
				case 1: // +
					memory[memory[i + 3]] = param1 + param2;
					i += 4;
					break;
				case 2: // *
					memory[memory[i + 3]] = param1 * param2;
					i += 4;
					break;
				case 3: // input
					try {
//						System.out.println(name + " taking from queue of size " + inputQueue.size());
						memory[memory[i + 1]] = inputQueue.take();
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
							outputQueue.put(memory[memory[i + 1]]);
						}
						catch (InterruptedException e) {
							e.printStackTrace();
						}
					}
					else {
						try {
//							System.out.println(name + " posting in queue of size " + outputQueue.size());
							outputQueue.put(memory[i + 1]);
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
						memory[memory[i + 3]] = 1;
					}
					else {
						memory[memory[i + 3]] = 0;
					}
					i += 4;
					break;
				case 8: // equals
					if (param1 == param2) {
						memory[memory[i + 3]] = 1;
					}
					else {
						memory[memory[i + 3]] = 0;
					}
					i += 4;
					break;
				case 99:
					break outer;
			}
		}
	}

	public int getMemory(int position) {
		return memory[position];
	}

	public void setMemory(int position, int value) {
		memory[position] = value;
	}
}
