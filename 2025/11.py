import functools

import utils

lines = utils.get_input(__file__)


def part1():
    devices = {
        device: outputs.split()
        for line in lines.splitlines()
        for device, outputs in [line.split(": ")]
    }
    stack = ["you"]
    paths = 0
    while stack:
        current = stack.pop()
        for output in devices[current]:
            if output == "out":
                paths += 1
            else:
                stack.append(output)
    return paths


def count_paths(devices, start, end):
    @functools.lru_cache
    def dfs(current):
        if current == end:
            return 1
        paths = 0
        for output in devices.get(current, []):
            paths += dfs(output)
        return paths

    return dfs(start)


def part2():
    devices = {
        device: outputs.split()
        for line in lines.splitlines()
        for device, outputs in [line.split(": ")]
    }
    return (
        count_paths(devices, "svr", "dac")
        * count_paths(devices, "dac", "fft")
        * count_paths(devices, "fft", "out")
    ) + (
        count_paths(devices, "svr", "fft")
        * count_paths(devices, "fft", "dac")
        * count_paths(devices, "dac", "out")
    )


if __name__ == "__main__":
    with utils.perf_counter():
        print(part1())
    with utils.perf_counter():
        print(part2())
