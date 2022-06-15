package tcp_conn

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"strings"
	"wechat_server/core/user"
	"wechat_server/interactive/convert"
	"wechat_server/interactive/imodel"
)

// tcpConnMap ipaddress:port : 对应tcp连接
var tcpConnMap map[string]*TcpConn

type TcpInfo struct {
	Ip   string
	Port string
}

type TcpConn struct {
	conn net.Conn
}

func CreateTcpConn(conn net.Conn) *TcpConn {
	return &TcpConn{conn: conn}
}

func TcpSendMsg(ip, port string, msg imodel.Message) error {
	tc, err := GetConnByUser(&TcpInfo{
		Ip:   ip,
		Port: port,
	})
	if err != nil {
		panic(any(err))
	}
	return tc.SendMsg(msg)
}

// SendMsg 发送tcp消息
func (tc *TcpConn) SendMsg(msg imodel.Message) (err error) {
	sendUser := convert.FromUser(user.GetMyInfo())
	msg.SendUser = sendUser
	bytes, err := json.Marshal(msg)
	if err != nil {
		err = errors.Wrap(err, "[TcpConn] [SendMsg] json.Marshal(msg) error")
		return err
	}
	_, err = tc.conn.Write(bytes)
	if err != nil {
		err = errors.Wrap(err, "[TcpConn] [SendMsg] tc.conn.Write(bytes) error")
		return err
	}
	return
}

func (tc *TcpConn) ParseMsg() (msg imodel.Message, err error) {
	// 解析消息来源
	remoteAddr := tc.conn.RemoteAddr().String()
	remoteSli := strings.Split(remoteAddr, ":")
	remoteIp := remoteSli[0]

	// 解析消息内容
	// todo:验证一个问题：当正在处理conn中接收的消息正在进行时，同一个连接发送过来新的请求内容会发生什么事？
	// todo:限制每条消息长度 另外如果需要接收大文件，考虑使用流式读取
	buf := make([]byte, 20000)
	length, err := tc.conn.Read(buf)
	if err != nil {
		err = errors.Wrap(err, "[TcpConn] [ParseMsg] tc.conn.Read(buf) error")
		return msg, err
	}
	err = json.Unmarshal(buf[:length], &msg)
	if err != nil {
		err = errors.Wrap(err, "[TcpConn] [ParseMsg] json.Unmarshal(buf[:length], &msg) error")
		return msg, err
	}

	// 校验
	if remoteIp != msg.SendUser.Ip {
		err = errors.New("[TcpConn] [ParseMsg] 消息发送方伪造信息")
		return msg, err
	}
	return msg, err
}

func GetConnByUser(tcpInfo *TcpInfo) (*TcpConn, error) {
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
	tc := &TcpConn{
		conn: conn,
	}
	return tc, err
}
