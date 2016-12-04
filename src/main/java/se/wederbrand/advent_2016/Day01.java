package se.wederbrand.advent_2016;

import com.google.common.collect.Iterables;

import java.util.ArrayList;
import java.util.Scanner;

public class Day01 {

	public int part1(String input) {
		int ns = 0;
		int ew = 0;
		int facing = 0;

		Scanner scanner = new Scanner(input).useDelimiter("\\s*,\\s*");
		while (scanner.hasNext()) {
			String instruction = scanner.next();
			char direction = instruction.charAt(0);
			int steps = Integer.parseInt(instruction.substring(1));

			facing = getFacing(facing, direction);

			switch (facing) {
				case 0:
					ns += steps;
					break;
				case 90:
					ew += steps;
					break;
				case 180:
					ns -= steps;
					break;
				case 270:
					ew -= steps;
					break;
			}
		}

		return Math.abs(ns) + Math.abs(ew);
	}

	public int part2(String input) {
		int ns = 0;
		int ew = 0;
		int facing = 0;

		ArrayList<String> trace = new ArrayList<String>();
		trace.add(ns + ":" + ew);

		Scanner scanner = new Scanner(input).useDelimiter("\\s*,\\s*");
		while (scanner.hasNext()) {
			String instruction = scanner.next();
			char direction = instruction.charAt(0);
			int steps = Integer.parseInt(instruction.substring(1));

			facing = getFacing(facing, direction);

			int distance = walkUntilRevisited(trace, facing, steps);
			if (distance != -1) {
				return distance;
			}
		}

		return -1;
	}

	private int walkUntilRevisited(ArrayList<String> trace, int facing, int steps) {
		String last = Iterables.getLast(trace);
		String[] split = last.split(":");
		int ns = Integer.valueOf(split[0]);
		int ew = Integer.valueOf(split[1]);

		for (int i = 0; i < steps; i++) {
			switch (facing) {
				case 0:
					ns++;
					break;
				case 90:
					ew++;
					break;
				case 180:
					ns--;
					break;
				case 270:
					ew--;
					break;
			}

			if (trace.contains(ns + ":" + ew)) {
				return Math.abs(ns) + Math.abs(ew);
			}

			trace.add(ns + ":" + ew);
		}

		return -1;
	}

	private int getFacing(int facing, char direction) {
		switch (direction) {
			case 'L':
				facing -= 90;
				break;
			case 'R':
				facing += 90;
				break;
		}

		facing = (facing + 360) % 360;
		return facing;
	}
}
