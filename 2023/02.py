import utils

lines = utils.get_input(__file__).splitlines()


def part1():
    red = 12
    green = 13
    blue = 14
    sum_ids = 0

    for id_, line in enumerate(lines):
        game = line.partition(": ")[2]
        sets = game.split("; ")
        for set_ in sets:
            cubes = set_.split(", ")
            for cube in cubes:
                count = int(cube.split(" ")[0])
                if "red" in cube:
                    if count > red:
                        break
                elif "green" in cube:
                    if count > green:
                        break
                elif "blue" in cube:
                    if count > blue:
                        break
            else:
                continue
            break

        else:
            sum_ids += id_ + 1

    print(sum_ids)


def part2():
    sum_powers = 0

    for line in lines:
        game = line.partition(": ")[2]
        sets = game.split("; ")
        red = 0
        green = 0
        blue = 0
        for set_ in sets:
            cubes = set_.split(", ")
            for cube in cubes:
                count = int(cube.split(" ")[0])
                if "red" in cube:
                    red = max(count, red)
                elif "green" in cube:
                    green = max(count, green)
                elif "blue" in cube:
                    blue = max(count, blue)

        power = red * green * blue
        sum_powers += power

    print(sum_powers)


if __name__ == "__main__":
    part1()
    part2()
