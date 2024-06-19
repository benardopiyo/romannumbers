# RomanNumbers

* `RomanNumbers` is a program designed to convert a given number (within a specific range) into its Roman numeral representation and provide a step-by-step calculation of the conversion. 

* The program has a limit of 4000. If the input number is invalid, such as a non-numeric string or a number outside the valid range (1-3999), the program will print an error message.


## Functionality

* The program includes the following features:

    - Convert a number (given as an argument) into a Roman numeral.
    - Print the Roman numeral and its calculation.
    - Handle invalid inputs and numbers outside the range (1-3999).

* Roman Numerals

    I - 1
    V - 5
    X - 10
    L - 50
    C - 100
    D - 500
    M - 1000

* Subtractive Notation:

    IV - 4
    IX - 9
    XL - 40
    XC - 90
    CD - 400
    CM - 900

* Example

Given the number 1732, the program will output:


M+D+CC+X+X+I+I
MDCCXXXII

## Usage

To run the program, use the following command:

```bash
$ go run . <number>
```

### Examples

```bash
$ go run . hello
ERROR: cannot convert to roman digit
```

```bash
$ go run . 123
C+X+X+I+I+I
CXXIII
```

```bash
$ go run . 999
(M-C)+(C-X)+(X-I)
CMXCIX
```

```bash
$ go run . 3999
M+M+M+(M-C)+(C-X)+(X-I)
MMMCMXCIX
```

```bash
$ go run . 4000
ERROR: cannot convert to roman digit
```
