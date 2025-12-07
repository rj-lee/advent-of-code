import utils

lines = utils.get_input(__file__)


def part1():
    rows = [list(line) for line in lines.splitlines()]
    total = 0
    for y, row in enumerate(rows):
        if y == len(rows) - 1:
            break
        beams = [x for x, char in enumerate(row) if char in "S|"]
        for next_y, next_char in enumerate(rows[y + 1]):
            if next_y in beams:
                if next_char == "^":
                    rows[y + 1][next_y - 1] = "|"
                    rows[y + 1][next_y + 1] = "|"
                    total += 1
                else:
                    rows[y + 1][next_y] = "|"
    return total


def part2():
    rows = lines.splitlines()
    width = len(rows[0])
    beams = [0] * width
    beams[rows[0].index("S")] = 1
    for row in rows[1:]:
        for i in range(width):
            if row[i] == "^" and beams[i] > 0:
                beams[i - 1] += beams[i]
                beams[i + 1] += beams[i]
                beams[i] = 0
    return sum(beams)


if __name__ == "__main__":
    print(part1())
    print(part2())
