package service

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"go-websocket/core"
)

type ConsoleService struct {
	Config				*core.Config
	WebSocketService	*WebSocketService
}

func (t *ConsoleService) Run(){
	//执行协程
	var ch chan string
	go t.main(ch)
}

func (t *ConsoleService) main(ch chan string){
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

func (t *ConsoleService) welcome(){
	port := strconv.Itoa(t.Config.Port)
	fmt.Print("-------------------------------- GO WEBSOCKET --------------------------------------- \n")
	fmt.Print("Service  	Listen                       	processes 	status \n")
	fmt.Print("WebSocket	ws://0.0.0.0:"+ port +"            	1   	     	[ok] \n")
	if t.Config.Test {
		fmt.Print("Test     	http://127.0.0.1:"+ port +"/test   	1   	     	[ok] \n")
	}
	if t.Config.Swagger {
		fmt.Print("Swagger  	http://127.0.0.1:"+ port +"/swagger	1   	     	[ok] \n")
	}
	fmt.Print("------------------------------------------------------------------------------------- \n")
	fmt.Print("[*] Waiting for commad. To exit press CTRL+C \n")
	fmt.Print("[*] Please enter a command( enter 'help' to view ): \n")
}

func (t *ConsoleService) help(){
	fmt.Print("[*] help		Help message. \n")
	fmt.Print("[*] count		Client coonect count. \n")
}

func (t *ConsoleService) count(){
	fmt.Print("[*] Client connect count : " + strconv.Itoa(len(t.WebSocketService.Connects)) + "\n")
}
