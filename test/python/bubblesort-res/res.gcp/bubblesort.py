from random import *
import datetime
import sys

def bubblesort_test(workload):

	randomValues = [] 

	for i in range(0, workload):
		randomValues.append(randint(1, workload))

	workload = workload-1

	t1 = datetime.datetime.now()
	while True:
		swapped = False
		for i in range(0, workload):

			prevValue = randomValues[i]
			succValue = randomValues[i+1]

			if prevValue > succValue:
				temp = randomValues[i+1]
				randomValues[i+1] = randomValues[i]
				randomValues[i] = temp
				swapped = True

		if swapped == False:
			break
	t2 = datetime.datetime.now()
	t3 = t2 - t1
	elapsed_ms = (t3.seconds * 1000) + (t3.microseconds / 1000)
	return elapsed_ms



workload = int(sys.argv[1])
numTest = int(sys.argv[2])
testFile = sys.argv[3]

file = open(testFile,'w') 


for i in range(0, numTest):
	elapsed_ms = bubblesort_test(workload)
	elapsed_ms_str = str(elapsed_ms)
	file.write(elapsed_ms_str)
	file.write('\n')

file.close()