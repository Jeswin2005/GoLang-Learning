#!/bin/bash

# For loop
echo "For loop"

# Loop through explicit list
for i in 1 2 3 4 5
do
    echo "$i"
done

# Loop using brace expansion
for i in {1..5}   
do
    echo "$i"
done

# C-style loop
for ((i=0; i<5; i++))
do
    echo "$i"
done

# While loop
echo ""
echo "While loop"
a=0          
while [ $a -lt 10 ] 
do
    echo "$a"
    ((a++))
done

# Case statement
echo ""
echo "Enter a number (1-7): "
read day

case $day in
    1)
        echo "Monday"
        ;;
    2)
        echo "Tuesday"
        ;;
    3)
        echo "Wednesday"
        ;;
    4)
        echo "Thursday"
        ;;
    5)
        echo "Friday"
        ;;
    6)
        echo "Saturday"
        ;;
    7)
        echo "Sunday"
        ;;
    *)
        echo "Invalid input"
        ;;
esac

