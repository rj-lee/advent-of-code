from pathlib import Path


def get_input(filename: str):
    day = Path(filename).stem
    with open(f"./input/{day}.txt", encoding="utf-8") as file:
        return file.read()
