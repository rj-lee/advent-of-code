from math import prod

import utils

lines = utils.get_input(__file__)


def part1():
    boxes = lines.splitlines()
    circuits = [{box} for box in boxes]
    distances = []
    for i, box1 in enumerate(boxes):
        for box2 in boxes[i + 1 :]:
            distance = sum(
                (int(a) - int(b)) ** 2 for a, b in zip(box1.split(","), box2.split(","))
            )
            distances.append((distance, box1, box2))
    distances.sort(key=lambda x: x[0])
    for _ in range(10):
        distance, box1, box2 = distances.pop(0)
        circuit1 = next(c for c in circuits if box1 in c)
        circuit2 = next(c for c in circuits if box2 in c)
        if circuit1 is not circuit2:
            circuit1.update(circuit2)
            circuits.remove(circuit2)
    circuits.sort(key=len, reverse=True)
    return prod(len(c) for c in circuits[:3])


def part2():
    boxes = lines.splitlines()
    circuits = [{box} for box in boxes]
    distances = []
    for i, box1 in enumerate(boxes):
        for box2 in boxes[i + 1 :]:
            distance = sum(
                (int(a) - int(b)) ** 2 for a, b in zip(box1.split(","), box2.split(","))
            )
            distances.append((distance, box1, box2))
    distances.sort(key=lambda x: x[0])
    for distance, box1, box2 in distances:
        circuit1 = next(c for c in circuits if box1 in c)
        circuit2 = next(c for c in circuits if box2 in c)
        if circuit1 is not circuit2:
            circuit1.update(circuit2)
            circuits.remove(circuit2)
        if len(circuits) == 1:
            return int(box1.split(",")[0]) * int(box2.split(",")[0])


if __name__ == "__main__":
    print(part1())
    print(part2())
