#!/bin/bash

max=0.99  
min=0.01  

X=0.5
Y=0.5

curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "figure $X $Y"
curl -X POST http://localhost:17000 -d "update"

step=0.01

while true; do
    if (( $(echo "$X >= $max" | bc -l) )); then
        step=-0.01
    elif (( $(echo "$X <= $min" | bc -l) )); then
        step=0.01
    fi
    
    curl -X POST http://localhost:17000 -d "move $step $step"
    curl -X POST http://localhost:17000 -d "update"
    X=$(echo "$X+$step" | bc -l)
    Y=$(echo "$Y+$step" | bc -l)
done