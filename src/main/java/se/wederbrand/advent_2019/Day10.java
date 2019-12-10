package se.wederbrand.advent_2019;

import java.util.Comparator;
import java.util.HashMap;
import java.util.SortedSet;
import java.util.TreeSet;

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
			double innerAngel = getDegrees(targetX, targetY, innerX, innerY);

			// go through them again
			for (String outer : map.keySet()) {
				if (outer.equalsIgnoreCase(inner)) {
					continue;
				}
				int outerX = Integer.parseInt(outer.split(",")[0]);
				int outerY = Integer.parseInt(outer.split(",")[1]);

				double outerDistance = Math.hypot(Math.abs(targetX - outerX), Math.abs(targetY - outerY));
				double outerAngel = getDegrees(targetX, targetY, outerX, outerY);

				if (outerDistance > innerDistance && innerAngel == outerAngel) {
					// erase the ones that can't be seen
					mapCopy.remove(getKey(outerX, outerY));
				}
			}
		}

		return mapCopy;
	}

	public static double getDegrees(int startX, int startY, int endX, int endY) {
		double rad = Math.atan2((endY - startY), (endX - startX));
		double deg = rad * (180 / Math.PI);
		deg += 90;
		if (deg > 360) {
			deg -=360;
		}
		if (deg < 0) {
			deg += 360;
		}
		return deg;
	}

	private String getKey(int x, int y) {
		return x + "," + y;
	}

	public int vaporize200(int x, int y) {
		int removed = 0;
		while (true) {
			// start up
			double angel = getDegrees(0, 0, 0, -1);
			// find all visible
			HashMap<String, String> visibleMap = getVisibleMap(x, y);
			SortedSet<String> sortedMap = new TreeSet<>(new Comparator<String>() {
				@Override
				public int compare(String o1, String o2) {
					int targetX1 = Integer.parseInt(o1.split(",")[0]);
					int targetY1 = Integer.parseInt(o1.split(",")[1]);
					double degrees1 = getDegrees(x, y, targetX1, targetY1);

					int targetX2 = Integer.parseInt(o2.split(",")[0]);
					int targetY2 = Integer.parseInt(o2.split(",")[1]);
					double degrees2 = getDegrees(x, y, targetX2, targetY2);

					return Double.compare(degrees1, degrees2);
				}
			});

			sortedMap.addAll(visibleMap.keySet());

			for (String key : sortedMap) {
				map.remove(key);
				removed++;
				if (removed == 200) {
					String[] split = key.split(",");
					int resultX = Integer.parseInt(split[0]);
					int resultY = Integer.parseInt(split[1]);

					return resultX*100+resultY;
				}
			}
		}
	}

}
