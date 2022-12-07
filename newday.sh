#!/bin/bash

mkdir $1;
cd $1;
touch README.md
touch input_test.txt
touch input.txt
cp ../template.go ./solution.go
