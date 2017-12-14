package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day14Test {
	private static final String INPUT = "jzgqcdpd";

	@Test
	public void testPart1() {
		assertEquals(8108, new Day14().part1( "flqrgnkx"));
	}

	@Test
	public void actualPart1() {
		System.out.println(new Day14().part1(INPUT));
	}

	@Test
	public void testPart2() {
		assertEquals(1242, new Day14().part2( "flqrgnkx"));
	}

	@Test
	public void actualPart2() {
		System.out.println(new Day14().part2(INPUT));
	}

}
