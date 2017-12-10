package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day10Test {
	private static final String INPUT = "14,58,0,116,179,16,1,104,2,254,167,86,255,55,122,244";

	@Test
	public void testPart1() {
		assertEquals(12, new Day10().part1(5, "3,4,1,5"));
	}

	@Test
	public void actualPart1() {
		System.out.println(new Day10().part1(256, INPUT));
	}

	@Test
	public void testPart2() {
		assertEquals("a2582a3a0e66e6e86e3812dcb672a272", new Day10().part2(256, ""));
		assertEquals("33efeb34ea91902bb2f59c9920caa6cd", new Day10().part2(256, "AoC 2017"));
		assertEquals("3efbe78a8d82f29979031a4aa0b16a9d", new Day10().part2(256, "1,2,3"));
		assertEquals("63960835bcdc130f0b66d7ff4f6a5a8e", new Day10().part2(256, "1,2,4"));
	}

	@Test
	public void actualPart2() {
		System.out.println(new Day10().part2(256, INPUT));
	}

}
