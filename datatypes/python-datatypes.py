
# run command - python3 python-datatypes.py
# https://www.w3schools.com/python/python_datatypes.asp

# Text Type:	str
print("*"*40)
print("1) Text type")
print("*"*40)

s = "no strings attached"
print(type(s) , " : ", s)

# Numeric Types:	int, float, complex
print("*"*40)
print("2) Numeric type")
print("*"*40)

n = 101
print(type(n) , " : ", n)

pi = 3.141592653589793238
print(type(pi) , " : ", pi)

comp = complex(3,5)
print(type(comp) , " : ", comp)

# Sequence Types:	list, tuple, range
print("*"*40)
print("3) Sequence type")
print("*"*40)

l = ["lets","get","back","to","the","list"]
print(type(l) , " : ", l)

t = ("lets","get","back","to","the","tuple")
print(type(t) , " : ", t)

# Mapping Type:	dict
print("*"*40)
print("4) Mapping type")
print("*"*40)

d = {"key": "point", "no" : "way", "get": "on"}
print(type(d) , " : ", d)

# Set Types:	set, frozenset
print("*"*40)
print("5) Set type")
print("*"*40)

s = {"get", "set", "go"}
print(type(s) , " : ", s)

froz = frozenset({"frozen", "glacier", "should", "not", "melt"})
print(type(froz) , " : ", froz)

# Boolean Type:	bool
print("*"*40)
print("6) Boolean type")
print("*"*40)

b = True
print(type(b) , " : ", b)

# Binary Types:	bytes, bytearray, memoryview
print("*"*40)
print("7) Binary type")
print("*"*40)

b = b"bytes"
print(type(b) , " : ", b)

barr = bytearray(5)
print(type(barr) , " : ", barr)

m = memoryview(b)
print(type(m) , " : ", m)

# None Type:	NoneType
print("*"*40)
print("8) None type")
print("*"*40)

noone = None
print(type(noone) , " : ", noone)


