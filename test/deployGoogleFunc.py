import os
import sys
import csv 
from macro import readCSV


def createGoMod( package ):
	os.system("go mod init "+package)
	os.system("go build")


def setupTriggerEvent( type, params ):
	trigger = ""


	if ( type.find("[TRIGGER]-HTTP") != -1 ):
		trigger = "--trigger-http "
	elif ( type.find("[TRIGGER]-PUBSUB") != -1 ):
		trigger = "--trigger-topic "+params[0]
	elif ( type.find("[TRIGGER]-STORAGE") != -1 ):	
		trigger = ("--trigger-resource "+params[0]+
				   " --trigger-event "+params[1])


	return trigger


def deployGoogleFunction( runtime, func):

	runtimePath = ("./"+ runtime +"/"+ func)

	funcParamsFile = (runtimePath + "/deployGoogleFunc.csv")
	values = readCSV( funcParamsFile )

	funcName 	= values[0]
	funcSource  = values[1]
	runtime		= values[2]
	memory		= values[3]
	region		= values[4]
	timeout		= values[5]

	trigger     = setupTriggerEvent(values[6], values[7:])


	depIns = ("cd "+ runtimePath+" && gcloud functions deploy "+funcName+" --entry-point "+funcSource+
			  " --runtime "+runtime+" --memory "+memory+" --region "+region+" --timeout "+timeout+" "+
			  trigger)

	os.system( depIns )



runtime = sys.argv[1]
func = sys.argv[2]
#provider = sys.argv[3]

if (runtime.find("golang") != -1):
	gomodPackage = sys.argv[3]

	createGoMod(gomodPackage)

deployGoogleFunction(runtime,func)