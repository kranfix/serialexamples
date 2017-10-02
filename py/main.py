import serial
import threading
from time import sleep

isReading = True

def sRead():
    while isReading:
        answer = s.read(1)
        if answer:
            print "Rx:", answer
    print "Serialport: Finishing reading"

try:
    s = serial.Serial('/dev/ttyUSB0',9600,timeout=0.050)
    print "Serialport: Opened"
    tRx = threading.Thread(target=sRead)
    tRx.start()

    led = True
    while True:
        if led:
            m = b'A'
        else:
            m = b'B'
        print "Tx:", m
        s.write(m)
        led = not led
        sleep(2)

except KeyboardInterrupt:
    print "" # Good format
    isReading = False
    #s.write(b'Z')
    #print "Tx: Z"
    tRx.join()
    s.close()
    print "Serialport: Closed"

93AQ02GYY19
sudo apt remove node-red
npm i -g node-red
npm i node-red-node-serialport
node-red
