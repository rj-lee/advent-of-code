import utils

lines = utils.get_input(__file__).splitlines()


def part1():
    total = 0

    for line in lines:
        card = line.split(": ")[1].split(" | ")
        winning = set(card[0].split())
        scratched = set(card[1].split())
        if matches := len(winning & scratched):
            total += 2 ** (matches - 1)

    print(total)


def part2():
    cards = [1] * len(lines)

    for i, line in enumerate(lines):
        card = line.split(": ")[1].split(" | ")
        winning = set(card[0].split())
        scratched = set(card[1].split())
        matches = len(winning & scratched)
        for j in range(matches):
            cards[i + j + 1] += cards[i]

    print(sum(cards))


if __name__ == "__main__":
    part1()
    part2()
