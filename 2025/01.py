import utils

lines = utils.get_input(__file__)


def part1():
    password = 0
    dial = 50
    rotations = lines.splitlines()
    for rotation in rotations:
        direction = rotation[0]
        clicks = int(rotation[1:])
        delta = -clicks if direction == "L" else clicks
        dial += delta
        dial %= 100
        if dial == 0:
            password += 1
    return password


def part2():
    password = 0
    dial = 50
    rotations = lines.splitlines()
    for rotation in rotations:
        direction = rotation[0]
        clicks = int(rotation[1:])
        for _ in range(clicks):
            dial += -1 if direction == "L" else 1
            dial %= 100
            if dial == 0:
                password += 1
        dial %= 100
    return password


if __name__ == "__main__":
    print(part1())
    print(part2())
