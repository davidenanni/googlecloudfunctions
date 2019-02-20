import random
import datetime

def montecarlo(request):
	
	request_json = request.get_json(silent=True)
	workload = int(request_json['workload'])
	t1 = datetime.datetime.now()

	numGoals = 0

	for i in range(0, workload):
		x = random.random()
		y = random.random()

		if (x ** 2 + y ** 2 <= 1.0):
			numGoals += 1	

	t2 = datetime.datetime.now()
	t3 = t2 - t1
	elapsed_ms = (t3.seconds * 1000) + (t3.microseconds / 1000)
	return '{}'.format(elapsed_ms)