package main

import (
	"fmt"
	"main/pkg"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:2222")
	pkg.ErrorHandler(err)

	defer dial.Close()

	// Memberikan waktu selama 5 detik untuk client
	// terhubung dengan server, jika lebih akan timeout
	dial.SetReadDeadline(time.Now().Add(5 * time.Second))

	// Mengirim pesan "Request" ke server
	dial.Write([]byte("[Request]"))

	// Menyimpan data pesan
	buffer := make([]byte, 1024)
	_, err = dial.Read(buffer)
	pkg.ErrorHandler(err)

	// Menampilkan status jawaban dari server
	fmt.Println("Server:", string(buffer))
}