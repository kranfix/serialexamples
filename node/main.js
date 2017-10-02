var SerialPort = require('serialport');
var port = new SerialPort('/dev/ttyUSB0', {
  baudRate: 9600
});


let led = true;
let m = 'A';
setInterval(() => {
  if(led){
    m = 'A';
  } else {
    m = 'B';
  }

  port.write(m, function(err) {
    if (err) {
      return console.log('Error on write: ', err.message);
    }
    console.log('Tx:',m);
    led = !led;
  });
},2000);


port.on('open', function() {
  console.log('Serialport: Opened');
});

port.on('readable', function () {
  console.log('Rx:', port.read().toString());
});

port.on('close', function() {
  console.log('Serialport: Closed');
});
