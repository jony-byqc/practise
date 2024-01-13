package main

import (
	"bytes"
	"strings"
	"testing"
)

const (
	username = "wjq"
	password = "123456"
	ip       = "192.168.60.157"
	port     = 22
	cmds     = "show clock;sudo cd /data;123456;ls;exit"
)

func Test_SSH(t *testing.T) {
	var cipherList []string
	session, err := connect(username, password, ip, "", port, cipherList)
	if err != nil {
		t.Error(err)
		return
	}
	defer session.Close()

	cmdlist := strings.Split(cmds, ";")
	stdinBuf, err := session.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}

	var outbt, errbt bytes.Buffer
	session.Stdout = &outbt

	session.Stderr = &errbt
	err = session.Shell()
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range cmdlist {
		c = c + "\n"
		stdinBuf.Write([]byte(c))
	}
	session.Wait()
	t.Log((outbt.String() + errbt.String()))
	return
}
