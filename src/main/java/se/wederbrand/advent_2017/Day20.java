package se.wederbrand.advent_2017;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

public class Day20 {

	public int part1(String input) {
		String[] particles = input.split(System.lineSeparator());

		int minAccIndex=-1;
		int minAcc=Integer.MAX_VALUE;
		for (int i = 0; i < particles.length; i++) {
			String particle = particles[i];
			String[] split = particle.split(">,");
			String p = split[0].trim();
			String v = split[1].trim();
			String a = split[2].trim();

			String[] aSplit = a.substring(3, a.length() - 1).split(",");
			int aX = Integer.parseInt(aSplit[0]);
			int aY = Integer.parseInt(aSplit[1]);
			int aZ = Integer.parseInt(aSplit[2]);

			int acc = Math.abs(aX) + Math.abs(aY) + Math.abs(aZ);
			if (acc < minAcc){
				minAccIndex = i;
				minAcc = acc;
			}
		}

		return minAccIndex;
	}

	public int part2(String input) {
		String[] particlesInput = input.split(System.lineSeparator());

		List<Particle> particles = new ArrayList<>();
		for (String particle : particlesInput) {
			String[] split = particle.split(">,");
			String p = split[0].trim();
			String v = split[1].trim();
			String a = split[2].trim();

			String[] pSplit = p.substring(3, p.length()).split(",");
			int pX = Integer.parseInt(pSplit[0].trim());
			int pY = Integer.parseInt(pSplit[1].trim());
			int pZ = Integer.parseInt(pSplit[2].trim());

			String[] vSplit = v.substring(3, v.length()).split(",");
			int vX = Integer.parseInt(vSplit[0].trim());
			int vY = Integer.parseInt(vSplit[1].trim());
			int vZ = Integer.parseInt(vSplit[2].trim());

			String[] aSplit = a.substring(3, a.length() - 1).split(",");
			int aX = Integer.parseInt(aSplit[0].trim());
			int aY = Integer.parseInt(aSplit[1].trim());
			int aZ = Integer.parseInt(aSplit[2].trim());

			particles.add(new Particle(pX, pY, pZ, vX, vY, vZ, aX, aY, aZ));
		}

		particles.sort(Comparator.comparing(Particle::acceleration).reversed());

		int escapers = 0;

		outerLoop:
		while (particles.size() > 0) {
			System.out.println(particles.size());
			// all move
			for (Particle particle : particles) {
				particle.step();
			}

			// find all that collide and remove them
			for (Particle one : particles) {
				for (Particle two : particles) {
					if (one != two && one.hasCollided(two)) {
						System.out.println("collision!");
						one.setCollided();
						two.setCollided();
					}
				}
			}

			particles.removeIf(Particle::isCollided);

			// find all that have the highest acceleration and if they are also the furthest and heading away delete them
			if (particles.size() == 0) {
				continue;
			}
			Particle highestAcc = particles.get(0);
			if (!highestAcc.headingAway()) {
				// not heading away
				continue;
			}

			for (int i = 1; i < particles.size(); i++) {
				Particle particle = particles.get(i);
				if (highestAcc.distance() < particle.distance()) {
					// currently not furthest away
					continue outerLoop;
				}
			}

			particles.remove(0);
			escapers++;
		}


		return escapers;
	}


	private class Particle {
		private int px;
		private int py;
		private int pz;
		private int vx;
		private int vy;
		private int vz;
		private final int ax;
		private final int ay;
		private final int az;
		private boolean collided;

		public Particle(int px, int py, int pz, int vx, int vy, int vz, int ax, int ay, int az) {
			this.px = px;
			this.py = py;
			this.pz = pz;
			this.vx = vx;
			this.vy = vy;
			this.vz = vz;
			this.ax = ax;
			this.ay = ay;
			this.az = az;
		}

		public void step() {
			vx += ax;
			vy += ay;
			vz += az;

			px += vx;
			py += vy;
			pz += vz;
		}

		public double distance() {
			return Math.sqrt(px * px + py * py + pz * pz);
		}

		public double acceleration() {
			return Math.sqrt(ax * ax + ay * ay + az * az);
		}

		public boolean hasCollided(Particle two) {
			return px==two.px && py==two.py && pz==two.pz;
		}

		public boolean headingAway() {
			// not yet travelling in the correct direction
			if (vx*ax < 0) return false;
			if (vy*ay < 0) return false;
			if (vz*az < 0) return false;

			// not yet travelled passed origo
			if (vx*px < 0) return false;
			if (vy*py < 0) return false;
			if (vz*pz < 0) return false;

			// heading away
			return true;
		}

		public boolean isCollided() {
			return collided;
		}

		public void setCollided() {
			this.collided = true;
		}
	}

}

