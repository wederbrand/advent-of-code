package advent_2017;


import java.util.Arrays;
import java.util.HashMap;
import java.util.Scanner;

public class Day04 {
	public int part1(String input) {
		int result = 0;

		Scanner scanner = new Scanner(input);
		outerLooop:
		while (scanner.hasNextLine()) {
			String line = scanner.nextLine();

			HashMap words = new HashMap();

			Scanner lineScanner = new Scanner(line);
			while (lineScanner.hasNext()) {
				String word = lineScanner.next();
				if (words.containsKey(word)) {
					continue outerLooop;
				}
				words.put(word, 1);
			}
			result++;
		}

		return result;
	}

	public int part2(String input) {
		int result = 0;

		Scanner scanner = new Scanner(input);
		outerLooop:
		while (scanner.hasNextLine()) {
			String line = scanner.nextLine();

			HashMap words = new HashMap();

			Scanner lineScanner = new Scanner(line);
			while (lineScanner.hasNext()) {
				String word = lineScanner.next();
				word = sortWord(word);
				if (words.containsKey(word)) {
					continue outerLooop;
				}
				words.put(word, 1);
			}
			result++;
		}

		return result;
	}

	private String sortWord(String word) {
		char[] chars = word.toCharArray();
		Arrays.sort(chars);
		return new String(chars);
	}

}

