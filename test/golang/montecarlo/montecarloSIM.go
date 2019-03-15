package main

import(
		"fmt"
		"math/rand"
		"strconv"
		"math"
		"os"
		_"bufio"
)

func montecarlo(numLanci int64)(float64){
	numLanciArea := 0

	var i int64

    for i=0; i<numLanci; i++{

        x := rand.Float64()
        y := rand.Float64()                

        if ( math.Pow(x,float64(2)) + math.Pow(y,float64(2)) <= 1){
           numLanciArea++
        }
    }

    return float64(numLanciArea)/float64(numLanci)
}

func main(){

	numTest, _ := strconv.ParseInt(os.Args[1], 10, 64)

	pigreco_sim := montecarlo(numTest)*4

	fmt.Println(pigreco_sim)
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