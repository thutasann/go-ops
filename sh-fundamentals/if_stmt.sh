#! /bin/bash

# SIMPLE IF STATEMENT
NAME="Brad"
if [ "$NAME" == "Brad" ]
then
  echo "Your Name is Brad"
fi

# IF-ELSE
if [ "$NAME" == "James" ]
then
  echo "Your Name is James"
else 
  echo "Your Name is not James"
fi

# ELSE-IF
if [ "$NAME" == "James" ]
then
  echo "Your Name is James"
elif [ "$NAME" == "Brad" ]
then 
  echo "Else if Your Name is Brad"
else 
  echo "In the Else"
fi