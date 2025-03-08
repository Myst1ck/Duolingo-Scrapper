from shared.types.filters import FilterOptions
from shared.types.input import DuolingoResponse, DuolingoWordResponse


def filter_values(data: DuolingoResponse, options: FilterOptions | None) -> list[DuolingoWordResponse]:
    filtered_data: list[DuolingoWordResponse] = []
    
    for index, word_response in enumerate(data.get('learnedLexemes')):
        if options != None and options.get('amount') != None and index > options.get('amount'):
            break
        
        if options != None and options.get('last_word') != None and word_response.get('text').lower() == options.get('last_word').lower():
            break
        
        filtered_data.append(word_response)
    
    return filtered_data