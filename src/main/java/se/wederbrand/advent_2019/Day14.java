package se.wederbrand.advent_2019;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.atomic.AtomicLong;

public class Day14 {
	HashMap<String, Formula> formulas = new HashMap<>();
	HashMap<String, AtomicLong> storage = new HashMap<>();

	public Day14(String input) {
		String[] lines = input.split(System.lineSeparator());
		for (String line : lines) {
			Formula formula = new Formula(line);
			formulas.put(formula.getOutput(), formula);
		}
	}

	public long part1() {
		storage.put("FUEL", new AtomicLong(-1));

		fixStorage();

		return storage.get("ORE").get()*-1;
	}

	public long part2() {
		long min = 1;
		long max = 1_000_000_000;

		long lastMid=-1;
		while (true) {
			long mid = (min + max) / 2;
			if (mid == lastMid) {
				return mid;
			}
			lastMid = mid;
			storage.clear();
			storage.put("FUEL", new AtomicLong(-mid));
			fixStorage();
			if (storage.get("ORE").get() < -1000000000000L) {
				// to much fuel
				max = mid;
			}
			else {
				// too little fuel
				min = mid;
			}
		}
	}

	private void fixStorage() {
		while (true) {
			Map.Entry<String, AtomicLong> nextNotOre = getNextNotOre();
			if (nextNotOre == null) break;
			Formula formula = formulas.get(nextNotOre.getKey());
			double numberToGet = nextNotOre.getValue().get() * -1.0;
			long multiplesToGet = (long) Math.ceil(numberToGet / formula.outputQuantity);
			for (Map.Entry<String, Long> entry : formula.getInput().entrySet()) {
				AtomicLong atomicLong = storage.computeIfAbsent(entry.getKey(), s -> new AtomicLong(0));
				long delta = -entry.getValue() * multiplesToGet;
				atomicLong.getAndAdd(delta);
			}
			AtomicLong atomicLong = storage.get(nextNotOre.getKey());
			long delta = formula.outputQuantity * multiplesToGet;
			atomicLong.getAndAdd(delta);
		}
	}

	private Map.Entry<String, AtomicLong> getNextNotOre() {
		for (Map.Entry<String, AtomicLong> entry : storage.entrySet()) {
			if (entry.getKey().equals("ORE")) continue;
			if (entry.getValue().get() < 0) {
				return entry;
			}
		}
		return null;
	}

	private class Formula {
		String output;
		long outputQuantity;
		HashMap<String, Long> input = new HashMap<>();

		public Formula(String line) {
			String[] split = line.split("=>");
			String[] outputString = split[1].trim().split((" "));
			output = outputString[1].trim();
			outputQuantity = Long.parseLong(outputString[0].trim());

			String[] inputStrings = split[0].split(",");
			for (String inputString : inputStrings) {
				String[] inputStringSplit = inputString.trim().split(" ");
				input.put(inputStringSplit[1].trim(), Long.valueOf(inputStringSplit[0].trim()));
			}
		}

		public String getOutput() {
			return output;
		}

		public HashMap<String, Long> getInput() {
			return input;
		}

		public long getOutputQuantity() {
			return outputQuantity;
		}
	}
}
