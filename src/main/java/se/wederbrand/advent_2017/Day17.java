package se.wederbrand.advent_2017;

import java.util.ArrayList;

public class Day17 {
	ArrayList<Integer> memory = new ArrayList<>();
	int position;

	public int part1(int steps) {
		memory.add(position, 0);
		position = 0;
		for (int i = 1; i <= 2017; i++) {
			position = (position + steps) % memory.size() + 1;
			memory.add(position, i);
		}
		return memory.get(position+1);
	}

	public int part2(int steps) {
		int length = 1;
		position = 0;
		int value=0;
		for (int i = 1; i <= 50000000; i++) {
			position = (position + steps) % length + 1;
			length++;
			if (position == 1) {
				value = i;
			}
		}
		return value;
	}

}
