package main

import (
  "github.com/tarm/serial"
  "fmt"
  "time"
)

func main() {
  c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
  s, err := serial.OpenPort(c)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("Serialport: Opened")
  isReading := true;
  defer func() {
    isReading = false;
    s.Close()
    fmt.Println("Serialport: Closed")
  }()

  // Reading process
  go func() {
    buf := make([]byte, 4)
    for {
      n, err := s.Read(buf)
      if err == nil {
        //fmt.Printf("Rx: %q\n", buf[:n])
        fmt.Printf("Rx: %s\n", buf[:n])
      } else {
        fmt.Println(err)
      }
    }
    fmt.Println("Serialport: Finishing reading")
  }()


  var led = true;
  m := []byte("A")
  for {
    if led {
      m = []byte("A")
    } else {
      m = []byte("B")
    }
    n, err := s.Write(m)
    if err == nil {
      fmt.Printf("Tx: %s\n",m[:n])
    } else {
      fmt.Println(err)
    }

    led = !led;
    time.Sleep(2 * time.Second)
  }
}
