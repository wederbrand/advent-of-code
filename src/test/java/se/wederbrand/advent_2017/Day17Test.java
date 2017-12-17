package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day17Test {
	String INPUT = "";

	@Test
	public void testPart1() {
		assertEquals(638, new Day17().part1(3));
	}

	@Test
	public void actualPart1() {
		System.out.println(new Day17().part1(354));
	}

	@Test
	public void actualPart2() {
		System.out.println(new Day17().part2(354));
	}

}
