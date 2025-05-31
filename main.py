import os

from filterer.filterer import filter_values
from reader.reader import read_input, read_output
from writer.writer import write_data

INPUT_PATH = os.path.join(os.path.dirname(__file__), 'values/input.json')
OUTPUT_PATH = os.path.join(os.path.dirname(__file__), 'values/output.csv')

SHOULD_AUTOMATICALLY_TAKE_LATEST_WORD = True
LATEST_WORD = 'endlich'


def main() -> None:
    initial_input = read_input(INPUT_PATH)
    
    last_word = LATEST_WORD if not SHOULD_AUTOMATICALLY_TAKE_LATEST_WORD else read_output(OUTPUT_PATH)[0][0]
    data = filter_values(initial_input, {'last_word': last_word}) # type: ignore
    
    write_data(OUTPUT_PATH, data)


if __name__ == '__main__':
    main()
