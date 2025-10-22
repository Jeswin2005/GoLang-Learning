#!/bin/bash

#Arithmetic Operators
a=10
b=3

echo "Arithmetic Operators"
echo "a = $a, b = $b"
echo "Addition: $((a + b))"
echo "Subtraction: $((a - b))"
echo "Multiplication: $((a * b))"
echo "Division: $((a / b))"
echo "Modulus: $((a % b))"
echo "Increment a: $((++a))"
echo "Decrement b: $((--b))"


#Relational Operator
x=5
y=10

echo ""
echo "Relational Operators"
if [ $x -eq $y ]; then
    echo "$x is equal to $y"
else
    echo "$x is not equal to $y"
fi

if [ $x -lt $y ]; then
    echo "$x is less than $y"
fi

if [ $y -gt $x ]; then
    echo "$y is greater than $x"
fi


#Logical Operators
echo ""
echo "Logical Operators"
if [ $x -lt 10 ] && [ $y -gt 5 ]; then
    echo "Both conditions are true"
fi

if [ $x -lt 10 ] || [ $y -lt 5 ]; then
    echo "At least one condition is true"
fi

if ! [ $x -eq $y ]; then
    echo "x is not equal to y"
fi


#File Operators
file="operators.sh"
echo ""
echo "File Test Operators"
if [ -e "$file" ]; then
    echo "File '$file' exists"
fi

if [ -f "$file" ]; then
    echo "File '$file' is a regular file"
fi

if [ -r "$file" ]; then
    echo "File '$file' is readable"
fi
# -x - executable
# -w - writable
# -d - check if directory
