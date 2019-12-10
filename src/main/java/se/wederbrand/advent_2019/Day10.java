package se.wederbrand.advent_2019;

import java.util.HashMap;
import java.util.Map;

public class Day10 {
	HashMap<String, String> map = new HashMap<>();

	public Day10(String mapString) {
		String[] yRows = mapString.split(System.lineSeparator());
		for (int y = 0; y < yRows.length; y++) {
			String[] xValues = yRows[y].split("");
			for (int x = 0; x < xValues.length; x++) {
				if (xValues[x].equalsIgnoreCase("#")) {
					map.put(getKey(x, y), xValues[x]);
				}
			}
		}
	}

	public int findBest() {
		int max = 0;
		for (String key : map.keySet()) {
			int x = Integer.parseInt(key.split(",")[0]);
			int y = Integer.parseInt(key.split(",")[1]);
			int numberDetected = getNumberDetected(x, y);
			if (numberDetected > max) {
				max = numberDetected;
			}
		}

		return max;
	}

	public int getNumberDetected(int targetX, int targetY) {
		HashMap<String, String> mapCopy = new HashMap<>(map);

		for (String inner : map.keySet()) {
			// go through all asteroids
			int innerX = Integer.parseInt(inner.split(",")[0]);
			int innerY = Integer.parseInt(inner.split(",")[1]);

			if (innerX == targetX && innerY == targetY) {
				continue;
			}

			double innerDistance = Math.hypot(Math.abs(targetX - innerX), Math.abs(targetY - innerY));
			double innerAngel = Math.atan2(((double) targetY - innerY), ((double) targetX - innerX));

			// go through them again
			for (String outer : map.keySet()) {
				if (outer.equalsIgnoreCase(inner)) {
					continue;
				}
				int outerX = Integer.parseInt(outer.split(",")[0]);
				int outerY = Integer.parseInt(outer.split(",")[1]);

				double outerDistance = Math.hypot(Math.abs(targetX - outerX), Math.abs(targetY - outerY));
				double outerAngel = Math.atan2(((double) targetY - outerY), ((double) targetX - outerX));

				if (outerDistance > innerDistance && innerAngel == outerAngel) {
					// erase the ones that can't be seen
					mapCopy.remove(getKey(outerX, outerY));
				}
			}
		}

		return mapCopy.size() - 1;
	}

	private String getKey(int x, int y) {
		return x + "," + y;
	}

}
