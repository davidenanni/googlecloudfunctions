import random
import datetime    

workload = 50
numGoals = 0

t1 = datetime.datetime.now()

for i in range(0, workload):
    x = random.random()
    y = random.random()

    if (x ** 2 + y ** 2 <= 1.0):
    	numGoals += 1	

t2 = datetime.datetime.now()
t3 = t2 - t1
elapsed_ms = (t3.seconds * 1000) + (t3.microseconds / 1000)

print(elapsed_ms)
print(numGoals/workload*4)