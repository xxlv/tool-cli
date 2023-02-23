package main

import (
	"os"
	"os/user"
	"strings"
)

func AddAlias(aliasString string) error {

	currentUser, err := user.Current()
	if err != nil {
		return nil
	}
	homeDir := currentUser.HomeDir

	I("正在处理命令" + aliasString + " Home:" + homeDir)

	if !strings.Contains(aliasString, "=") {
		E("alias exp must contains =")
		return nil
	}

	s := strings.Split(aliasString, "=")
	c := s[1]
	// 开始
	if !strings.HasPrefix(c, `"`) {
		c = `"` + c
	}
	if !strings.HasSuffix(c, `"`) {
		c = c + `"`
	}
	cmd := "alias " + s[0] + "=" + c
	cnf := homeDir + "/.oh-my-zsh/oh-my-zsh.sh"
	file, err := os.OpenFile(cnf, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		E(err.Error())
	}
	defer file.Close()

	data := []byte(cmd + "\n")
	_, err = file.Write(data)

	if err != nil {
		E(err.Error())
	}
	// source  config
	result, _ := runShell("bash", "-c", "source "+cnf)
	D("bash -c source " + cnf)
	I("flush config" + result)
	return nil
}
