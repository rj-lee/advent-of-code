import utils

lines = utils.get_input(__file__)


def part1():
    sections = lines.split("\n\n")
    shapes = [[list(row) for row in shape.splitlines()[1:]] for shape in sections[:-1]]
    regions = [
        (
            [int(v) for v in dimensions.split("x")],
            [int(v) for v in requirements.split()],
        )
        for dimensions, requirements in (
            region.split(": ") for region in sections[-1].splitlines()
        )
    ]
    total = 0
    for (x, y), requirements in regions:
        region = [["."] * x for _ in range(y)]
        # solution = solve(shapes, region, requirements)
        if x * y >= 9 * sum(requirements):
            total += 1
    return total


if __name__ == "__main__":
    with utils.perf_counter():
        print(part1())
