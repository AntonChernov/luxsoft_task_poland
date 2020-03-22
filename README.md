#### Python task Implementation
in task use only standard python library
Python version for test task 3.8
Please Note all letters after "#" it is comments do not put it in CLI 

You can create a python virtual environment

```python
python3.8 -m venv <venv name>
```
After activate the  <venv name> yse next command:
```python
source <venv name>/bin/activate
```

Also you can yse python without venv.
```python
python3.8 path to file/token_validator.py -t=ab,bc,c -i=abbcbcf
```


Commands:
Token:
```python
python token_validator.py -t=ab,bc,c -i=abbcbcf # Result: false
python token_validator.py -t=ab,bc,c -i=abbcbc  # Result: true
python token_validator.py -t=ab,bc,c -i=ab      # Result: false
python token_validator.py -t=ab,bc,c -i=abb     # Result: true 
```

Common points:
```python
python common_points.py -a1=4.0,8.2 -a2=6.1,10.3 # Result: [6.1, 8.2]

python common_points.py -a1=1.0,3.5 -a2=3.5,6.1  # Result: [3.5]
```
#### GO Implementation (Not finished! But i tr yto do it :) ):

```go
go run token_validator.go --tokens asd,dsa --inputstr asdfer

go run common_points.go --first 1.0,3.5 --second 3.5,6.1
```



## Task description:

 
Task 1:
1. Write a function/method that accepts a list of string tokens and a separate text string as an input and checks if the latter string can be represented as a concatenation of any subset of tokens from the first argument, where each token can be used multiple times. Examples:
[ "ab", "bc", "a" ] vs. "abc" = true
[ "a", "ab", "bc" ] vs. "ab" = true
[ "a", "ab", "bc" ] vs. "" = true
[ "ab", "bc" ] <-> "abc" = false
[ "ab", "bc", "c" ] <-> "abbcbc" = true
2. Estimate the computational complexity of your code
3. Let's assume the list of tokens can be preprocessed somehow once, and then the result is to be multiple times applied to long input strings. E.g. in Golang notation:

func preprocess(tokens []string) func(input string) bool {
...
}

validator := preprocess(someTokens)
for _, input := range longListOfInputs {
fmt.Printf("%s: %v\n", input, validator(input))
}

We don't care about preprocessor() performance much, but we want validator() to run as far as possible. Could you suggest any implementation that satisfies this extra requirement? if so, what's the computational complexity of that implementation?

Task 2:
1. The input is a list of intervals (two float points); write a function/method that merges any N intervals in the list that have common points (intervals [1.0, 3.5] and [3.5, 6.1] have a common point of 3.5; [4.0, 8.2] and [6.1, 10.3] have common points in the interval [6.1, 8.2]) and returns the merged list (where no two intervals intersect). it is allowed to modify the input. Try to avoid allocating extra memory for the output.
2. What is the computational complexity of your implementation?

The preferred language is Golang (e.g. gives extra points), though any other programming language would be accepted as well (if it's something uncommon make sure the code could be understood by a person not familiar with that language). The solutions could be passed back in an email (inline/attachment(s)) or as a link to some public online repository.



