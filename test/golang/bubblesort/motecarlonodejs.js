exports.montecarlo = async (req, res) => {
	var reqJSON = JSON.parse(req.body) 
	var workload = parseInt(reqJSON["numb"], 10)

	var numGoals = 0
	var x = 0.0
	var y = 0.0

	var hrstart = process.hrtime()

	for (var i = 0; i < workload; i++){
		x = Math.random()
		y = Math.random()

		if ( Math.pow(x,2) + Math.pow(y,2) <= 1.0)
			numGoals++
	}

	var pigreco = numGoals/workload

	var hrend = process.hrtime(hrstart)
	var exeTime = hrend[1] / 1000000 

	res.send(pigreco.toString());
};