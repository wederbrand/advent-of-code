package se.wederbrand.advent_2017;

import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class Day07 {
	public String part1(String input) {
		HashMap<String, String> programs = new HashMap<>();

		Scanner scanner = new Scanner(input);

		while(scanner.hasNextLine()) {
			String line = scanner.nextLine();
			String[] split = line.split("->");
			String name = split[0].split("\\s")[0];

			if (split.length == 2) {
				programs.put(name, split[1]);
			}
			else {
				programs.put(name, "");
			}
		}

		// build the tree

		String bottomProgram = "";

		outerloop:
		while(true) {
			for (Map.Entry<String, String> entry : programs.entrySet()) {

				if (bottomProgram.equals("") || entry.getValue().contains(bottomProgram)) {
					bottomProgram = entry.getKey();
					continue outerloop;
				}
			}

			return bottomProgram;
		}
	}

	public String part2(String input) {
		HashMap<String, String> programs = new HashMap<>();

		Scanner scanner = new Scanner(input);

		while(scanner.hasNextLine()) {
			String line = scanner.nextLine();
			String[] split = line.split("->");
			String name = split[0].split("\\s")[0];

			if (split.length == 2) {
				programs.put(name, split[1]);
			}
			else {
				programs.put(name, "");
			}
		}

		// build the tree

		String bottomProgram = "";

		outerloop:
		while(true) {
			for (Map.Entry<String, String> entry : programs.entrySet()) {

				if (bottomProgram.equals("") || entry.getValue().contains(bottomProgram)) {
					bottomProgram = entry.getKey();
					continue outerloop;
				}
			}

			return bottomProgram;
		}
	}

}

