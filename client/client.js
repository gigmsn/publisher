const WebSocket = require('ws');
const Chance = require('chance');

const chance = new Chance();
const ws = new WebSocket('ws://server:3000/ws', {
  origin: 'http://server:3000/'
});

var counter = 0;

ws.on('open', function open() {
  console.log('connected');
  ws.send(chance.sentence());
});

ws.on('close', function close() {
  console.log('disconnected');
});

ws.on('message', function incoming(data) {
  console.log(data)
  setTimeout(function timeout() {
    console.log('sending message #'+counter)
    ws.send(chance.sentence());
    counter ++;
  }, 100);
});
