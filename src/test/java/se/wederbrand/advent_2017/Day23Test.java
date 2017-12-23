package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day23Test {
	String INPUT =
		"set b 79\n" +
		"set c b\n" +
		"jnz a 2\n" +
		"jnz 1 5\n" +
		"mul b 100\n" +
		"sub b -100000\n" +
		"set c b\n" +
		"sub c -17000\n" +
		"set f 1\n" +
		"set d 2\n" +
		"set e 2\n" +
		"set g d\n" +
		"mul g e\n" +
		"sub g b\n" +
		"jnz g 2\n" +
		"set f 0\n" +
		"sub e -1\n" +
		"set g e\n" +
		"sub g b\n" +
		"jnz g -8\n" +
		"sub d -1\n" +
		"set g d\n" +
		"sub g b\n" +
		"jnz g -13\n" +
		"jnz f 2\n" +
		"sub h -1\n" +
		"set g b\n" +
		"sub g c\n" +
		"jnz g 2\n" +
		"jnz 1 3\n" +
		"sub b -17\n" +
		"jnz 1 -23";

	@Test
	public void actualPart1() {
		System.out.println(new Day23().part1(INPUT));
	}

	@Test
	public void actualPart2() {
		System.out.println(new Day23().part2());
	}

}
