package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day18Test {
	String INPUT = "set i 31\n" +
		"set a 1\n" +
		"mul p 17\n" +
		"jgz p p\n" +
		"mul a 2\n" +
		"add i -1\n" +
		"jgz i -2\n" +
		"add a -1\n" +
		"set i 127\n" +
		"set p 622\n" +
		"mul p 8505\n" +
		"mod p a\n" +
		"mul p 129749\n" +
		"add p 12345\n" +
		"mod p a\n" +
		"set b p\n" +
		"mod b 10000\n" +
		"snd b\n" +
		"add i -1\n" +
		"jgz i -9\n" +
		"jgz a 3\n" +
		"rcv b\n" +
		"jgz b -1\n" +
		"set f 0\n" +
		"set i 126\n" +
		"rcv a\n" +
		"rcv b\n" +
		"set p a\n" +
		"mul p -1\n" +
		"add p b\n" +
		"jgz p 4\n" +
		"snd a\n" +
		"set a b\n" +
		"jgz 1 3\n" +
		"snd b\n" +
		"set f 1\n" +
		"add i -1\n" +
		"jgz i -11\n" +
		"snd a\n" +
		"jgz f -16\n" +
		"jgz a -19";

	@Test
	public void testPart1() {
		assertEquals(4, new Day18().part1("set a 1\n" +
			"add a 2\n" +
			"mul a a\n" +
			"mod a 5\n" +
			"snd a\n" +
			"set a 0\n" +
			"rcv a\n" +
			"jgz a -1\n" +
			"set a 1\n" +
			"jgz a -2"));
	}

	@Test
	public void testPart2() {
		assertEquals(3, new Day18().part2("snd 1\n" +
			"snd 2\n" +
			"snd p\n" +
			"rcv a\n" +
			"rcv b\n" +
			"rcv c\n" +
			"rcv d"));
	}

	@Test
	public void actualPart1() {
		System.out.println(new Day18().part1(INPUT));
	}

	@Test
	public void actualPart2() {
		System.out.println(new Day18().part2(INPUT));
	}

}
