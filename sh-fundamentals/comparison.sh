#! /bin/bash

#########
# val1 -eq val2
# val1 -ne val2
# val1 -gt val2
# val1 -ge val2
# val1 -lt val2
# val1 -le val2
#########
NUM1=3
NUM2=5

if [ "$NUM1" -gt "$NUM2" ]; then
    echo "$NUM1 is greater than $NUM2"
else
    echo "$NUM1 is less than $NUM2"
fi
