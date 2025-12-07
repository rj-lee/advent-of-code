from math import prod

import utils

lines = utils.get_input(__file__)
operations = {"+": sum, "*": prod}


def part1():
    rows = [line.split() for line in lines.splitlines()]
    columns = zip(*rows)
    total = 0
    for column in columns:
        numbers = [int(x) for x in column[:-1]]
        operator = column[-1]
        total += operations[operator](numbers)
    return total


def part2():
    rows = [list(line) for line in lines.splitlines()]
    columns = reversed(list(zip(*rows)))
    total = 0
    numbers = []
    for column in columns:
        if all(c == " " for c in column):
            continue
        number = int("".join(column[:-1]))
        operator = column[-1]
        numbers.append(number)
        if operator in "+*":
            total += operations[operator](numbers)
            numbers = []
    return total


if __name__ == "__main__":
    print(part1())
    print(part2())
