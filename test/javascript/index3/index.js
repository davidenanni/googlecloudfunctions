exports.helloWorld3 = function helloWorld3 (req, res) {

  //if (req.body.message === undefined) {

  //  res.status(400).send('err error');

  //} else {

    console.log(req.body.message);

    res.status(200).send('Hello World GCP v3!');

  //}

};
