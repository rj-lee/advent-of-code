import utils

lines = utils.get_input(__file__)


def part1():
    ranges, available = lines.split("\n\n")
    fresh = []
    for r in ranges.splitlines():
        start, end = r.split("-")
        fresh.append((int(start), int(end)))
    total = 0
    for a in available.splitlines():
        for start, end in fresh:
            if start <= int(a) <= end:
                total += 1
                break
    return total


def part2():
    ranges, _ = lines.split("\n\n")
    fresh = []
    for r in ranges.splitlines():
        start, end = r.split("-")
        fresh.append((int(start), int(end)))
    fresh.sort(key=lambda x: x[0])
    merged = []
    for start, end in fresh:
        if not merged or start > merged[-1][1]:
            merged.append([start, end])
        else:
            merged[-1][1] = max(merged[-1][1], end)
    return sum(end - start + 1 for start, end in merged)


if __name__ == "__main__":
    print(part1())
    print(part2())
