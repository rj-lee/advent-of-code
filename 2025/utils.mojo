from pathlib import Path


def get_input(filename: String):
    day = Path(filename).stem
    with open(f"./input/{day}.txt", encoding="utf-8") as file:
        return file.read()


def main():
    get_input("day01.txt")
