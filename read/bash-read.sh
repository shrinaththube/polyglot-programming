#!/bin/bash

# run command - bash bash-read.sh
# https://www.cyberciti.biz/faq/unix-howto-read-line-by-line-from-file/
# https://www.javatpoint.com/bash-read-file

filename="sample-read.txt"


echo "****************************************"
echo "1) Read line by line"
echo "****************************************"

while IFS= read -r line
do
  echo "$line"
done < "$filename"
echo "$line"

echo "****************************************"
echo "2) Read everything"
echo "****************************************"

value=$(<$filename)  
echo "$value"  