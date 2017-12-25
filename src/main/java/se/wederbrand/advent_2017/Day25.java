package se.wederbrand.advent_2017;

import java.util.Arrays;
import java.util.HashMap;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day25 {
	private HashMap<Long, Integer> tape = new HashMap<>();
	private HashMap<String, State> states = new HashMap<>();
	private long tapeIndex = 0;

	public long part1(String input) {
		String[] split = input.split(System.lineSeparator());
		Pattern pattern = Pattern.compile("(\\d+)");
		Matcher matcher = pattern.matcher(split[1]);
		matcher.find();
		long iterations = Long.parseLong(matcher.group(1));

		int i = 3;
		while (i < split.length) {
			State state = new State(Arrays.copyOfRange(split, i, i + 9));
			states.put(state.getState(), state);
			i += 10;
		}

		State currentState = states.get(String.valueOf(split[0].charAt(15)));

		for (int j = 0; j < iterations; j++) {
			currentState = states.get(currentState.doIt());
		}

		long total = 0;
		for (Integer integer : tape.values()) {
			total += integer;
		}

		return total;
	}

	private class State {
		private final Pattern IN_STATE_PATTERN = Pattern.compile("In state (.):");
		private final Pattern CURRENT_VALUE_PATTERN = Pattern.compile("If the current value is (.):");
		private final Pattern WRITE_VALUE_PATTERN = Pattern.compile("Write the value (.)");
		private final Pattern MOVE_VALUE_PATTERN = Pattern.compile("Move one slot to the (.+)\\.");
		private final Pattern CONTINUE_VALUE_PATTERN = Pattern.compile("Continue with state (.)");


		final private String state;

		final private int ifValueA;
		final private int writeA;
		final private String directionA;
		final private String nextStateA;

		final private int writeB;
		final private String directionB;
		final private String nextStateB;

		State(String[] strings) {
			Matcher matcher = IN_STATE_PATTERN.matcher(strings[0]);
			matcher.find();
			state = matcher.group(1);

			matcher = CURRENT_VALUE_PATTERN.matcher(strings[1]);
			matcher.find();
			ifValueA = Integer.parseInt(matcher.group(1));
			matcher = WRITE_VALUE_PATTERN.matcher(strings[2]);
			matcher.find();
			writeA = Integer.parseInt(matcher.group(1));
			matcher = MOVE_VALUE_PATTERN.matcher(strings[3]);
			matcher.find();
			directionA = matcher.group(1);
			matcher = CONTINUE_VALUE_PATTERN.matcher(strings[4]);
			matcher.find();
			nextStateA = matcher.group(1);

			matcher = WRITE_VALUE_PATTERN.matcher(strings[6]);
			matcher.find();
			writeB = Integer.parseInt(matcher.group(1));
			matcher = MOVE_VALUE_PATTERN.matcher(strings[7]);
			matcher.find();
			directionB = matcher.group(1);
			matcher = CONTINUE_VALUE_PATTERN.matcher(strings[8]);
			matcher.find();
			nextStateB = matcher.group(1);
		}

		String getState() {
			return state;
		}

		String doIt() {
			if (tape.getOrDefault(tapeIndex, 0) == ifValueA) {
				tape.put(tapeIndex, writeA);
				if (directionA.equals("right")) {
					tapeIndex++;
				}
				else if (directionA.equals("left")) {
					tapeIndex--;
				}
				return nextStateA;
			}
			else {
				tape.put(tapeIndex, writeB);
				if (directionB.equals("right")) {
					tapeIndex++;
				}
				else if (directionB.equals("left")) {
					tapeIndex--;
				}
				return nextStateB;
			}
		}
	}
}

