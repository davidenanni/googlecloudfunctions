const request = require('request').defaults({
    encoding: null
  });

let url = "https://api.telegram.org/token/sendMessage?chat_id=chat_id&text=$message";
/**
 * Responds to any HTTP request that can provide a "message" field in the body.
 *
 * @param {!Object} req Cloud Function request context.
 * @param {!Object} res Cloud Function response context.
 */
exports.bmi = (req, res) => {
   
  if (req.body.message === undefined) {
    // This is an error case, as "message" is required.
    res.status(400).send('No message defined!');
  } else {
    // Everything is okay.
    console.log(req.body.message);
    let message = req.body.message.text;
        
    let re = /\d\d\/\d\d\d/;
    if(!message.match(re)) {
      sendMessage("Put Weight in Kilograms/Height in centimeters to calculate your BMI. Ex 87/188");
      return res.status(200).end();
    }
   
    let weight = message.split("/")[0];
    let height = message.split("/")[1]/100;
    
    let num_result =  weight/Math.pow(height,2);
    console.log(num_result)
    if(num_result >= 25) {
      sendMessage("You are in overweight range."); 
    } else if(num_result > 18.5 && num_result < 25) {
        sendMessage("You are in healthy range");
    } else {
        sendMessage("You are in underweight range.");
    }
    
    return res.status(200).end();            
  }
};

function sendMessage(message) {
  let newUrl = url.replace("$message", message)
  request.get(newUrl, (error, response, body) => {
      let json = JSON.parse(body);
      console.log(json);
    });
}