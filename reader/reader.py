import json

from shared.types.input import DuolingoResponse

def read_input(path: str) -> DuolingoResponse:
    with open(path, 'r') as file:
        data: DuolingoResponse = json.load(file)
        return data