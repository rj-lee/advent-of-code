import itertools
import re

import utils
import z3

lines = utils.get_input(__file__)
pattern = re.compile(r"\[(.*)\] (.*) {(.*)}")


def part1():
    total = 0
    # pylint: disable=too-many-nested-blocks
    for line in lines.splitlines():
        match = pattern.match(line)
        assert match
        correct_str, buttons_str, _ = match.groups()
        correct = {i for i, x in enumerate(correct_str) if x == "#"}
        buttons = [
            tuple(int(x) for x in b.strip("()").split(",")) for b in buttons_str.split()
        ]
        presses = 1
        while True:
            combos = itertools.combinations(buttons, presses)
            found = False
            for combo in combos:
                lights = set()
                for button in combo:
                    for light in button:
                        if light in lights:
                            lights.remove(light)
                        else:
                            lights.add(light)
                if lights == correct:
                    total += presses
                    found = True
                    break
            if found:
                break
            presses += 1
    return total


def part2():
    total = 0
    for line in lines.splitlines():
        match = pattern.match(line)
        assert match
        _, buttons_str, joltages_str = match.groups()
        joltages = [int(x) for x in joltages_str.split(",")]
        buttons = {
            tuple(int(x) for x in b.strip("()").split(",")) for b in buttons_str.split()
        }
        presses = [z3.Int(i) for i in range(len(buttons))]
        optimizer = z3.Optimize()
        optimizer.add(p >= 0 for p in presses)
        for j_idx, required in enumerate(joltages):
            contributing = [
                presses[b_idx]
                for b_idx, button in enumerate(buttons)
                if j_idx in button
            ]
            optimizer.add(z3.Sum(contributing) == required)
        total_presses = z3.Sum(presses)
        objective = optimizer.minimize(total_presses)
        optimizer.check()
        min_presses = objective.value()
        assert isinstance(min_presses, z3.IntNumRef)
        total += min_presses.as_long()
    return total


if __name__ == "__main__":
    with utils.perf_counter():
        print(part1())
    with utils.perf_counter():
        print(part2())
