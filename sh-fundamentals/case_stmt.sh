#! /bin/bash

read -p "Are you 21 or over? Y/N " ANSWER

case "$ANSWER" in
    [yY] | [yY][eE][sS])
        echo "You can hav a beer :)"
        ;;
    [nN] | [nN][oO])
        echo "You are a kid"
        ;;
    *)
        echo "Please enter y/yes or n/no"
esac