const request = require('request').defaults({
    encoding: null
  });


const {Storage} = require('@google-cloud/storage');


	// Instantiate a storage client
const storage = new Storage();


let url = "https://api.telegram.org/bot779199701:AAEuqK6B4YxM-uKPZzSU5x6QYVIPZP9xbRA/sendMessage?chat_id=$chatID&text=$message";
/**
 * Responds to any HTTP request that can provide a "message" field in the body.
 *
 * @param {!Object} req Cloud Function request context.
 * @param {!Object} res Cloud Function response context.
 */
exports.bmi = (req, res) => {

const reqTelegram = require('request').defaults({
    encoding: null
  });

   
  if (req.body.message === undefined) {
    // This is an error case, as "message" is required.
    res.status(400).send('No message defined!');
  } else {
    // Everything is okay.
    console.log(req.body.message);

    message = req.body.message

    chat = message.chat
    console.log("[CHAT]")
    console.log(chat)
    console.log(chat.id)


    photo = message.photo
    console.log("[PHOTO]")
    console.log(photo[0].file_id)

	urlFilePath = "https://api.telegram.org/bot779199701:AAEuqK6B4YxM-uKPZzSU5x6QYVIPZP9xbRA/getFile?file_id="+photo[0].file_id	
  	console.log(urlFilePath)


  reqTelegram.get(urlFilePath, (error, response, body) => {
      let json = JSON.parse(body);
      console.log(json.result.file_path);

      downloadFile = "https://api.telegram.org/file/bot779199701:AAEuqK6B4YxM-uKPZzSU5x6QYVIPZP9xbRA/"+json.result.file_path

      console.log(downloadFile)
      //const reqTelegram = require('request').defaults({
      //	encoding: null
  	  //	});

      //reqTelegram.get(downloadFile, (error, response, body) => {
      //		let json = JSON.parse(body);
      //		console.log(json);

      //		const bucket = storage.bucket("davnan-telegrambot");
      //		bucket.upload(""+body);
      //});
    });


  sendMessage(chat.id,"Your position is shared!")


	//


    //
    //

    return res.status(200).end();            
  }
};

function sendMessage(chatID, message) {
  let newUrl = url.replace("$message", "\""+message+"\"")
  newUrl = newUrl.replace("$chatID", chatID)
  console.log(newUrl)
  request.get(newUrl, (error, response, body) => {
      let json = JSON.parse(body);
      console.log(json);
    });
}