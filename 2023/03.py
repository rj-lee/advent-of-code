import re

import utils

lines = utils.get_input(__file__).splitlines()
height = len(lines)
width = len(lines[0])


def adjacent_chars(row: int, start: int, end: int):
    adjacents = {}

    if row > 0:
        for col in range(max(0, start - 1), min(width, end + 1)):
            adjacents[(row - 1, col)] = lines[row - 1][col]

    if start > 0:
        adjacents[(row, start - 1)] = lines[row][start - 1]

    if end < width:
        adjacents[(row, end)] = lines[row][end]

    if row < height - 1:
        for col in range(max(0, start - 1), min(width, end + 1)):
            adjacents[(row + 1, col)] = lines[row + 1][col]

    return adjacents


def part1():
    total = 0
    for row, line in enumerate(lines):
        parts = re.finditer(r"\d+", line)
        for part in parts:
            start = part.start()
            end = part.end()
            number = int(part.group())
            adjacents = adjacent_chars(row, start, end)
            if any(a != "." for a in adjacents.values()):
                total += number

    print(total)


def part2():
    total = 0
    gears = {}

    for row, line in enumerate(lines):
        parts = re.finditer(r"\d+", line)
        for part in parts:
            start = part.start()
            end = part.end()
            number = int(part.group())
            adjacents = adjacent_chars(row, start, end)
            for coord, symbol in adjacents.items():
                if symbol == "*":
                    if gears.get(coord):
                        total += gears[coord] * number
                    else:
                        gears[coord] = number

    print(total)


if __name__ == "__main__":
    part1()
    part2()

# Elegant 12 line solution
# https://www.reddit.com/r/adventofcode/comments/189m3qw/comment/kbs9g3g
# https://topaz.github.io/paste/#XQAAAQAcAgAAAAAAAAA0m0pnuFI8c+fPp4G3Y5M2miSs3R6AnrKm3fbDkugpdVsCgQOTZL/yzxGwy/N08UeBslxE7G36XluuSq4Y/2FE0+4nPcdj9lpkrrjBk5HRCFLEKuPjUV8tYPx04VDoJ1c6yyLzScmAGwNvzpPoqb5PkRyyy4dSEcuEDe/k0/U7h7pZVh4eTrNAIPsTNZohcltxuwuA4lrZSN37i0QZiufFpvLVyhV/dLBnmSr+2jwFcFE+W6OEIFQxK6MIJ2z7TWKj8lg6yV4yhJzTm+c+QHh2omzhGVLd2WdcHdhjmCyC+Btbr3yCqemYb/6tMUvz8VchnyHstx7QKKeLVmTOEyYqHH/qRDhlKXSQ23RWuPibCf4quQUPGpPDRsH4KITzLbIUVUdssnSp6ffcHO+dAISdzBOiznl5/+PI+jE=
