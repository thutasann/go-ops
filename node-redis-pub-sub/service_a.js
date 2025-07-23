// @ts-check
const { initPubSub, publish } = require('./pub_sub');

async function start() {
  await initPubSub();

  setInterval(() => {
    const message = {
      user: 'Alice',
      text: 'Hello from Service A',
      timestamp: Date.now(),
    };
    publish('chat-message', message);
    console.log('Published:', message);
  }, 5000);
}

start();
