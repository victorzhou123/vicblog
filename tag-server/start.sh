#!/bin/bash

# define
server() {
    ./category-server
}

proto() {
    make proto
}

# run
proto
server