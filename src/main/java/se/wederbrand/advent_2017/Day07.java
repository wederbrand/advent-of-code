package se.wederbrand.advent_2017;

import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class Day07 {
	public String part1(String input) {
		HashMap<String, Program> programs = getTree(input);
		return getBottomProgram(programs).getName();
	}

	public int part2(String input) {
		HashMap<String, Program> programs = getTree(input);
		Program bottomProgram = getBottomProgram(programs);

		return bottomProgram.getCorrectWeight(-1);
	}

	private Program getBottomProgram(HashMap<String, Program> programs) {
		Program bottomProgram = null;

		outerloop:
		while(true) {
			for (Map.Entry<String, Program> entry : programs.entrySet()) {
				HashMap<String, Program> childs = entry.getValue().getChilds();

				if (bottomProgram == null || (childs.size() > 0 && childs.containsKey(bottomProgram.getName()))) {
					bottomProgram = entry.getValue();
					continue outerloop;
				}
			}

			return bottomProgram;
		}
	}

	private HashMap<String, Program> getTree(String input) {
		HashMap<String, Program> programs = new HashMap<>();

		Scanner scanner = new Scanner(input);

		while(scanner.hasNextLine()) {
			String line = scanner.nextLine();
			String[] split = line.split("->");
			String name = split[0].split("[\\s()]")[0];
			int weight = Integer.parseInt(split[0].split("[\\s()]")[2]);

			Program value = new Program(name, weight);
			programs.put(name, value);

			if (split.length == 2) {
				value.addChildName(split[1]);
			}
		}

		for (Map.Entry<String, Program> entry : programs.entrySet()) {
			String childNames = entry.getValue().getChildNames();
			if (childNames != null) {
				String[] split = childNames.split(",");
				for (String childName : split) {
					childName = childName.trim();
					entry.getValue().addChild(programs.get(childName));
				}
			}
		}
		return programs;
	}

	private class Program {
		private final String name;
		private final int weight;
		private HashMap<String, Program> childs = new HashMap<>();
		private String childNames;

		public Program(String name, int weight) {
			this.name = name;
			this.weight = weight;
		}

		public String getName() {
			return name;
		}

		public int getWeight() {
			return weight;
		}

		public int getTotalWeight() {
			int totalWeight = this.weight;
			for (Program program : childs.values()) {
				totalWeight += program.getTotalWeight();
			}

			return totalWeight;
		}

		public void addChild(Program child) {
			this.childs.put(child.getName(), child);
		}

		public HashMap<String, Program> getChilds() {
			return childs;
		}

		public void addChildName(String childNames) {
			this.childNames = childNames;
		}

		public String getChildNames() {
			return childNames;
		}

		public int getCorrectWeight(int expected) {
			HashMap<Integer, Integer> weights = new HashMap<>();
			for (Program program : childs.values()) {
				int totalWeight = program.getTotalWeight();
				Integer currentCount = weights.get(totalWeight);
				if (currentCount == null) {
					currentCount = 0;
				}
				currentCount++;
				weights.put(totalWeight, currentCount);
			}

			int correctWeight = -1;
			int badWeight = -1;
			for (Map.Entry<Integer, Integer> entry : weights.entrySet()) {
				if (entry.getValue() == 1) {
					// this is bad
					badWeight = entry.getKey();
				}
				else {
					// this is good
					correctWeight = entry.getKey();
				}
			}

			if (badWeight == -1) {
				// this is the bad program. Return the diff.
				return this.weight - (getTotalWeight() - expected);
			}

			// find the bad one, if any
			for (Program program : childs.values()) {
				if (program.getTotalWeight() == badWeight) {
					return program.getCorrectWeight(correctWeight);
				}
			}

			return 0;
		}
	}
}

