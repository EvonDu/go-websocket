package service

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Console struct {
	WebSocketService *WebSocketService
}

func (t *Console) Run(ch chan string){
	//帮助信息
	t.welcome()
	//等待命令
	for{
		//获取键盘输入
		//fmt.Print("commad-> ")
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		input = strings.Replace(input, " ", "", -1)
		input = strings.Replace(input, "\n", "", -1)
		input = strings.Replace(input, "\r", "", -1)
		//执行帮助
		switch input {
		case "help":
			t.help()
		case "count":
			t.count()
		default:
			fmt.Print("[*] Error : Unknown command, please output 'help' to view the document. \n")
		}
	}
}

func (t *Console) welcome(){
	fmt.Print("-------------------------------- GO WEBSOCKET --------------------------------------- \n")
	fmt.Print("Service        Listen                          processes      status \n")
	fmt.Print("WebSocket      ws://0.0.0.0:8080               1              [ok] \n")
	fmt.Print("Http           http://127.0.0.1:8080/test      1              [ok] \n")
	fmt.Print("[*] Waiting for commad. To exit press CTRL+C \n")
	fmt.Print("[*] Please enter a command( enter 'help' to view ): \n")
}

func (t *Console) help(){
	fmt.Print("[*] help		Help message. \n")
	fmt.Print("[*] count		Client coonect count. \n")
}

func (t *Console) count(){
	fmt.Print("[*] Client connect count : " + strconv.Itoa(len(t.WebSocketService.Connects)) + "\n")
}