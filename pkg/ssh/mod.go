package ssh

import (
	"bytes"

	"golang.org/x/crypto/ssh"
)

// Client SSH 客户端
type Client struct {
	host     string
	username string
	password string
	Instance *ssh.Client
}

// NewSession 创建新的 SSH 会话
func (p *Client) NewSession() (*Session, error) {
	session, err := p.Instance.NewSession()
	return &Session{
		Instance: session,
	}, err
}

// NewClient 创建新的 SSH 客户端
func NewClient(host, username, password string) (*Client, error) {
	// SSH 连接配置
	config := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host+":22", config)
	return &Client{
		host:     host,
		username: username,
		password: password,
		Instance: client,
	}, err
}

// Session SSH 会话
type Session struct {
	Instance *ssh.Session
}

// RunCommand session 会话下执行 ssh 命令
func (session *Session) RunCommand(command string) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Instance.Stdout = &stdoutBuf
	err := session.Instance.Run(command)
	return stdoutBuf.String(), err
}
