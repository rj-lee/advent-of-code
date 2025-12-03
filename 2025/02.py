import utils

lines = utils.get_input(__file__)


def part1():
    total = 0
    id_ranges = lines.split(",")
    for id_range in id_ranges:
        start, end = id_range.split("-")
        for id_int in range(int(start), int(end) + 1):
            id_string = str(id_int)
            length = len(id_string)
            if id_string[: length // 2] == id_string[length // 2 :]:
                total += id_int
    return total


def part2():
    total = 0
    id_ranges = lines.split(",")
    for id_range in id_ranges:
        start, end = id_range.split("-")
        for id_int in range(int(start), int(end) + 1):
            id_string = str(id_int)
            length = len(id_string)
            for part_length in range(1, length // 2 + 1):
                chunks = [
                    id_string[i : i + part_length]
                    for i in range(0, length, part_length)
                ]
                if len(chunks) > 1 and len(set(chunks)) == 1:
                    total += id_int
                    break
    return total


if __name__ == "__main__":
    print(part1())
    print(part2())
