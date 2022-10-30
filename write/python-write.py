
# run command - python3 python-write.py

# https://www.w3schools.com/python/python_file_write.asp
# https://www.pythontutorial.net/python-basics/python-write-text-file/

filename = "sample-write.txt"

print("1) Write in one shot")

fileoperator = open(filename, "a")
fileoperator.write('"A man is but a product of his thoughts. What he thinks he becomes." -MK Gandhi \n')
fileoperator.write('\n')
fileoperator.close()

print("2) Write from list of lines")

quotes = ['"The sign of intelligence is that you are constantly wondering. Idiots are always dead sure \
about every damn thing they are doing in their life." -Sadhguru', \
'"When you are inspired by some great purpose, some extraordinary project, all your thoughts \
break their bonds; your mind transcends limitations; your conscious expands in every direction; \
and you find yourself in a great, new and wonderful world." -Patanjali']

with open(filename, "a") as fo:
    for quote in quotes:
        fo.write(quote)
        fo.write('\n\n')

print("3) Write in one shot")

more_quotes = ['"Dream, dream, dream. Dreams transform into thoughts and thoughts result in action." -APJ Abdul Kalam', \
               '"If you resist change, you resist life." -Sadhguru']
with open(filename, 'a') as f:
    f.write('\n\n'.join(more_quotes))
    f.write('\n\n')