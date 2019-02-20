import sys

def func(numb):
	if numb == 1:
		return 1
	else:
		return numb*func(numb-1)


print(func(int(sys.argv[1])))