package uipc

import (
	"net"
	"time"
)

const (
	TIMEOUT_DEFAULT = 10
	BUFSIZE_DEFAULT = 65535
)

type SESSION struct {
	timeout time.Duration
	url     string
	conn    net.Conn
	err     error
	wbuf    chan []byte
	rbuf    chan []byte

	uipc *UIPC
}

type UIPC struct {
	IP   string
	PORT string

	lis net.Conn

	session map[string]*SESSION
}

func NewUIPC(ip, port string) *UIPC {
	return &UIPC{IP: ip, PORT: port}
}

func RecvTask(s *SESSION) {

}

func (u *UIPC) NewSession(ip, port string) (*SESSION, error) {
	var err error
	url := ip + ":" + port

	s1, b := u.session[url]
	if b == true {
		return s1, nil
	}

	var s = &SESSION{url: url}

	s.timeout = time.Duration(TIMEOUT_DEFAULT * time.Second)
	s.wbuf = make(chan []byte, BUFSIZE_DEFAULT)
	s.rbuf = make(chan []byte, BUFSIZE_DEFAULT)
	s.uipc = u
	s.err = nil

	s.conn, err = net.DialTimeout("udp", url, s.timeout)
	if err != nil {
		return nil, err
	}

	u.session[url] = s

	go RecvTask(s)

	return s, nil
}

func (s *SESSION) Close() {
	s.Close()
	delete(s.uipc.session, s.url)
}
