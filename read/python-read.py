# run command - python3 python-read.py

# https://www.freecodecamp.org/news/how-to-read-files-in-python/
# https://www.pythontutorial.net/python-basics/python-read-text-file/

filename = "sample-read.txt"

print("*"*40)
print("1) Read line by line")
print("*"*40)
with open(filename) as f:
    for line in f.readlines():
        print(line)

print("*"*40)
print("2) Read entire file")
print("*"*40)

file_operator = open(filename,'r')
print(file_operator.read())
file_operator.close()