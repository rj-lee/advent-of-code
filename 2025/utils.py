import contextlib
import time
from pathlib import Path


def get_input(filename: str):
    day = Path(filename).stem
    with open(f"./input/{day}.txt", encoding="utf-8") as file:
        return file.read()


@contextlib.contextmanager
def perf_counter():
    start = time.perf_counter()
    yield
    end = time.perf_counter()
    print(f"Took {end - start:.6f} seconds")
