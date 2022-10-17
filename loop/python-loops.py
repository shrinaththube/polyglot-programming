# run command - python3 python-loops.py

cool_list = ["i" , "am", "reading", "all" , "elements", "of", "this","list"]

# while loop
print("*"*40)
print("1) Let's start with 'while' loop")
print("*"*40)

ind = 0
while(ind < len(cool_list)):
    # Don't forget to increment index by 1
    print(cool_list[ind])    
    ind +=1
# ------------------------------------------------------------------------------------------- #

# for loop access element by element
print("*"*40)
print("2) Move on to 'for' loop")
print("*"*40)

for ele in cool_list:
    print(ele)

# ------------------------------------------------------------------------------------------- #

# for loop access elements using index
print("*"*40)
print("3) Another variation 'for' loop using indexing ")
print("*"*40)

for i in range(len(cool_list)):
    print(cool_list[i])
