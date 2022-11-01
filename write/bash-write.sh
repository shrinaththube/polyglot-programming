#!/bin/bash

# run command - bash bash-loops.sh
# ref -
# https://www.countryliving.com/life/g4559/halloween-quotes/
# https://linuxize.com/post/bash-write-to-file/
# https://superuser.com/questions/154936/echo-text-with-new-line-in-bash

filename="sample-write.txt"

echo "1) Write a small string in a file"

# -e allows to adde newline char at the end
echo -e '"Trick or treat, bag of sweets, ghosts are walking down the street." -Unknown\n' >> $filename

# ------------------------------------------------------------------------------------------- #

echo "2) Write multi-line information in a file"

cat << EOF >> $filename
"A mask tells us more than a face." -Oscar Wilde

"Halloween was confusing. All my life my parents said, 'Never take candy from strangers.' And then they dressed me up and said, 'Go beg for it.' " -Rita Rudner

"We make up horrors to help us cope with the real ones." -Stephen King

EOF

# ------------------------------------------------------------------------------------------- #

echo "3) Write a list of lines line by line into a file"

spooky=('"Do you believe in destiny? That even the powers of time can be altered for a single purpose?" -Dracula (1992)' 
'"On Halloween, witches come true; wild ghosts escape from dreams. Each monster dances in the park." -Nick Gordon' 
'"There is something haunting in the light of the moon." -Joseph Conrad' '"By the pricking of my thumbs, something 
wicked this way comes." -William Shakespeare')

# if you don't hadd double quotes around spooky list then for will access each word separately
for quote in "${spooky[@]}"
do
    (echo $quote ; echo "") >> $filename
    # echo "" >> $filename
done
