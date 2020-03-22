import argparse


def create_array(input_str):
    return [float(item) for item in input_str.split(',')]


def check_common(input_first, input_second):
    res = []
    for el in input_first: # loop for first array
        if el in input_second: # check if element present in second array if yes -->
            res.append(el) # add element in to result array
    input_first_common = sum(input_first)/len(input_first)
    print(input_first_common)
    input_second_common = sum(input_second)/len(input_second)
    print(input_second_common)
    if input_first_common in input_second:
        res.append(input_first_common)
    if input_second_common in input_first:
        res.append(input_second_common)

    print(res)


if __name__ == '__main__':
    '''
   In program we have  two lists with Unknown length O(n) Worst case O(n log(n))

   '''
    new_parser = argparse.ArgumentParser()
    new_parser.add_argument('-a1',
                            '--array1',
                            # action='store_true',
                            help='Tokens separate by ","')
    new_parser.add_argument('-a2',
                            '--array2',
                            # action='store_true',
                            help='Input str')

    args = new_parser.parse_args()
    print(args)
    a1 = create_array(args.array1)
    a2 = create_array(args.array2)
    check_common(a1, a2)
