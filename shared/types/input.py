from typing import TypedDict

class DuolingoWordResponse(TypedDict):
    text: str
    translations: list[str]

class DuolingoResponse(TypedDict):
    learnedLexemes: list[DuolingoWordResponse]