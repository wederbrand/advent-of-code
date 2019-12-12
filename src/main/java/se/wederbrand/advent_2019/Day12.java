package se.wederbrand.advent_2019;

import java.util.ArrayList;
import java.util.List;

import static java.lang.Math.abs;

public class Day12 {
	List<Moon> moons = new ArrayList<>();

	public Day12(String input) {
		String[] split = input.split(System.lineSeparator());
		for (String moonData : split) {
			moons.add(new Moon(moonData));
		}
	}

	public long part1(int ticks) {
		for (int i = 0; i < ticks; i++) {
			// change velocity
			for (Moon moon1 : moons) {
				for (Moon moon2 : moons) {
					moon1.gravityX(moon2);
					moon1.gravityY(moon2);
					moon1.gravityZ(moon2);
				}
			}

			for (Moon moon : moons) {
				moon.moveX();
				moon.moveY();
				moon.moveZ();
			}

		}

		long energy = 0L;
		for (Moon moon : moons) {
			energy += moon.energy();
		}
		return energy;
	}

	public long part2() {
		long xSteps = 0;
		while (true) {
			xSteps++;
			if (xSteps % 1_000_000 == 0) {
				System.out.println(xSteps / 1_000_000);
			}
			for (Moon moon1 : moons) {
				for (Moon moon2 : moons) {
					moon1.gravityX(moon2);
				}
			}

			boolean identical = true;
			for (Moon moon : moons) {
				moon.moveX();
				if (!moon.atStartX()) {
					identical = false;
				}
			}

			if (identical) {
				break;
			}
		}

		long ySteps = 0;
		while (true) {
			ySteps++;
			if (ySteps % 1_000_000 == 0) {
				System.out.println(ySteps / 1_000_000);
			}
			for (Moon moon1 : moons) {
				for (Moon moon2 : moons) {
					moon1.gravityY(moon2);
				}
			}

			boolean identical = true;
			for (Moon moon : moons) {
				moon.moveY();
				if (!moon.atStartY()) {
					identical = false;
				}
			}

			if (identical) {
				break;
			}
		}

		long zSteps = 0;
		while (true) {
			zSteps++;
			if (zSteps % 1_000_000 == 0) {
				System.out.println(zSteps / 1_000_000);
			}
			for (Moon moon1 : moons) {
				for (Moon moon2 : moons) {
					moon1.gravityZ(moon2);
				}
			}

			boolean identical = true;
			for (Moon moon : moons) {
				moon.moveZ();
				if (!moon.atStartZ()) {
					identical = false;
				}
			}

			if (identical) {
				break;
			}
		}

		// 2427676415834048 is too high;
		// 303459551979256 is correct
		return lcm(xSteps, lcm(ySteps, zSteps));
	}

	long gcd(long a, long b) {
		if (a == 0) {
			return b;
		}
		return gcd(b % a, a);
	}

	long lcm(long a, long b) {
		return (a * b) / gcd(a, b);
	}

	private class Moon {
		long x;
		long y;
		long z;

		long origX;
		long origY;
		long origZ;

		long vx = 0;
		long vy = 0;
		long vz = 0;

		public Moon(String moonData) {
			// <x=-1, y=0, z=2>
			String[] split = moonData.split("[,=>]");
			this.x = Integer.parseInt(split[1]);
			this.y = Integer.parseInt(split[3]);
			this.z = Integer.parseInt(split[5]);

			this.origX = Integer.parseInt(split[1]);
			this.origY = Integer.parseInt(split[3]);
			this.origZ = Integer.parseInt(split[5]);
		}

		public void gravityX(Moon moon2) {
			if (x < moon2.x) {
				vx++;
			}
			else if (x > moon2.x) {
				vx--;
			}
		}

		public void gravityY(Moon moon2) {
			if (y < moon2.y) {
				vy++;
			}
			else if (y > moon2.y) {
				vy--;
			}
		}

		public void gravityZ(Moon moon2) {
			if (z < moon2.z) {
				vz++;
			}
			else if (z > moon2.z) {
				vz--;
			}
		}

		public void moveX() {
			x += vx;
		}

		public void moveY() {
			y += vy;
		}

		public void moveZ() {
			z += vz;
		}

		public long energy() {
			return (abs(x) + abs(y) + abs(z)) * (abs(vx) + abs(vy) + abs(vz));
		}

		@Override
		public String toString() {
			return "Moon{" +
				"x=" + x +
				", y=" + y +
				", z=" + z +
				", vx=" + vx +
				", vy=" + vy +
				", vz=" + vz +
				'}';
		}

		public boolean atStartX() {
			if (vx != 0) {
				return false;
			}

			if (x != origX) {
				return false;
			}

			return true;
		}

		public boolean atStartY() {
			if (vy != 0) {
				return false;
			}

			if (y != origY) {
				return false;
			}

			return true;
		}

		public boolean atStartZ() {
			if (vz != 0) {
				return false;
			}

			if (z != origZ) {
				return false;
			}

			return true;
		}
	}
}
