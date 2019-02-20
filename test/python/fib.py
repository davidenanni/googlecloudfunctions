import sys

def func(numb):
	if numb == 1 or numb == 2:
		return 1
	else:
		return func(numb-1)+func(numb-2)

print(func(int(sys.argv[1])))		