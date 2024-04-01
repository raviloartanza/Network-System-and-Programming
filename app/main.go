package main

import (
	"fmt"
	"main/pkg"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:2222")
	pkg.ErrorHandler(err)

	// Menampilkan status bahwa server telah tersambung dengan client
	fmt.Printf("Bound to %q\n", listener.Addr())

	defer listener.Close()

	// Agar server dapat menerima koneksi yang baru terhubung
	// dan mengatur koneksi yang sudah terbuat 
	for {
		conn, err := listener.Accept()
		pkg.ErrorHandler(err)

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Menampilkan status bahwa koneksi
	// antara client dan server telah berhasil
	fmt.Println("[Accepted]")

	// Memberikan batas waktu kepada client dan server
	// dalam hal menulis dan membaca pesan
	deadline := time.Now().Add(15 * time.Second)
	conn.SetReadDeadline(deadline)
	conn.SetWriteDeadline(deadline)

	// Menyimpan data pesan
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	pkg.ErrorHandler(err)

	// Mengirim status "Pesan diterima" ke client
	conn.Write([]byte("[Message Received]"))

	defer conn.Close()
}