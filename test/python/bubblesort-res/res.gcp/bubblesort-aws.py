from random import *
import datetime
import json

def bubblesort(event, context):
    
    reqBody = event['body']
    reqBodyJSON = json.loads(reqBody)
   
    workload = int(reqBodyJSON['numb'])
    
    randomValues = []
    
    for i in range(0,workload):
    	randomValues.append(randint(1,workload))
    
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
	ElapsedmsTime = (t3.seconds * 1000) + (t3.microseconds / 1000)    
    
    result = str(ElapsedmsTime)
    
    # TODO implement
    return {
        'statusCode': 200,
        'body': json.dumps(""+result+"")
    }