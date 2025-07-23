// @ts-check

const { initPubSub, subscribe } = require('./pub_sub');
const { Server } = require('socket.io');
const http = require('http');

const server = http.createServer();
const io = new Server(server, {
  cors: { origin: '*' },
});

server.listen(3000, () => {
  console.log('Service B - Socket.IO server running on port 3000');
});

async function start() {
  await initPubSub();

  subscribe('chat-message', (message) => {
    console.log('Service B received:', message);
    io.emit('new-message', message); // Emit to clients
  });
}

start();
