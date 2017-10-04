package main

import (
  "github.com/tarm/serial"
  "fmt"
  "time"
)

func main() {
  c := &serial.Config{Name: "/dev/ttyUSB0",
     Baud: 9600,
     ReadTimeout: time.Millisecond * 5}
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


  led := true
  m := []byte("A")
  buf := make([]byte, 4)

  tick := time.Tick(250 * time.Millisecond)
  boom := time.After(1500 * time.Millisecond)
  fmt.Println("Debug")
  for {
    select {
    case <-tick:
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

      led = !led
    case <-boom:
      return
    default:
      if (isReading){
        n, err := s.Read(buf)
        if err == nil && n != 0 {
          //fmt.Printf("Rx: %q\n", buf[:n])
          fmt.Printf("Rx: %s\n", buf[:n])
        } /*else {
          fmt.Println(err)
        }*/
      } else {
        fmt.Println("Serialport: Finishing reading")
      }
    }
  }
}
