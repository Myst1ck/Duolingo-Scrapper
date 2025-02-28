from shared.types.input import DuolingoWordResponse
import csv
import os


def remove_if_exists(path: str) -> None:
    if os.path.exists(path):
        os.remove(path)


def write_data(path: str, data: list[DuolingoWordResponse]) -> None:
    remove_if_exists(path)

    with open(path, mode='w', newline='') as file:
        writer = csv.writer(file)

        for word_response in data:
            writer.writerow([word_response.get('text'),
                            ', '.join(word_response.get('translations'))])
