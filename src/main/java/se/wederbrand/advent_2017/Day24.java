package se.wederbrand.advent_2017;

import java.util.ArrayList;

public class Day24 {
	public long part1(String input) {
		String[] split = input.split(System.lineSeparator());
		ArrayList<Component> components = new ArrayList<>(split.length);
		for (String s : split) {
			components.add(new Component(s));
		}

		return getMax(0, components);
	}

	public long part2(String input) {
		String[] split = input.split(System.lineSeparator());
		ArrayList<Component> components = new ArrayList<>(split.length);
		for (String s : split) {
			components.add(new Component(s));
		}

		return getMax2(0, components).getStrength();
	}

	private long getMax(int i, ArrayList<Component> components) {
		ArrayList<Component> subComponents = getMatchingComponents(i, components);
		long localMax = 0;
		for (Component component : subComponents) {
			components.remove(component);
			long max = component.value() + getMax(component.otherSide(i), components);
			if (max > localMax) {
				localMax = max;
			}
			components.add(component);
		}
		return localMax;
	}

	private Max getMax2(int i, ArrayList<Component> components) {
		ArrayList<Component> subComponents = getMatchingComponents(i, components);
		Max localMax = new Max(0,0);
		for (Component component : subComponents) {
			components.remove(component);
			Max max = getMax2(component.otherSide(i), components).addValue(component.value());
			if (max.getLength() > localMax.getLength() || (max.getLength() == localMax.getLength() && max.getStrength() > localMax.getStrength())) {
				localMax = max;
			}
			components.add(component);
		}
		return localMax;
	}

	private ArrayList<Component> getMatchingComponents(int i, ArrayList<Component> components) {
		ArrayList<Component> matchingComponents = new ArrayList<>(components.size());

		for (Component component : components) {
			if (component.matches(i)) {
				matchingComponents.add(component);
			}
		}

		return matchingComponents;
	}


	private class Component {
		final private int a;
		final private int b;

		public Component(String s) {
			a = Integer.parseInt(s.split("/")[0]);
			b = Integer.parseInt(s.split("/")[1]);
		}

		public int otherSide(int i) {
			if (i == a) {
				return b;
			}
			else {
				return a;
			}
		}

		public boolean matches(int i) {
			return a == i || b == i;
		}

		public long value() {
			return a + b;
		}
	}

	private class Max {
		private final int length;
		private final long strength;

		public Max(int length, long strength) {
			this.length = length;
			this.strength = strength;
		}

		public int getLength() {
			return length;
		}

		public long getStrength() {
			return strength;
		}

		public Max addValue(long value) {
			return new Max(length +1 , strength + value);
		}
	}
}

