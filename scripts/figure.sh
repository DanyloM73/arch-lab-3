#!/bin/bash

curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "bgrect 0.1 0.1 0.9 0.9"
curl -X POST http://localhost:17000 -d "figure 0.5 0.5"
curl -X POST http://localhost:17000 -d "figure 0.3 0.3"
curl -X POST http://localhost:17000 -d "figure 0.7 0.7"
curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "update"