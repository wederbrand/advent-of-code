package se.wederbrand.advent_2019;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day04Test {
	public static final String INPUT = "109165-576723";

	@Test
	public void testPart1() throws Exception {
		assertEquals(true, new Day04().part1internal("111111"));
		assertEquals(false, new Day04().part1internal("223450"));
		assertEquals(false, new Day04().part1internal("123789"));
	}

	@Test
	public void actualPart1() throws Exception {
		System.out.println(new Day04().part1(INPUT));
	}

	@Test
	public void testPart2() throws Exception {
		assertEquals(true, new Day04().part2internal("112233"));
		assertEquals(false, new Day04().part2internal("123444"));
		assertEquals(true, new Day04().part2internal("111144"));
	}

	@Test
	public void actualPart2() throws Exception {
		// 1756 is wrong
		System.out.println(new Day04().part2(INPUT));
	}
}

