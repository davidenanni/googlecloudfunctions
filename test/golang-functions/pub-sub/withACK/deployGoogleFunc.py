import os
import csv 
from macro import readCSV

funcParams = "./deployGoogleFunc.csv"


def createGoMod( package ):
	os.system("go mod init "+package)
	os.system("go build")


def deployGoogleFunction():
	values = readCSV( funcParams )

	funcName 	= values[0]
	funcSource  = values[1]
	runtime		= values[2]
	memory		= values[3]
	region		= values[4]

	trigger = values[6]

	if trigger.find("--trigger-topic")!=-1:
		trigger = trigger+" "+values[7]

	depIns = ("gcloud functions deploy "+funcName+" --entry-point "+funcSource+
			  " --runtime "+runtime+" --memory "+memory+" --region "+region+" "+
			  trigger)

	os.system( depIns )


createGoMod("helloworld")
deployGoogleFunction()