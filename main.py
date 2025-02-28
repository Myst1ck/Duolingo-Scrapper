import os

from filterer.filterer import filter_values
from reader.reader import read_input
from writer.writer import write_data

INPUT_PATH = os.path.join(os.path.dirname(__file__), 'values/input.json')
OUTPUT_PATH = os.path.join(os.path.dirname(__file__), 'values/output.csv')
LATEST_WORD = 'Visum'


def main() -> None:
    intial_input = read_input(INPUT_PATH)
    data = filter_values(intial_input, {'last_word': LATEST_WORD})
    write_data(OUTPUT_PATH, data)


if __name__ == '__main__':
    main()
