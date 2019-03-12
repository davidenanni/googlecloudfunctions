import csv
from os import listdir
from os.path import isfile, join

from libcloud.compute.types import Provider
from libcloud.compute.providers import get_driver


crdPath = "./conf/crds/"

def getGCECredentials():
    serviceAccountEmail, privateKeyPath, projectID = getCredentials()

    return get_driver(Provider.GCE)(serviceAccountEmail, privateKeyPath, project=projectID)


def getCredentials():
    crdsFile = crdPath+"crds.csv"
    values = readCSV(crdsFile)

    projectID           = values[0]
    serviceAccountEmail = values[1]
    privateKeyPath      = crdPath+""+values[2]

    return serviceAccountEmail, privateKeyPath, projectID


def readFile( filename ):

    rowsFile    = open( filename )

    return rowsFile.read()


def readCSV(filePath):
    values = []

    with open(filePath) as csvfile:
	    csvreader = csv.reader(csvfile, delimiter=',')
	    for row in csvreader:
	    	values.append(row[1])

    return values


def getDirsList(directoryPath):
	return [d for d in listdir(directoryPath) if not isfile(join(directoryPath, d))]


def getFilesList(directoryPath):
	return [f for f in listdir(directoryPath) if isfile(join(directoryPath, f))]
