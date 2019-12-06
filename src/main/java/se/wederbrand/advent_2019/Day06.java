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

	public int part2(String input) {
		String[] lines = input.split(System.lineSeparator());
		for (String line : lines) {
			String[] split = line.split("\\)");
			Stuff center = map.computeIfAbsent(split[0], Stuff::new);
			Stuff orbiter = map.computeIfAbsent(split[1], Stuff::new);
			center.has(orbiter);
			orbiter.orbits(center);
		}

		Stuff you = map.get("YOU");
		Stuff san = map.get("SAN");

		Stuff com = map.get("COM");
		com.recursiveMagic(0);
		ArrayList<Stuff> youPath = you.getPathTo(com);
		ArrayList<Stuff> sanPath = san.getPathTo(com);

		// find forking node
		Stuff forkNode = null;
		for (Stuff stuff : youPath) {
			if (sanPath.contains(stuff)) {
				forkNode = stuff;
			}
			else {
				// no longer matching nodes
				break;
			}
		}

		return (you.distance - forkNode.distance - 1) + (san.distance - forkNode.distance - 1) ;
	}


	private class Stuff {
		private String name;
		private int distance;
		private List<Stuff> orbiters = new ArrayList<>();
		private Stuff orbits;

		public Stuff(String name) {
			this.name = name;
		}

		public void has(Stuff orbiter) {
			orbiters.add(orbiter);
		}

		public void orbits(Stuff stuff) {
			this.orbits = stuff;
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

		public ArrayList<Stuff> getPathTo(Stuff target) {
			ArrayList<Stuff> list;
			if (target == orbits) {
				list = new ArrayList<>();
			}
			else {
				list = orbits.getPathTo(target);
			}
			list.add(this);
			return list;
		}
	}
}
