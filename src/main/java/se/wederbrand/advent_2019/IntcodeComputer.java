package se.wederbrand.advent_2019;

import java.util.Arrays;
import java.util.HashMap;
import java.util.concurrent.BlockingQueue;

public class IntcodeComputer implements Runnable {
	private final String name;
	private HashMap<Integer, Long> memory;
	int relativeBase = 0;
	private BlockingQueue<Long> inputQueue;
	private BlockingQueue<Long> outputQueue;

	public IntcodeComputer(String name, String code, BlockingQueue<Long> inputQueue, BlockingQueue<Long> outputQueue) {
		this.name = name;
		this.memory = new HashMap<>();
		long[] longs = Arrays.stream(code.split(",")).mapToLong(Long::parseLong).toArray();
		for (int i = 0; i < longs.length; i++) {
			this.memory.put(i, longs[i]);
		}
		this.inputQueue = inputQueue;
		this.outputQueue = outputQueue;
	}

	@Override
	public void run() {
		int i = 0;
		while (true) {
			Long value = memory.get(i);
			int opCode = (int) (value % 100);
			if (opCode == 99) {
				return;
			}

			int c = (int) (value / 100 % 10);
			int b = (int) (value / 1000 % 10);
			int a = (int) (value / 10000 % 10);

			int paramPosition1 = getParamPosition(c, i + 1);
			Long param1 = getParam(c, i + 1);

			int paramPosition2 = getParamPosition(b, i + 2);
			Long param2 = getParam(b, i + 2);

			int paramPosition3 = getParamPosition(a, i + 3);

			switch (opCode) {
				case 1: // +
					memory.put(paramPosition3, param1 + param2);
					i += 4;
					break;
				case 2: // *
					memory.put(paramPosition3, param1 * param2);
					i += 4;
					break;
				case 3: // input
					try {
//						System.out.println(name + " taking from queue of size " + inputQueue.size());
						memory.put(Math.toIntExact(paramPosition1), inputQueue.take());
					}
					catch (InterruptedException e) {
						e.printStackTrace();
					}
					i += 2;
					break;
				case 4: // output
					try {
//							System.out.println(name + " posting in queue of size " + outputQueue.size());
						outputQueue.put(param1);
					}
					catch (InterruptedException e) {
						e.printStackTrace();
					}
					i += 2;
					break;
				case 5: // jump if true
					if (param1 != 0) {
						i = Math.toIntExact(param2);
					}
					else {
						i += 3;
					}
					break;
				case 6: // jump if false
					if (param1 == 0) {
						i = Math.toIntExact(param2);
					}
					else {
						i += 3;
					}
					break;
				case 7: // less than
					if (param1 < param2) {
						memory.put(paramPosition3, 1L);
					}
					else {
						memory.put(paramPosition3, 0L);
					}
					i += 4;
					break;
				case 8: // equals
					if (param1.equals(param2)) {
						memory.put(paramPosition3, 1L);
					}
					else {
						memory.put(paramPosition3, 0L);
					}
					i += 4;
					break;
				case 9: // relative base
					relativeBase += param1;
					i += 2;
					break;
			}
		}
	}

	private Long getParam(int mode, int key) {
		if (mode == 1) { // immediate mode
			return memory.getOrDefault(key, 0L);
		}
		int position = Math.toIntExact(memory.getOrDefault(key, 0L));
		if (mode == 2) { // relative mode
			position += relativeBase;
		}

		return memory.getOrDefault(position, 0L);
	}

	private int getParamPosition(int mode, int key) {
		if (mode == 1) { // immediate mode
			return -1;
		}

		int position = 0;
		try {
			position = Math.toIntExact(memory.getOrDefault(key, 0L));
		}
		catch (Exception e) {
			// ignore, it happens that we try to fetch values we won't use
		}
		if (mode == 2) { // relative mode
			position += relativeBase;
		}

		return position;
	}

	public long getMemory(int position) {
		return memory.get(position);
	}

	public void setMemory(int position, long value) {
		memory.put(position, value);
	}
}
