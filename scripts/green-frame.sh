#!/bin/bash

curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "bgrect 0.01 0.01 0.99 0.99"
curl -X POST http://localhost:17000 -d "update"