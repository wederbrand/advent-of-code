package se.wederbrand.advent_2019;

import java.util.HashMap;

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
		int bestX = 0;
		int bestY = 0;
		for (String key : map.keySet()) {
			int x = Integer.parseInt(key.split(",")[0]);
			int y = Integer.parseInt(key.split(",")[1]);
			int numberDetected = getNumberDetected(x, y);
			if (numberDetected > max) {
				max = numberDetected;
				bestX = x;
				bestY = y;
			}
		}

		System.out.println(bestX + " " + bestY);

		return max;
	}

	public int getNumberDetected(int targetX, int targetY) {
		HashMap<String, String> visibleMap = getVisibleMap(targetX, targetY);
		return visibleMap.size() - 1;
	}

	private HashMap<String, String> getVisibleMap(int targetX, int targetY) {
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

		return mapCopy;
	}

	private String getKey(int x, int y) {
		return x + "," + y;
	}

	public int vaporize200(int x, int y) {
		int removed = 0;
		while (true) {
			// start up
			double angel = Math.atan2(-1.0, 0.0);
			// find all visible
			HashMap<String, String> visibleMap = getVisibleMap(x, y);

			do {
				// jump to the next possible
				double[] nextTarget = findNextTarget(x, y, angel, visibleMap);
				// remove it
				int targetX = (int) nextTarget[0];
				int targetY = (int) nextTarget[1];
				double nextAngel = nextTarget[2];
				visibleMap.remove(getKey(targetX, targetY));
				map.remove(getKey(targetX, targetY));
				removed++;
				if (removed == 200) {
					return targetX*100 + targetY;
				}
				angel = nextAngel;
			} while (angel < Math.atan2(-1.0, 0.0));
		}
	}

	private double[] findNextTarget(int targetX, int targetY, double targetAngel, HashMap<String, String> visibleMap) {
		double minAngelDiff = Double.MAX_VALUE;
		int nextTargetX = 0;
		int nextTargetY = 0;
		double nextAngel = 0.0;

		for (String key : visibleMap.keySet()) {
			int x = Integer.parseInt(key.split(",")[0]);
			int y = Integer.parseInt(key.split(",")[1]);
			double angel = Math.atan2(((double) targetY - y), ((double) targetX - x));

			if (angel-targetAngel < minAngelDiff) {
				minAngelDiff = angel-targetAngel;
				nextTargetX = x;
				nextTargetY = y;
				nextAngel = angel;
			}
		}

		return new double[]{nextTargetX, nextTargetY, nextAngel};
	}
}
