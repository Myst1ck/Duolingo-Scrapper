import json
import csv

from shared.types.input import DuolingoResponse


def read_input(path: str) -> DuolingoResponse:
    with open(path, 'r') as file:
        data: DuolingoResponse = json.load(file)
        return data


def read_output(path: str) -> list[list[str]]:
    with open(path, 'r') as file:
        reader = csv.reader(file)
        data = list(reader)
        return data
