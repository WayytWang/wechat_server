package cache

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
)

type TcpInfo struct {
	Ip   string
	Port string
}

// tcpConnMap ipaddress:port : 对应tcp连接
var tcpConnMap map[string]net.Conn

func GetConnByUser(tcpInfo *TcpInfo) (net.Conn, error) {
	// todo:缓存tcp连接提升效率
	//userId := user.ID
	//conn,ok := tcpConnMap[userId]
	//if ok {
	//	// todo :检验连接是否存活
	//	return conn,nil
	//}
	address := fmt.Sprintf("%s:%s", tcpInfo.Ip, tcpInfo.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		err = errors.New(fmt.Sprintf("dial %s error", address))
		return nil, err
	}
	return conn, err
}
