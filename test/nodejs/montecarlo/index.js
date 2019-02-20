const escapeHtml = require('escape-html');

/**
 * HTTP Cloud Function.
 *
 * @param {Object} req Cloud Function request context.
 *                     More info: https://expressjs.com/en/api.html#req
 * @param {Object} res Cloud Function response context.
 *                     More info: https://expressjs.com/en/api.html#res
 */
exports.montecarlo = (req, res) => {
  workload = parseInt(req.body.workload, 10);

  var numGoals = 0

  var timestamp = new Date().getTime();

  for (var i = 0; i < workload; i++){
  	x = Math.random()
  	y = Math.random()

  	if ( Math.pow(x,2) + Math.pow(y,2) <= 1.0)
  		numGoals++
  }

  pigreco = numGoals/workload

  var timestampAfter = new Date().getTime();
  
  var exeTime = timestampAfter - timestamp

  res.send(`${escapeHtml(''+exeTime)}`);
};