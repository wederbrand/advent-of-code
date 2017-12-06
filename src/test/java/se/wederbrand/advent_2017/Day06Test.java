package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day06Test {

	public static final String INPUT = "10\t3\t15\t10\t5\t15\t5\t15\t9\t2\t5\t8\t5\t2\t3\t6";

	@Test
	public void testPart1() throws Exception {
		assertEquals(5, new Day06().part1("0\t2\t7\t0"));
	}

	@Test
	public void testPart2() throws Exception {
		assertEquals(4, new Day06().part2("0\t2\t7\t0"));
	}

	@Test
	public void actualPart1() throws Exception {
		System.out.println(new Day06().part1(INPUT));
	}

	@Test
	public void actualPart2() throws Exception {
		System.out.println(new Day06().part2(INPUT));
	}


}
