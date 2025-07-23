// @ts-check
const { createClient } = require('redis');

/** @type { import('redis').RedisClientType} */
let publisher;

/** @type { import('redis').RedisClientType} */
let subscriber;

async function initPubSub() {
  publisher = createClient();
  subscriber = createClient();

  await publisher.connect();
  await subscriber.connect();

  console.log('Redis PubSub initialized');
}

function publish(channel, message) {
  return publisher.publish(channel, JSON.stringify(message));
}

function subscribe(channel, handler) {
  subscriber.subscribe(channel, (msg) => {
    const data = JSON.parse(msg);
    handler(data);
  });
}

const pubSub = {
  initPubSub,
  publish,
  subscribe,
};

module.exports = pubSub;
