import os
import sys
import csv 
from macro import readCSV

funcParams = "./deployGoogleFunc-"


def createGoMod( package ):
	os.system("go mod init "+package)
	os.system("go build")


def setupTriggerEvent( type, params ):
	trigger = ""

	if ( type.find("[TRIGGER]-PUBSUB") != -1 ):
		trigger = "--trigger-topic "+params[0]

	if ( type.find("[TRIGGER]-STORAGE") != -1 ):	
		trigger = ("--trigger-resource "+params[0]+
				   " --trigger-event "+params[1])

	return trigger


def deployGoogleFunction( gofunc ):
	funcParamsFile = (funcParams + "" + gofunc + ".csv")
	values = readCSV( funcParamsFile )

	funcName 	= values[0]
	funcSource  = values[1]
	runtime		= values[2]
	memory		= values[3]
	region		= values[4]

	trigger     = setupTriggerEvent(values[5], values[6:])


	depIns = ("gcloud functions deploy "+funcName+" --entry-point "+funcSource+
			  " --runtime "+runtime+" --memory "+memory+" --region "+region+" "+
			  trigger)

	os.system( depIns )


gomod = sys.argv[1]
gofunc = sys.argv[2]

createGoMod(gomod)
deployGoogleFunction(gofunc)