package socks5

import (
	"fmt"
	"log"
	"net"
)

type Server interface {
	Run() error
}

type ServerName struct {
	IP   string
	Port int
}

func (s *ServerName) Run() error {
	address := fmt.Sprintf("%s:%d", s.IP, s.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for { //不停接收客户端的请求
		coon, err := listener.Accept() //三次握手获取连接
		if err != nil {
			log.Printf("conection failure from %s:%s", coon.RemoteAddr(), err)
			continue
		}

		//成功获取连接,开始协商等步骤

		go func() { //创建一个方程捕获错误(能写if判断)
			defer coon.Close() //延迟连接关闭，资源回收
			err := handleConnection(coon)
			if err != nil {
				log.Printf("handle conection failure from %s:%s", coon.RemoteAddr(), err)
			}

		}()

	}

}

func handleConnection(coon net.Conn) error {
	//协商过程
	if err := auth(coon); err != nil {
		return err
	}
	//请求过程

	//转发过程

	return nil //没有错误返回空值
}

func auth(coon net.Conn) error {
	//主要从连接中读取报文，后向连接中写入报文

	return nil
}
