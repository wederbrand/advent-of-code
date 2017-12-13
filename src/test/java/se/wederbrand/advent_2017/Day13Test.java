package se.wederbrand.advent_2017;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day13Test {
	private static final String INPUT = "0: 5\n" +
		"1: 2\n" +
		"2: 3\n" +
		"4: 4\n" +
		"6: 6\n" +
		"8: 4\n" +
		"10: 8\n" +
		"12: 6\n" +
		"14: 6\n" +
		"16: 14\n" +
		"18: 6\n" +
		"20: 8\n" +
		"22: 8\n" +
		"24: 10\n" +
		"26: 8\n" +
		"28: 8\n" +
		"30: 10\n" +
		"32: 8\n" +
		"34: 12\n" +
		"36: 9\n" +
		"38: 20\n" +
		"40: 12\n" +
		"42: 12\n" +
		"44: 12\n" +
		"46: 12\n" +
		"48: 12\n" +
		"50: 12\n" +
		"52: 12\n" +
		"54: 12\n" +
		"56: 14\n" +
		"58: 14\n" +
		"60: 14\n" +
		"62: 20\n" +
		"64: 14\n" +
		"66: 14\n" +
		"70: 14\n" +
		"72: 14\n" +
		"74: 14\n" +
		"76: 14\n" +
		"78: 14\n" +
		"80: 12\n" +
		"90: 30\n" +
		"92: 17\n" +
		"94: 18";

	@Test
	public void testPart1() {
		assertEquals(24, new Day13().part1( "0: 3\n" +
			"1: 2\n" +
			"4: 4\n" +
			"6: 4"));
	}

	@Test
	public void actualPart1() {
		System.out.println(new Day13().part1(INPUT));
	}

	@Test
	public void testPart2() {
		assertEquals(10, new Day13().part2( "0: 3\n" +
			"1: 2\n" +
			"4: 4\n" +
			"6: 4"));
	}

	@Test
	public void actualPart2() {
		System.out.println(new Day13().part2(INPUT));
	}

}
