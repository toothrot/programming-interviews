#!/bin/bash

for i in $(find -name 'problem_*.sql'); do
  echo "Solution for $i"
  psql postgres < $i;
done
