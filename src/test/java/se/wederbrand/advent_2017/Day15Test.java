package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day15Test {
	@Test
	public void testPart1() {
		assertEquals(588, new Day15().part1( 65, 8921));
	}

	@Test
	public void actualPart1() {
		System.out.println(new Day15().part1(883, 879));
	}

	@Test
	public void testPart2() {
		assertEquals(309, new Day15().part2( 65, 8921));
	}

	@Test
	public void actualPart2() {
		System.out.println(new Day15().part2(883, 879));
	}

}
