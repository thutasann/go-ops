#! /bin/bash

# WHILE LOOP - Read through a file line by line
LINE=1
while read -r CURRENT_LINE
    do 
        echo "$LINE: $CURRENT_LINE"
        ((LINE++))
done < "./lorem.txt"