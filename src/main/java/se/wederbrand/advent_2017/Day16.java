package se.wederbrand.advent_2017;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;

public class Day16 {
	private char[] programs;

	public String part1(int size, String input) {
		createPrograms(size);
		List<Move> moves = getMoves(input);
		danceOnce(moves);
		return getResults();
	}

	private void danceOnce(List<Move> moves) {
		for (Move move : moves) {
			move.dance();
		}
	}

	public String part2(int size, String input) {
		createPrograms(size);
		List<Move> moves = getMoves(input);
		List<String> shortCuts = new ArrayList<>();
		for (int i = 0; i < 1000000000; i++) {
			String before = getResults();
			shortCuts.add(before);
			danceOnce(moves);
			String after = getResults();
			if (shortCuts.contains(after)) {
				// now I've seen it all
				break;
			}
		}

		return shortCuts.get(1000000000 % shortCuts.size());
	}

	private String getResults() {
		StringBuilder result = new StringBuilder();
		for (char c : programs) {
			result.append(c);
		}
		return result.toString();
	}

	private List<Move> getMoves(String input) {
		List<Move> moves = new ArrayList<>();
		Scanner scanner = new Scanner(input).useDelimiter(",");
		while (scanner.hasNext()) {
			String move = scanner.next();

			switch (move.charAt(0)) {
				case 's':
					moves.add(createSpin(move.substring(1)));
					break;
				case 'x':
					moves.add(createExchange(move.substring(1)));
					break;
				case 'p':
					moves.add(createPartner(move.substring(1)));
					break;
			}
		}
		return moves;
	}

	private void createPrograms(int size) {
		programs = new char[size];
		char program = 'a';
		for (int i = 0; i < size; i++) {
			programs[i] = program++;
		}
	}

	private Move createSpin(String spin) {
		int spinSize = Integer.parseInt(spin);
		return new SpinMove(spinSize);
	}

	private Move createExchange(String exchange) {
		String[] split = exchange.split("/");
		int src = Integer.parseInt(split[0]);
		int dest = Integer.parseInt(split[1]);
		return new ExchangeMove(src, dest);
	}

	private Move createPartner(String partner) {
		char src = partner.charAt(0);
		char dest = partner.charAt(2);

		return new PartnerMove(src, dest);
	}

	private interface Move {
		void dance();
	}

	private class SpinMove implements Move {
		private final int spinSize;

		SpinMove(int spinSize) {
			this.spinSize = spinSize;
		}

		@Override
		public void dance() {
			char[] newPrograms = new char[programs.length];
			System.arraycopy(programs, programs.length - spinSize, newPrograms, 0, spinSize);
			System.arraycopy(programs, 0, newPrograms, spinSize, programs.length - spinSize);

			programs = newPrograms;
		}
	}

	private class ExchangeMove implements Move {
		private final int src;
		private final int dest;

		ExchangeMove(int src, int dest) {
			this.src = src;
			this.dest = dest;
		}

		@Override
		public void dance() {
			char slask = programs[src];
			programs[src] = programs[dest];
			programs[dest] = slask;
		}
	}

	private class PartnerMove implements Move {
		private final char src;
		private final char dest;

		PartnerMove(char src, char dest) {
			this.src = src;
			this.dest = dest;
		}

		@Override
		public void dance() {
			for (int i = 0; i < programs.length; i++) {
				char program = programs[i];
				if (program == src) {
					programs[i] = dest;
				}
				if (program == dest) {
					programs[i] = src;
				}
			}
		}
	}
}
