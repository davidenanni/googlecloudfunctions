from random import random
import time
import json

def montecarlo(event):
	
	reqBody = event['body']
	reqBodyJSON = json.loads(reqBody)
   
	workload = int(reqBodyJSON['numb'])
	
	numGoals = 0

	t1 = time.process_time_ns()
	
	for i in range(0, workload):
		x = random()
		y = random()
		if (x ** 2 + y ** 2 <= 1.0):
			numGoals += 1
	
	pigreco = float(numGoals)/float(workload) 
	
	elapsed_time = time.process_time_ns() - t1
	elapsed_time_ms = elapsed_time / 1000000.0
	
	# TODO implement
	return {
		'statusCode': 200,
		'body': json.dumps(elapsed_time_ms)
	}