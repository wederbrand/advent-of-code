package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day03Test {
	@Test
	public void testPart1() throws Exception {
		assertEquals(0, new Day03().part1(1));
		assertEquals(3, new Day03().part1(12));
		assertEquals(2, new Day03().part1(23));
		assertEquals(31   , new Day03().part1(1024));
	}

	@Test
	public void actualPart1() throws Exception {
		System.out.println(new Day03().part1(277678));
	}

	@Test
	public void testPart2() throws Exception {
	}

	@Test

	public void actualPart2() throws Exception {
		System.out.println(new Day03().part2(277678));
	}


}
