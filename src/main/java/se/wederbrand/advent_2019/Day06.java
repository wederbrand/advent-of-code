package se.wederbrand.advent_2019;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class Day06 {
	HashMap<String, Stuff> map = new HashMap<>();

	public int part1(String input) {
		String[] lines = input.split(System.lineSeparator());
		for (String line : lines) {
			String[] split = line.split("\\)");
			Stuff center = map.computeIfAbsent(split[0], Stuff::new);
			Stuff orbiter = map.computeIfAbsent(split[1], Stuff::new);
			center.has(orbiter);
		}

		Stuff com = map.get("COM");
		return com.recursiveMagic(0);
	}


	private class Stuff {
		private String name;
		private int distance;
		private List<Stuff> orbiters = new ArrayList<>();

		public Stuff(String name) {
			this.name = name;
		}

		public void has(Stuff orbiter) {
			orbiters.add(orbiter);
		}

		@Override
		public String toString() {
			return "Stuff{" +
				"name='" + name + '\'' +
				'}';
		}

		public int recursiveMagic(int distance) {
			this.distance = distance;
			if (orbiters.size() == 0) {
				return distance;
			}

			int sum = 0;
			for (Stuff orbiter : orbiters) {
				sum += orbiter.recursiveMagic(distance+1);
			}
			return distance + sum;
		}
	}
}
