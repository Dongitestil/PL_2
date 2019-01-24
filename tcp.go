package main

import "net"
import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {

	var f, flag int
	f = 0
	flag = 0
	for (flag!=1) {
		fmt.Print("Выберите режим приложения(1 - клиент/2 - сервер) ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		f, _ := strconv.Atoi(text[0:1])
		if (f!=1)&&(f!=2) {
			fmt.Print("!Некорректный выбор!\n")
		} else { 
			 flag = 1
			}
	}
	
	if (f == 1) {
		fmt.Print("!Некорректный выбор!\n")
		// connect to this socket
		conn, _ := net.Dial("tcp", "127.0.0.1:8081")
		for { 
		var hash_string,skey_initial string
		protectorcl := new(Protector) 
		hash_string = protectorcl.get_hash_str()
		skey_initial = protectorcl.get_session_key()
		protectorcl.set_key(skey_initial)
		protectorcl.set_hash(hash_string)
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ваше сообщение: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, hash_string + "\n")
		fmt.Fprintf(conn, skey_initial + "\n")
		// listen for reply
		key_mess, _ := bufio.NewReader(conn).ReadString('\n')
		serv_key := strings.Trim(key_mess,"\n")
		new_key:=protectorcl.next_session_key()
		fmt.Println("Текущий ключ: " + new_key) 
		if new_key == serv_key {
			fmt.Println("Ключи совпали ",text)
		}
		}
	}

	if f == 2 {
		fmt.Println("Launching server...")

		// listen on all interfaces
		ln, _ := net.Listen("tcp", ":8081")
	
		// accept connection on port
		conn, _ := ln.Accept()
	
		// run loop forever (or until ctrl-c)
		for {
		//var hash_string, skey_initial string
		protectorserv := new(Protector)
		// will listen for message to process ending in newline (\n)
		hash_string, _ := bufio.NewReader(conn).ReadString('\n')
		skey_initial, _ := bufio.NewReader(conn).ReadString('\n')	
		hash_string = strings.Trim(hash_string,"\n")
		skey_initial = strings.Trim(skey_initial,"\n")
		protectorserv.set_key(skey_initial)
		protectorserv.set_hash(hash_string)  
		// output message received
		fmt.Println("Hash Received:", protectorserv.get_hash())
		fmt.Println("key Received:", protectorserv.get_key())
		new_key := protectorserv.next_session_key()
		fmt.Println("Текущий ключ: " + new_key+"\n")
		// send new string back to client
		conn.Write([]byte(new_key + "\n"))
		}
	}
  }