exports.handler = async (req,res) => {
    // TODO implement

    var reqJSON = JSON.parse(req.body)
    var workload = parseInt(reqJSON["numb"], 10)
    
    var a = [];

    for (var i = 0; i < workload; i++){
        a.unshift(Math.round(Math.random() * workload))
    }

    var hrstart = process.hrtime()

    var swapped;
    do {
        swapped = false;
        for (var i=0; i < a.length-1; i++) {
            if (a[i] > a[i+1]) {
                var temp = a[i];
                a[i] = a[i+1];
                a[i+1] = temp;
                swapped = true;
            }
        }
    } while (swapped);

  var hrend = process.hrtime(hrstart)
  
  var exeTime = hrend[1] / 1000000    
    
  const response = {
    statusCode: 200,
    body: ''+exeTime,
  };
  
  return response;
};