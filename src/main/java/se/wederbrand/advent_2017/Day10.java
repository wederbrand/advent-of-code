package se.wederbrand.advent_2017;

public class Day10 {
	public int part1(int size, String input) {
		int skipLength = 0;
		int[] memory = new int[size];
		for (int i = 0; i < memory.length; i++) {
			memory[i] = i;
		}

		String[] reverseSizes = input.split(",");

		int index = 0;
		for (String reverseSizeString : reverseSizes) {
			int reverseSize = Integer.parseInt(reverseSizeString);
			int[] targetArray = new int[reverseSize];

			for (int i = 0; i < reverseSize; i++) {
				targetArray[i] = memory[(index + i) % size];
			}

			for (int i = 0; i < reverseSize; i++) {
				memory[(index + i) % size] = targetArray[reverseSize - i - 1];
			}

			index += reverseSize + skipLength;
			skipLength++;
		}

		return memory[0] * memory[1];
	}

	public String part2(int size, String input) {
		int skipLength = 0;
		int[] m = new int[size];
		for (int i = 0; i < m.length; i++) {
			m[i] = i;
		}

		char[] chars = new char[input.length() + 5];
		char[] a = input.toCharArray();
		char[] b = {17, 31, 73, 47, 23};
		System.arraycopy(a, 0, chars, 0, a.length);
		System.arraycopy(b, 0, chars, a.length, b.length);

		int index = 0;
		for (int j = 0; j < 64; j++) {
			for (char reverseSizeChar : chars) {
				int reverseSize = reverseSizeChar;
				int[] targetArray = new int[reverseSize];

				for (int i = 0; i < reverseSize; i++) {
					targetArray[i] = m[(index + i) % size];
				}

				for (int i = 0; i < reverseSize; i++) {
					m[(index + i) % size] = targetArray[reverseSize - i - 1];
				}

				index += reverseSize + skipLength;
				skipLength++;
			}
		}

		StringBuilder result = new StringBuilder();
		for (int i = 0; i < 16; i++) {
			int dense = m[16 * i] ^ m[16 * i + 1] ^ m[16 * i + 2] ^ m[16 * i + 3] ^ m[16 * i + 4] ^ m[16 * i + 5] ^ m[16 * i + 6] ^ m[16 * i + 7] ^ m[16 * i + 8] ^ m[16 * i + 9] ^ m[16 * i + 10] ^ m[16 * i + 11] ^ m[16 * i + 12] ^ m[16 * i + 13] ^ m[16 * i + 14] ^ m[16 * i + 15];
			String str = Integer.toHexString(dense);
			if (str.length() == 1) {
				result.append("0");
			}
			result.append(str);
		}
		return result.toString();
	}


}
