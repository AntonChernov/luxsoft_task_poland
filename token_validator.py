import argparse
import itertools


def letters(obj): # separate string to letters
    if isinstance(obj, list):
        # itertools.chain.from_iterable(obj) --> from ['asd', 'sdf'] to ['asdsdf']
        return list(set(itertools.chain.from_iterable(obj)))  # from ['asd', 'sdf'] to ['asdsdf'] but the string
        # contains only unique letters
    elif isinstance(obj, str):
        return list(set(obj))  # obj str contains only unique letters
    else:
        raise ValueError


def validate_token(token_str, input_str):
    if len(input_str) < len(token_str):
        print(False) # if len of tokens example less then input string no need to check so just return False
    else:
        tokens = letters(token_str)
        inputs = letters(input_str)
        if all(elem in tokens for elem in inputs): # if all elements from input we have in tokens so it is good input
            # if no bad input
            print(True)
        else:
            print(False)


if __name__ == '__main__':
    '''
    In program we have 2 two sickles (#8, #22) so in a best case O(n) Worst case O(n^2)
    Also about empty string i believe not good solution use "" as valid value. 
    '''
    new_parser = argparse.ArgumentParser()
    new_parser.add_argument('-t',
                            '--tokens',
                            # action='store_true',
                            help='Tokens separate by ","')
    new_parser.add_argument('-i',
                            '--inputstr',
                            # action='store_true',
                            help='Input str')

    args = new_parser.parse_args()
    validate_token(args.tokens.split(','), args.inputstr)
