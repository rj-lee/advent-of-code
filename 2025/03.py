import utils

lines = utils.get_input(__file__)


def part1():
    total = 0
    for line in lines.splitlines():
        digits = [int(d) for d in line]
        largest = max(digits)
        largest_index = digits.index(largest)
        if largest_index == len(digits) - 1:
            first = max(digits[:-1])
            second = largest
        else:
            first = largest
            second = max(digits[largest_index + 1 :])
        total += int(f"{first}{second}")
    return total


def part2():
    total = 0
    required = 12
    for line in lines.splitlines():
        digits = [int(d) for d in line]
        while len(digits) > required:
            # for i, digit in enumerate(digits):
            #     if i == len(digits) - 1 or digit < digits[i + 1]:
            #         digits.pop(i)
            #         break
            digits.pop(
                next(
                    i
                    for i, digit in enumerate(digits)
                    if i == len(digits) - 1 or digit < digits[i + 1]
                )
            )
        joltage = int("".join(str(d) for d in digits))
        total += joltage
    return total


if __name__ == "__main__":
    print(part1())
    print(part2())
