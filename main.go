package main

import (
    "net"
    "flag"
    "fmt"
)



func main() {
    //define the command line arguments
    serverPtr := flag.String("url", "localhost", "the url of the server to connect")
    portPtr := flag.Int("port", 6579, "the port to connect")
    flag.Parse()

    fmt.Println("stating the connection...")
    port := fmt.Sprintf(":%d", *portPtr)
    conn, err := net.Dial("tcp", *serverPtr + port)
    if err != nil {
        fmt.Printf("can't connect to the specified server... becoming one! Host: " + *serverPtr + " Port: " + port  )
        ln, err := net.Listen("tcp", port)
        if err != nil {
            fmt.Printf("error creating server: " + err.Error())
            return
        }
        for {
            conn, err := ln.Accept()
            if err != nil {
                fmt.Printf("error receiving message: " + err.Error())
                continue
            }
            dat, err := ioutil.ReadFile("./filetogo.txt")

            dataComing := make([]byte, 1000)
            _, err = conn.Read(dataComing)
            if err != nil {
                fmt.Printf("error reading  message: " + err.Error())
                continue
            }
            fmt.Println("chegou:")
            fmt.Println(string(dataComing))
        }
    }
    fmt.Println("conectado! pronto pra mandar: ")
    var dado string
    fmt.Scanln(&dado)
    conn.Write([]byte(dado))

}
