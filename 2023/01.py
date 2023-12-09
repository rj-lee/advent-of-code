import re

import utils

lines = utils.get_input(__file__).splitlines()


def part1():
    total = 0

    for line in lines:
        numbers = re.sub(r"[^\d]", "", line)
        total += int(numbers[0]) * 10 + int(numbers[-1])

    print(total)


def part2():
    total = 0
    replacements = {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }

    for line in lines:
        buffer = ""
        for char in line:
            buffer += char
            if digit := replacements.get(buffer):
                line = line.replace(buffer, digit + buffer[-1], 1)
                buffer = buffer[-1]
            elif not any(key.startswith(buffer) for key in replacements):
                buffer = buffer[1:]

        numbers = re.sub(r"[^\d]", "", line)
        total += int(numbers[0]) * 10 + int(numbers[-1])

    print(total)


if __name__ == "__main__":
    part1()
    part2()
