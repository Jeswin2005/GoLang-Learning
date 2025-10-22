#!/bin/bash

NAME[0]="apple"
NAME[1]="banana"
NAME[2]="mango"
NAME[3]="orange"
NAME[4]="guava"

echo "Array using @ : ${NAME[@]}"
echo "Array using * : ${NAME[*]}"
echo "First Element: ${NAME[0]}"
echo "Second Element: ${NAME[1]}"
