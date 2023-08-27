package remote

import "app.land.x/pkg/ssh"

func (r *Remote) SSHConnect() string {
	// return r.dep.Common.GetRemoteIP()
	return "hello"
}

func (r *Remote) SSHCommand(host string, username string, password string, command string) (string, error) {
	// return r.dep.Common.GetRemoteIP()
	sshClient, err := ssh.NewClient(host, username, password)
	if err != nil {
		return "", err
	}
	session, err := sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Instance.Close()
	return session.RunCommand(command)
}
