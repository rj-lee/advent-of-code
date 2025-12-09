from math import prod

import utils

lines = utils.get_input(__file__)


def part1():
    tiles = [[int(x) for x in line.split(",")] for line in lines.splitlines()]
    areas = []
    for i, tile1 in enumerate(tiles):
        for tile2 in tiles[i + 1 :]:
            areas.append(prod(abs(a - b) + 1 for a, b in zip(tile1, tile2)))
    return max(areas)


def edges(tile1, tile2):
    x1, y1 = tile1
    x2, y2 = tile2
    xl, xr = min(x1, x2), max(x1, x2)
    yb, yt = min(y1, y2), max(y1, y2)
    for x in range(xl, xr + 1):
        yield x, yb
        yield x, yt
    for y in range(yb + 1, yt):
        yield xl, y
        yield xr, y


def part2():
    red_tiles = [[int(x) for x in line.split(",")] for line in lines.splitlines()]
    rectangles = []
    for i, tile1 in enumerate(red_tiles):
        for tile2 in red_tiles[i + 1 :] + red_tiles[:1]:
            rectangles.append(
                (prod(abs(a - b) + 1 for a, b in zip(tile1, tile2)), tile1, tile2)
            )
    rectangles.sort(key=lambda x: x[0], reverse=True)
    tiles = {}
    for tile1, tile2 in zip(red_tiles, red_tiles[1:] + red_tiles[:1]):
        for x, y in edges(tile1, tile2):
            if not tiles.get(y):
                tiles[y] = [x, x]
            elif x < tiles[y][0]:
                tiles[y][0] = x
            elif x > tiles[y][1]:
                tiles[y][1] = x
    for area, tile1, tile2 in rectangles:
        if all(tiles[y][0] <= x <= tiles[y][1] for x, y in edges(tile1, tile2)):
            return area


if __name__ == "__main__":
    print(part1())
    print(part2())
