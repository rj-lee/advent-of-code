import utils

lines = utils.get_input(__file__)


def _surrounding_cells(size, x, y):
    cells = []
    for dy in [-1, 0, 1]:
        for dx in [-1, 0, 1]:
            if dx == 0 and dy == 0:
                continue
            nx, ny = x + dx, y + dy
            if 0 <= nx < size[0] and 0 <= ny < size[1]:
                cells.append((nx, ny))
    return cells


def part1():
    grid = []
    for line in lines.splitlines():
        row = []
        for char in line:
            row.append(char)
        grid.append(row)

    total = 0
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == ".":
                continue
            rolls = 0
            for nx, ny in _surrounding_cells((len(grid[0]), len(grid)), x, y):
                if grid[ny][nx] in ("@", "x"):
                    rolls += 1
            if rolls < 4:
                total += 1
                grid[y][x] = "x"
    return total


def part2():
    grid = []
    for line in lines.splitlines():
        row = []
        for char in line:
            row.append(char)
        grid.append(row)

    total = 0
    subtotal = -1
    while subtotal != 0:
        subtotal = 0
        for y, row in enumerate(grid):
            for x, cell in enumerate(row):
                if cell == ".":
                    continue
                rolls = 0
                for nx, ny in _surrounding_cells((len(grid[0]), len(grid)), x, y):
                    if grid[ny][nx] in ("@", "x"):
                        rolls += 1
                if rolls < 4:
                    subtotal += 1
                    total += 1
                    grid[y][x] = "x"
        for y, row in enumerate(grid):
            for x, cell in enumerate(row):
                if cell == "x":
                    grid[y][x] = "."
    return total


if __name__ == "__main__":
    print(part1())
    print(part2())
