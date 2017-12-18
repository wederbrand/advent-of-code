package se.wederbrand.advent_2017;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class Day18 {
	HashMap<String, Long> register = new HashMap<>();
	String[] instructions;
	int index = 0;
	List<Long> queue = new ArrayList<>();
	Day18 other;
	private boolean blocked = false;
	private int totalSent = 0;

	public Day18() {
	}

	public Day18(String input, long id) {
		instructions = input.split(System.lineSeparator());
		register.put("p", id);
	}

	public List<Long> getQueue() {
		return queue;
	}

	public void setOther(Day18 other) {
		this.other = other;
	}

	public void addToQueue(long value) {
		queue.add(value);
		blocked = false;
	}

	public boolean isBlocked() {
		return blocked;
	}

	public int totalSent() {
		return totalSent;
	}

	public void run() {
		while (true) {
			String[] split = instructions[index].split(" ");
			switch (split[0]) {
				case "snd":
					other.addToQueue(getValue(split[1]));
					totalSent++;
					break;
				case "set":
					register.put(split[1], getValue(split[2]));
					break;
				case "add":
					register.put(split[1], getValue(split[1]) + getValue(split[2]));
					break;
				case "mul":
					register.put(split[1], getValue(split[1]) * getValue(split[2]));
					break;
				case "mod":
					register.put(split[1], getValue(split[1]) % getValue(split[2]));
					break;
				case "rcv":
					if (queue.isEmpty()) {
						blocked = true;
						return;
					}
					register.put(split[1], queue.remove(0));
					break;
				case "jgz":
					if (getValue(split[1]) > 0) {
						index += getValue(split[2]);
						continue;
					}
					break;
			}

			index++;
		}
	}

	public long part1(String input) {
		String[] instructions = input.split(System.lineSeparator());

		long sound = 0;
		int index = 0;

		while (true) {
			String[] split = instructions[index].split(" ");
			switch (split[0]) {
				case "snd":
					sound = getValue(split[1]);
					break;
				case "set":
					register.put(split[1], getValue(split[2]));
					break;
				case "add":
					register.put(split[1], getValue(split[1]) + getValue(split[2]));
					break;
				case "mul":
					register.put(split[1], getValue(split[1]) * getValue(split[2]));
					break;
				case "mod":
					register.put(split[1], getValue(split[1]) % getValue(split[2]));
					break;
				case "rcv":
					if (getValue(split[1]) > 0) {
						return sound;
					}
				case "jgz":
					if (getValue(split[1]) > 0) {
						index += Integer.parseInt(split[2]);
						continue;
					}
					break;
			}

			index++;
		}
	}

	public long part2(String input) {
		Day18 id0 = new Day18(input, 0);
		Day18 id1 = new Day18(input, 1);

		id0.setOther(id1);
		id1.setOther(id0);

		while (!id0.isBlocked() || !id1.isBlocked()) {
			id0.run();
			id1.run();
		}

		return id1.totalSent();
	}

	private Long getValue(String s) {
		try {
			return Long.parseLong(s);
		}
		catch (NumberFormatException e) {
			Long aLoing = register.get(s);
			if (aLoing == null) {
				return 0L;
			}
			return aLoing;
		}
	}

}

