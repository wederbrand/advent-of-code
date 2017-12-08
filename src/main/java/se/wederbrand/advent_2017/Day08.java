package se.wederbrand.advent_2017;

import java.util.HashMap;
import java.util.Scanner;

public class Day08 {
	HashMap<String, Integer> memory = new HashMap<>();

	public int part1(String input) {

		Scanner scanner = new Scanner(input);
		while (scanner.hasNextLine()) {
			String instruction = scanner.nextLine();
			String[] split = instruction.split(("\\s"));
			String target = split[0];
			String incdec = split[1];
			int value = Integer.parseInt(split[2]);
			Integer memoryTargetValue = getMemoryValue(split[4]);
			String condition = split[5];
			int conditionValue = Integer.parseInt(split[6]);

			Integer memoryValue = getMemoryValue(target);


			switch(condition) {
				case "<":
					if (memoryTargetValue < conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case "<=":
					if (memoryTargetValue <= conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case ">":
					if (memoryTargetValue > conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case ">=":
					if (memoryTargetValue >= conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case "!=":
					if (memoryTargetValue != conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case "==":
					if (memoryTargetValue == conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;

				default: throw new RuntimeException("dude, this never happens");
			}

			memory.put(target, memoryValue);

		}

		int max = Integer.MIN_VALUE;
		for (Integer integer : memory.values()) {
			if (integer > max) {
				max = integer;
			}
		}

		return max;
	}

	public int part2(String input) {
		int max = Integer.MIN_VALUE;

		Scanner scanner = new Scanner(input);
		while (scanner.hasNextLine()) {
			String instruction = scanner.nextLine();
			String[] split = instruction.split(("\\s"));
			String target = split[0];
			String incdec = split[1];
			int value = Integer.parseInt(split[2]);
			Integer memoryTargetValue = getMemoryValue(split[4]);
			String condition = split[5];
			int conditionValue = Integer.parseInt(split[6]);

			Integer memoryValue = getMemoryValue(target);


			switch(condition) {
				case "<":
					if (memoryTargetValue < conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case "<=":
					if (memoryTargetValue <= conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case ">":
					if (memoryTargetValue > conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case ">=":
					if (memoryTargetValue >= conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case "!=":
					if (memoryTargetValue != conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;
				case "==":
					if (memoryTargetValue == conditionValue) memoryValue = adjustValue(memoryValue, incdec, value);
					break;

				default: throw new RuntimeException("dude, this never happens");
			}

			memory.put(target, memoryValue);

			if (memoryValue > max) {
				max = memoryValue;
			}

		}

		return max;
	}

	private Integer adjustValue(Integer memoryTargetValue, String incdec, int value) {
		if (incdec.equals("inc")) {
			return memoryTargetValue + value;
		}
		else {
			return memoryTargetValue - value;
		}
	}

	private Integer getMemoryValue(String target) {
		Integer memoryValue = memory.get(target);
		if (memoryValue == null) {
			memoryValue = 0;
		}

		return memoryValue;
	}

}

