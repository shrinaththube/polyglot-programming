#!/bin/bash

# run command - bash bash-loops.sh
# ref - https://www.geeksforgeeks.org/bash-scripting-for-loop/
# https://linuxhint.com/bash-for-loop-examples/
# https://www.cyberciti.biz/faq/bash-while-loop/

cool_list=("i" "am" "reading" "all" "elements" "of" "this" "list")

# while loop
echo "****************************************"
echo "1) Let's start with 'while' loop"
echo "****************************************"

len="${#cool_list[@]}"
i=0
while [ $i -le $(($len-1)) ]
do
    echo "$i ${cool_list[$i]}"
    i=$(( $i + 1 ))
done

# ------------------------------------------------------------------------------------------- #

# for loop
echo "****************************************"
echo "2) Move on to 'for' loop"
echo "****************************************"

len="${#cool_list[@]}"

for (( i=0; i<$len; i++ ));
do
    echo "${cool_list[$i]}"
done

# ------------------------------------------------------------------------------------------- #

# for loop
echo "****************************************"
echo "3) Another variation of 'for' loop"
echo "****************************************"

for ele in ${cool_list[@]};
do
    echo $ele
done