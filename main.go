package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/girishmallya123/concurrent_ftpserver/ftp_server"
)

var port int
var rootDirectory string

/* below is an init function that handles a few command line arguments
command line arguments:
	1) port (Integer) : Port number to run the ftp server.
	2) rootDir (String): String parameter that identifies the root directory for the ftp server
*/
func init() {
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&rootDirectory, "rootDir", "public", "root directory")
	flag.Parse()
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDirectory)
	if err != nil {
		log.Fatal(err)
	}
	ftp_server.Serve(ftp_server.NewConn(c, absPath))
}

func main() {
	//fmt.Println("Hello, Welcome to this ftp-server.")
	//fmt.Println("Launching the ftp server on port number ", port)

	server := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", server)
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
