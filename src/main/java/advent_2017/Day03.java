package advent_2017;


public class Day03 {
	public int part1(int target) {
		int size = getSize(target)+2;
		int offset = size/2;
		int[][] memory = new int[size][size];
		int x = offset;
		int y = offset;
		String heading = "r";

		for (int i = 1; i < target; i++) {
			memory[x][y] = i;
			System.out.println(i + " " + x + " " + y + " " + heading);

			switch (heading) {
				case "r":
					x++;
					break;
				case "u":
					y++;
					break;
				case "l":
					x--;
					break;
				case "d":
					y--;
					break;
			}

			switch (heading) {
				case "r":
					if (memory[x][y + 1] == 0) {
						heading = "u";
					}
					break;
				case "u":
					if (memory[x - 1][y] == 0) {
						heading = "l";
					}
					break;
				case "l":
					if (memory[x][y - 1] == 0) {
						heading = "d";
					}
					break;
				case "d":
					if (memory[x + 1][y] == 0) {
						heading = "r";
					}
					break;
			}
		}

		// x and y are now at the target space and there absolute sum from 0,0 is the distance
		return Math.abs(x-offset) + Math.abs(y-offset);
	}

	public int part2(int target) {
		int size = getSize(target)+2;
		int offset = size/2;
		int[][] memory = new int[size][size];
		int x = offset;
		int y = offset;
		String heading = "r";

		memory[x][y] = 1; // inital value, calculations still hold
		for (int i = 1; i < target; i++) {
			memory[x][y] = calculateSum(memory, x, y);

			if (memory[x][y] > target) {
				return(memory[x][y]);
			}

			switch (heading) {
				case "r":
					x++;
					break;
				case "u":
					y++;
					break;
				case "l":
					x--;
					break;
				case "d":
					y--;
					break;
			}

			switch (heading) {
				case "r":
					if (memory[x][y + 1] == 0) {
						heading = "u";
					}
					break;
				case "u":
					if (memory[x - 1][y] == 0) {
						heading = "l";
					}
					break;
				case "l":
					if (memory[x][y - 1] == 0) {
						heading = "d";
					}
					break;
				case "d":
					if (memory[x + 1][y] == 0) {
						heading = "r";
					}
					break;
			}
		}

		// x and y are now at the target space and there absolute sum from 0,0 is the distance
		return Math.abs(x-offset) + Math.abs(y-offset);
	}

	private int calculateSum(int[][] memory, int x, int y) {
		int sum = 0;
		for (int i = -1; i <= 1; i++) {
			for (int j = -1; j <= 1; j++) {
				sum+=memory[x+i][y+j];
			}
		}
		return sum;
	}

	private int getSize(int target) {
		int size = 0;
		while (size*size < target)
			size++;
		return size;
	}
}

