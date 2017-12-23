package se.wederbrand.advent_2017;

import java.util.HashMap;

public class Day23 {
	HashMap<String, Long> register = new HashMap<>();

	public long part1(String input) {
		String[] instructions = input.split(System.lineSeparator());

		int index = 0;

		long count = 0;
		while (true) {
			if (index < 0 || index >= instructions.length) {
				return count;
			}
			String[] split = instructions[index].split(" ");
			switch (split[0]) {
				case "set":
					register.put(split[1], getValue(split[2]));
					break;
				case "sub":
					register.put(split[1], getValue(split[1]) - getValue(split[2]));
					break;
				case "mul":
					register.put(split[1], getValue(split[1]) * getValue(split[2]));
					count++;
					break;
				case "jnz":
					if (getValue(split[1]) != 0) {
						index += Integer.parseInt(split[2]);
						continue;
					}
					break;
			}

			index++;
		}
	}

	public long part2() {
		long b;
		long c;
		long d;
		long e;
		long f;
		long g;
		long h = 0;

		//    00  set b 79
		//		01  set c b
		//		02  jnz a 2
		//		03  jnz 1 5
		//		04  mul b 100
		//		05  sub b -100000
		b = 79*100 + 100000;
		//		06  set c b
		//		07  sub c -17000
		c = b + 17000;

		do {
			//		08  set f 1
			f = 1;
			//		09  set d 2
			d = 2;

			primeloop:
			do {
				//		10  set e 2
				e = 2;

				do {
					//		11  set g d
					//		12  mul g e
					//		13  sub g b
					//		14  jnz g 2
					if (d*e == b) {
						//		15  set f 0
						// primes that factors b?
						f = 0;
						// getting here means we can skip the large loop, e and d is not essential any more
						break primeloop;
					}
					if (d*e > b) {
						// no matter what we add to e, it will never be a factor in b;
						break;
					}
					//		16  sub e -1
					e++;
					//		17  set g e
					//		18  sub g b
					//		19  jnz g -8
				} while (e != b);
				//		20  sub d -1
				d++;
				//		21  set g d
				//		22  sub g b
				//		23  jnz g -13
			} while (d != b);
			//		24  jnz f 2
			if (f == 0) {
				//		25  sub h -1
				h++;
			}
			//		26  set g b
			//		27  sub g c
			//		28  jnz g 2
			if (b == c) {
				//		29  jnz 1 3
				return h;
			}

			//		30  sub b -17
			b -= -17;

			//		31  jnz 1 -23
		} while (true);
	}

	private long getValue(String s) {
		try {
			return Long.parseLong(s);
		}
		catch (NumberFormatException e) {
			return register.getOrDefault(s, 0L);
		}
	}

}

