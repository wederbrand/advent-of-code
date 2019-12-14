package se.wederbrand.advent_2019;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class Day14Test {

	public static final String INPUT = "1 XVCBM, 12 SWPQ => 7 VMWSR\n" +
		"10 SBLTQ, 14 TLDR => 6 HJFPQ\n" +
		"1 VWHXC, 2 GZDQ, 3 PCLMJ => 4 VJPLN\n" +
		"9 MGVG => 7 WDPF\n" +
		"1 FBXD, 5 FZNZR => 6 GZDQ\n" +
		"5 TJPZ, 1 QNMZ => 5 SWPQ\n" +
		"12 XWQW, 1 HJFPQ => 8 JPKNC\n" +
		"15 CPNC, 2 TXKRN, 2 MTVQD => 9 LBRSX\n" +
		"5 VJPLN, 1 VSTRK, 2 GFQLV => 5 NLZKH\n" +
		"1 TLDR => 4 TNRZW\n" +
		"2 VCFM => 7 FZNZR\n" +
		"1 PSTRV, 5 RTDV => 8 VCFM\n" +
		"2 PSTRV => 9 SFWJG\n" +
		"4 XWQW => 2 BHPS\n" +
		"1 ZWFNW, 19 JKRWT, 2 JKDL, 8 PCLMJ, 7 FHNL, 22 MSZCF, 1 VSTRK, 7 DMJPR => 1 ZDGF\n" +
		"22 XVCBM, 8 TBLM => 1 MTVQD\n" +
		"101 ORE => 1 WBNWZ\n" +
		"6 VNVXJ, 1 FBXD, 13 PCLMJ => 9 MGVG\n" +
		"13 SHWB, 1 WDPF, 4 QDTW => 6 FHNL\n" +
		"9 VSTRK => 2 VZCML\n" +
		"20 LZCDB => 7 KNPM\n" +
		"2 LBRSX, 9 GRCD => 3 SHWB\n" +
		"5 BHPS => 6 SQJLW\n" +
		"1 RTDV => 6 GRCD\n" +
		"6 SBLTQ, 6 XWQW => 5 CPNC\n" +
		"153 ORE => 3 RTDV\n" +
		"6 LZCDB, 1 SBLTQ => 3 PCLMJ\n" +
		"1 RTDV, 2 TJPZ => 5 LZCDB\n" +
		"24 QNMZ => 4 TXKRN\n" +
		"19 PCLMJ, 7 VNVXJ => 6 RKRVJ\n" +
		"12 RKRVJ, 11 QNMZ => 3 JKRWT\n" +
		"4 SFWJG => 9 FBXD\n" +
		"16 WDPF, 4 TXKRN => 6 DMJPR\n" +
		"3 QNMZ => 1 VSTRK\n" +
		"9 VSTRK => 4 ZWFNW\n" +
		"7 QBWN, 1 TLDR => 4 QDTW\n" +
		"7 VJPLN, 1 NLZKH, 15 JPKNC, 3 SHWB, 1 MSZCF, 3 VMWSR => 6 QDHGS\n" +
		"14 QXQZ => 7 XWQW\n" +
		"152 ORE => 9 TJPZ\n" +
		"1 PJVJ, 10 QBWN, 19 NLZKH => 6 MSZCF\n" +
		"21 TLDR, 13 VNVXJ, 5 BHPS => 4 QBWN\n" +
		"1 GZDQ, 6 GRCD => 9 TLDR\n" +
		"4 BHPS => 8 MZBL\n" +
		"1 FZNZR => 2 VNVXJ\n" +
		"1 VNVXJ => 5 GFQLV\n" +
		"13 LZCDB => 2 QXQZ\n" +
		"3 MNFJX => 5 VWHXC\n" +
		"1 GZDQ, 2 VMWSR => 6 WZMHW\n" +
		"9 HJFPQ, 3 RKRVJ => 4 QNMZ\n" +
		"8 TJPZ => 9 SBLTQ\n" +
		"30 WBNWZ => 5 TBLM\n" +
		"1 PCLMJ => 3 GNMTQ\n" +
		"30 SQJLW, 3 QNMZ, 9 WDPF => 5 PJVJ\n" +
		"10 GRCD, 15 SBLTQ, 22 GFQLV => 4 XVCBM\n" +
		"30 PJVJ, 10 JPKNC, 3 DXFDR, 10 VZCML, 59 MZBL, 40 VWHXC, 1 ZDGF, 13 QDHGS => 1 FUEL\n" +
		"4 GNMTQ, 6 VMWSR, 19 RKRVJ, 5 FKZF, 4 VCFM, 2 WZMHW, 7 KNPM, 5 TNRZW => 7 DXFDR\n" +
		"152 ORE => 9 PSTRV\n" +
		"2 BHPS, 5 TXKRN, 2 PJVJ => 4 FKZF\n" +
		"2 XWQW, 2 VCFM, 13 BHPS => 8 MNFJX\n" +
		"3 XWQW => 2 JKDL";

	@Test
	public void testPart1() throws Exception {
		assertEquals(31, new Day14("10 ORE => 10 A\n" +
			"1 ORE => 1 B\n" +
			"7 A, 1 B => 1 C\n" +
			"7 A, 1 C => 1 D\n" +
			"7 A, 1 D => 1 E\n" +
			"7 A, 1 E => 1 FUEL").part1());

			assertEquals(165, new Day14("9 ORE => 2 A\n" +
				"8 ORE => 3 B\n" +
				"7 ORE => 5 C\n" +
				"3 A, 4 B => 1 AB\n" +
				"5 B, 7 C => 1 BC\n" +
				"4 C, 1 A => 1 CA\n" +
				"2 AB, 3 BC, 4 CA => 1 FUEL").part1());
	}

	@Test
	public void actualPart1() throws Exception {
		System.out.println(new Day14(INPUT).part1());
	}

	@Test
	public void testPart2() throws Exception {
		assertEquals(82892753, new Day14("157 ORE => 5 NZVS\n" +
			"165 ORE => 6 DCFZ\n" +
			"44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL\n" +
			"12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ\n" +
			"179 ORE => 7 PSHF\n" +
			"177 ORE => 5 HKGWZ\n" +
			"7 DCFZ, 7 PSHF => 2 XJWVT\n" +
			"165 ORE => 2 GPVTF\n" +
			"3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT").part2());
	}

	@Test
	public void actualPart2() throws Exception {
		System.out.println(new Day14(INPUT).part2());
	}


}
