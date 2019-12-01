package se.wederbrand.advent_2019;

import org.junit.Test;

import static org.junit.Assert.*;
import static org.junit.Assert.assertEquals;

public class Day01Test {

	public static final String INPUT = "89822\n" +
		"149236\n" +
		"106135\n" +
		"147663\n" +
		"91417\n" +
		"59765\n" +
		"66470\n" +
		"121156\n" +
		"148632\n" +
		"116660\n" +
		"90316\n" +
		"111666\n" +
		"142111\n" +
		"72595\n" +
		"139673\n" +
		"145157\n" +
		"77572\n" +
		"83741\n" +
		"79815\n" +
		"74693\n" +
		"139077\n" +
		"106066\n" +
		"125817\n" +
		"127827\n" +
		"103884\n" +
		"147289\n" +
		"81588\n" +
		"142651\n" +
		"69916\n" +
		"147214\n" +
		"71501\n" +
		"130067\n" +
		"60182\n" +
		"139195\n" +
		"115502\n" +
		"127751\n" +
		"95013\n" +
		"73411\n" +
		"125294\n" +
		"79809\n" +
		"118110\n" +
		"122547\n" +
		"145141\n" +
		"72231\n" +
		"138853\n" +
		"108119\n" +
		"139960\n" +
		"128665\n" +
		"107228\n" +
		"73416\n" +
		"54608\n" +
		"63811\n" +
		"72363\n" +
		"130546\n" +
		"61055\n" +
		"56786\n" +
		"127718\n" +
		"144953\n" +
		"149284\n" +
		"137318\n" +
		"109566\n" +
		"112866\n" +
		"148063\n" +
		"130570\n" +
		"67536\n" +
		"84011\n" +
		"123795\n" +
		"128098\n" +
		"51687\n" +
		"83758\n" +
		"59867\n" +
		"103122\n" +
		"77339\n" +
		"72126\n" +
		"71446\n" +
		"67162\n" +
		"112342\n" +
		"120248\n" +
		"137629\n" +
		"135736\n" +
		"139781\n" +
		"92512\n" +
		"105922\n" +
		"85458\n" +
		"148571\n" +
		"51173\n" +
		"135047\n" +
		"110175\n" +
		"93722\n" +
		"82611\n" +
		"128288\n" +
		"125225\n" +
		"104177\n" +
		"115081\n" +
		"78470\n" +
		"96167\n" +
		"138445\n" +
		"117778\n" +
		"100133\n" +
		"140047";

	@Test
	public void testPart1() throws Exception {
		assertEquals(2, new Day01().part1("12"));
		assertEquals(2, new Day01().part1("14"));
		assertEquals(654, new Day01().part1("1969"));
		assertEquals(33583, new Day01().part1("100756"));
	}

	@Test
	public void actualPart1() throws Exception {
		System.out.println(new Day01().part1(INPUT));
	}

	@Test
	public void testPart2() throws Exception {
		assertEquals(2, new Day01().part2("14"));
		assertEquals(966, new Day01().part2("1969"));
		assertEquals(50346, new Day01().part2("100756"));
	}

	@Test
	public void actualPart2() throws Exception {
		System.out.println(new Day01().part2(INPUT));
	}


}
