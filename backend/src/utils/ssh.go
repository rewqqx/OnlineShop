package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"time"
)

const (
	CERT_PASSWORD        = 1
	CERT_PUBLIC_KEY_FILE = 2
	DEFAULT_TIMEOUT      = 3 // second
)

type SSH struct {
	Ip        string
	User      string
	Password  string
	Port      int
	Signer    ssh.Signer
	PublicKey ssh.PublicKey
	session   *ssh.Session
	client    *ssh.Client
}

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(self, s)
}

func (self *ViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func (self *ViaSSHDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func (sshClient *SSH) readPublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func (sshClient *SSH) Connect(mode int) error {

	var sshConfig *ssh.ClientConfig
	var auth []ssh.AuthMethod
	if mode == CERT_PASSWORD {
		auth = []ssh.AuthMethod{ssh.Password(sshClient.Password)}
	} else if mode == CERT_PUBLIC_KEY_FILE {
		auth = []ssh.AuthMethod{
			ssh.PublicKeys(sshClient.Signer),
		}
	} else {
		return errors.New("Does not support mode")
	}

	sshConfig = &ssh.ClientConfig{
		User:            sshClient.User,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * DEFAULT_TIMEOUT,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshClient.Ip, sshClient.Port), sshConfig)
	if err != nil {
		return err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return err
	}

	sshClient.session = session
	sshClient.client = client

	return nil
}

func (sshClient *SSH) RunCmd(cmd string) {
	out, err := sshClient.session.CombinedOutput(cmd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func (sshClient *SSH) Close() {
	sshClient.session.Close()
	sshClient.client.Close()
}
