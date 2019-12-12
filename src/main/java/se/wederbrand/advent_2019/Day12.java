package se.wederbrand.advent_2019;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

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
			changeVelocity();
			moveAll();
		}

		long energy = 0L;
		for (Moon moon : moons) {
			energy += moon.energy();
		}
		return energy;
	}

	private void moveAll() {
		for (Moon moon : moons) {
			moon.move();
		}
	}

	private void changeVelocity() {
		for (Moon moon1 : moons) {
			for (Moon moon2 : moons) {
				moon1.gravity(moon2);
			}
		}
	}

	private class Moon {
		int x;
		int y;
		int z;

		int vx = 0;
		int vy = 0;
		int vz = 0;
		public Moon(String moonData) {
			// <x=-1, y=0, z=2>
			String[] split = moonData.split("[,=>]");
			this.x = Integer.parseInt(split[1]);
			this.y = Integer.parseInt(split[3]);
			this.z = Integer.parseInt(split[5]);
		}

		public void gravity(Moon moon2) {
			if (x < moon2.x) {
				vx++;
			}
			else if (x > moon2.x) {
				vx--;
			}
			if (y < moon2.y) {
				vy++;
			}
			else if (y > moon2.y) {
				vy--;
			}
			if (z < moon2.z) {
				vz++;
			}
			else if (z > moon2.z) {
				vz--;
			}
		}

		public void move() {
			x += vx;
			y += vy;
			z += vz;
		}

		public long energy() {
			return (abs(x) + abs(y) + abs(z)) * (abs(vx) + abs(vy) + abs(vz)) ;
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
	}
}
