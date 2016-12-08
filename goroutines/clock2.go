package main

import (
    "io"
    "log"
    "net"
    "time"
    "os"
)

func main()  {
    if len(os.Args[:]) > 1 {
      for _, port := range os.Args[1:] {
        listener, err := net.Listen("tcp", "localhost:" + port)
        if err != nil {
            log.Fatal(err)
        }
        for {
            conn, err := listener.Accept()
            if err != nil {
                log.Print(err)
                continue
            }
            go handleConn(conn)
        }
      }
    }
}

func handleConn(c net.Conn)  {
    defer c.Close()
    for {
        _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
        if err != nil {
            return
        }
        time.Sleep(1 * time.Second)
    }
}
