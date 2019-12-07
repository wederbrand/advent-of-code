package se.wederbrand.advent_2019;

import java.util.*;

public class Day07 {
	public List<int[]> getPermutations() {
		List<int[]> permutations = new ArrayList<>();

		for (int a = 0; a < 5; a++) {
			for (int b = 0; b < 5; b++) {
				if (b==a) continue;
				for (int c = 0; c < 5; c++) {
					if (c==a || c==b) continue;
					for (int d = 0; d < 5; d++) {
						if (d==a || d==b || d==c) continue;
						for (int e = 0; e < 5; e++) {
							if (e==a || e==b || e==c || e==d) continue;
							int[] permutation = {a, b, c, d, e};
							permutations.add(permutation);
						}
					}
				}
			}
		}

		return permutations;
	}

	public int machineOfMachines(int a, int b, int c, int d, int e, String input) {
		Stack<Integer> ints = new Stack<>();
		ints.push(0);
		ints.push(a);
		ints.push(machine(input, ints));
		ints.push(b);
		ints.push(machine(input, ints));
		ints.push(c);
		ints.push(machine(input, ints));
		ints.push(d);
		ints.push(machine(input, ints));
		ints.push(e);
		return machine(input, ints);
	}

	public int machineOfLoopingMachines(int a, int b, int c, int d, int e, String input) {
		Stack<Integer> ints = new Stack<>();
		ints.push(0);
		ints.push(a);
		ints.push(machine(input, ints));
		ints.push(b);
		ints.push(machine(input, ints));
		ints.push(c);
		ints.push(machine(input, ints));
		ints.push(d);
		ints.push(machine(input, ints));
		ints.push(e);
		ints.push(machine(input, ints));

		//  now loop
		try {
			int i = 0;
			while (true) {
				ints.push(machine(input, ints));
			}
		}
		catch (Exception ex) {
			System.out.println("done");
		}

		return ints.pop();
	}

	public int bestOfMachines(String input) {
		List<int[]> permutations = getPermutations();
		int max = 0;
		for (int[] permutation : permutations) {
			int i = machineOfMachines(permutation[0], permutation[1], permutation[2], permutation[3], permutation[4], input);
			if (i>max) {
				max = i;
			}
		}

		return max;
	}

	public int machine(String input, Stack<Integer> inputTo3) {
		int[] ints = Arrays.stream(input.split(",")).mapToInt(Integer::parseInt).toArray();
		int i = 0;
		outer:
		while (true) {
			int opCode = ints[i] % 100;
			int c = (ints[i]) / 100 % 10;
			int b = (ints[i]) / 1000 % 10;
			int a = (ints[i]) / 10000 % 10;

			int param1 = 0;
			int param2 = 0;
			try {
				param1 = c == 0 ? ints[ints[i + 1]] : ints[i + 1];
				param2 = b == 0 ? ints[ints[i + 2]] : ints[i + 2];
			}
			catch (ArrayIndexOutOfBoundsException e) {
				// ignore, it happens on some instructions
			}

			switch (opCode) {
				case 1: // +
					ints[ints[i + 3]] = param1 + param2;
					i += 4;
				break;
				case 2: // *
					ints[ints[i + 3]] = param1 * param2;
					i += 4;
				break;
				case 3: // input
					ints[ints[i + 1]] = inputTo3.pop();
					i += 2;
					break;
				case 4: // output
					if (c == 0) {
						return ints[ints[i + 1]];
					} else {
						return ints[i + 1];
					}
				case 5: // jump if true
					if (param1 != 0) {
						i = param2;
					}
					else {
						i+=3;
					}
					break;
				case 6: // jump if false
					if (param1 == 0) {
						i = param2;
					}
					else {
						i+=3;
					}
					break;
				case 7: // less than
					if (param1 < param2) {
						ints[ints[i + 3]] = 1;
					}
					else {
						ints[ints[i + 3]] = 0;
					}
					i+=4;
					break;
				case 8: // equals
					if (param1 == param2) {
						ints[ints[i + 3]] = 1;
					}
					else {
						ints[ints[i + 3]] = 0;
					}
					i+=4;
					break;
				case 99:
					throw new RuntimeException("done");
			}
		}
	}
}
