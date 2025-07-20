#! /bin/bash

# FUNCTION
function sayHello() {
    echo "Hello World"
}

sayHello

# FUNCTION WITH PARAMS
function greet() {
    echo "Hello, I am $1 and I am $2 years old"
}

greet "Brad" "36"