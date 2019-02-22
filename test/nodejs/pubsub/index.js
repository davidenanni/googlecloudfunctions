'use strict';

const {PubSub} = require('@google-cloud/pubsub');
const pubsub = new PubSub();
const Buffer = require('safe-buffer').Buffer;

exports.publish = (req, res) => {
  if (!req.body.topic) {
    res
      .status(500)
      .send(
        new Error(
          'Topic not provided. Make sure you have a "topic" property in your request'
        )
      );
    return;
  } else if (!req.body.message) {
    res
      .status(500)
      .send(
        new Error(
          'Message not provided. Make sure you have a "message" property in your request'
        )
      );
    return;
  }

  console.log(`Publishing message to topic ${req.body.topic}`);

  // References an existing topic
 
  const dataBuffer = Buffer.from(req.body.message);

  pubsub.topic(req.body.topic).publish(dataBuffer);
};