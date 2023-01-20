package socks5

import (
	"errors"
	"io"
)

type ClientAuthMessage struct {
	Version  byte   //版本号
	NMethods byte   //method的数量
	Methods  []byte //可变长度，用byte slice表示
}

// NewClientAuthMessage 读取报文并生成报文结构体
func NewClientAuthMessage(conn io.Reader) (*ClientAuthMessage, error) {
	//读取Version和NMethods
	buf := make([]byte, 2)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}
	//验证version是否合法
	if buf[0] != socks5Version {
		return nil, errors.New("protocol version not supported")

	}
	//读取NMethods个方法
	nMethods := buf[1]
	buf = make([]byte, nMethods)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}

	return &ClientAuthMessage{
		Version:  socks5Version,
		NMethods: nMethods,
		Methods:  buf,
	}, nil

}
