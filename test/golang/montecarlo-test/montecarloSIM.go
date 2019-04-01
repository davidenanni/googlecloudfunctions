package main

import(
		"fmt"
		"math/rand"
		"strconv"
		"math"
		"os"
		"time"
		_"bufio"
)

func montecarlo(numLanci int64)(float64,float64){
	numLanciArea := 0

	var i int64

    rand.Seed(time.Now().UnixNano())
    timestamp := time.Now()

    for i=0; i<numLanci; i++{

        x := rand.Float64()
        y := rand.Float64()                

        if ( math.Pow(x,float64(2)) + math.Pow(y,float64(2)) <= 1){
           numLanciArea++
        }
    }

    timestampCompleted := time.Now()

    executionTime := timestampCompleted.Sub(timestamp)

    executionTime_ms := float64(executionTime/time.Microsecond)/1000

    estimated_pi := float64(numLanciArea)/float64(numLanci)*4

    return estimated_pi,executionTime_ms
}

func main(){

	numTest, _ := strconv.ParseInt(os.Args[1], 10, 64)

	workload, _ := strconv.ParseInt(os.Args[2], 10, 64)


    file, err := os.Create("res-"+os.Args[2]+".csv")
    if err != nil {
        return
    }
    defer file.Close()

	var i int64

	for i=0; i<numTest; i++{

		pigreco_sim,timeExec := montecarlo(workload)	

		file.WriteString(fmt.Sprintf("%f",pigreco_sim)+","+fmt.Sprintf("%f",timeExec)+"\n")	
	}
}
	//
	//maxTest, _ := strconv.ParseInt(os.Args[2], 10, 64)	

	//fileTest, _ := os.Create("./sim.csv")

	//defer fileTest.Close()
	//writeFileTest := bufio.NewWriter(fileTest)

	/*var i int64

	for i=minTest; i<maxTest; i++{

		
		row := fmt.Sprintf("%d,%f\n",i,pigreco_sim)	
		writeFileTest.WriteString(row)
	}

	writeFileTest.Flush()
	//
	//fmt.Println(pigreco_sim*4)
}*/