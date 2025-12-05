import utils


def _surrounding_cells(
    size: Tuple[Int, Int], x: Int, y: Int
) -> List[Tuple[Int, Int]]:
    cells: List[Tuple[Int, Int]] = []
    for dy in [-1, 0, 1]:
        for dx in [-1, 0, 1]:
            if dx == 0 and dy == 0:
                continue
            nx, ny = x + dx, y + dy
            if 0 <= nx < size[0] and 0 <= ny < size[1]:
                cells.append((nx, ny))
    return cells^


def part1(lines: String):
    grid: List[List[Codepoint]] = []
    for line in lines.split("\n"):
        row: List[Codepoint] = []
        for char in line.codepoints():
            row.append(char)
        grid.append(row^)

    total = 0
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == Codepoint(ord(".")):
                continue
            rolls = 0
            for nx, ny in _surrounding_cells((len(grid[0]), len(grid)), x, y):
                if grid[ny][nx] in (Codepoint(ord("@")), Codepoint(ord("x"))):
                    rolls += 1
            if rolls < 4:
                total += 1
                grid[y][x] = Codepoint(ord("x"))
    print(total)


def part2(lines: String):
    grid: List[List[Codepoint]] = []
    for line in lines.split("\n"):
        row: List[Codepoint] = []
        for char in line.codepoints():
            row.append(char)
        grid.append(row^)

    total = 0
    subtotal = -1
    while subtotal != 0:
        subtotal = 0
        for y, row in enumerate(grid):
            for x, cell in enumerate(row):
                if cell == Codepoint(ord(".")):
                    continue
                rolls = 0
                for nx, ny in _surrounding_cells(
                    (len(grid[0]), len(grid)), x, y
                ):
                    if grid[ny][nx] in ("@", "x"):
                        rolls += 1
                if rolls < 4:
                    subtotal += 1
                    total += 1
                    grid[y][x] = Codepoint(ord("x"))
        for y, row in enumerate(grid):
            for x, cell in enumerate(row):
                if cell == Codepoint(ord("x")):
                    grid[y][x] = Codepoint(ord("."))
    print(total)


fn main():
    lines = utils.get_input(__file__)
    part1(lines)
    part2(lines)
