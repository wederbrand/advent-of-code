package se.wederbrand.advent_2019;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day12Test {

	public static final String INPUT = "<x=-16, y=15, z=-9>\n" +
		"<x=-14, y=5, z=4>\n" +
		"<x=2, y=0, z=6>\n" +
		"<x=-3, y=18, z=9>";

	@Test
	public void testPart1() throws Exception {
		assertEquals(179, new Day12("<x=-1, y=0, z=2>\n" +
			"<x=2, y=-10, z=-7>\n" +
			"<x=4, y=-8, z=8>\n" +
			"<x=3, y=5, z=-1>").part1(10));
	}

	@Test
	public void actualPart1() throws Exception {
		System.out.println(new Day12(INPUT).part1(1000));
	}

	@Test
	public void testPart2() throws Exception {
	}

	@Test
	public void actualPart2() throws Exception {
	}


}
