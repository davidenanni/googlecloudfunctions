const escapeHtml = require('escape-html');

/**
 * HTTP Cloud Function.
 *
 * @param {Object} req Cloud Function request context.
 *                     More info: https://expressjs.com/en/api.html#req
 * @param {Object} res Cloud Function response context.
 *                     More info: https://expressjs.com/en/api.html#res
 */
exports.bubblesort = (req, res) => {
  workload = parseInt(req.body.numb, 10);

  var a = [];

  for (var i = 0; i < workload; i++){
    a.unshift(Math.round(Math.random() * workload))
  }

  var hrstart = process.hrtime()//new Date().getTime();

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

  var hrend = process.hrtime(hrstart)//new Date().getTime();
  
  var exeTime = hrend[1] / 1000000//timestampAfter - timestamp

  res.writeHead(200, {"Content-Type": "text/plain"});
  res.write(''+exeTime);
  res.end();
};