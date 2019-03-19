exports.ping = function ping (req, res) {
    console.log(req.body.message);

    res.status(200).send('');
};