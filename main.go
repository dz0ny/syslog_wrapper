package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/mcuadros/go-syslog.v2"
)

var wrapCommand string
var syslogBind string

func init() {
	flag.StringVar(&wrapCommand, "cmd", "", "Command to wrap")
	flag.StringVar(&syslogBind, "syslog", "127.0.0.1:514", "Address of internal syslog UDP server")
}

func safeSplit(s string) []string {
	split := strings.Split(s, " ")

	var result []string
	var inquote string
	var block string
	for _, i := range split {
		if inquote == "" {
			if strings.HasPrefix(i, "'") || strings.HasPrefix(i, "\"") {
				inquote = string(i[0])
				block = strings.TrimPrefix(i, inquote) + " "
			} else {
				result = append(result, i)
			}
		} else {
			if !strings.HasSuffix(i, inquote) {
				block += i + " "
			} else {
				block += strings.TrimSuffix(i, inquote)
				inquote = ""
				result = append(result, block)
				block = ""
			}
		}
	}
	return result
}

func run(command, syslog string) {

	if os.Getenv("SYSLOG_SERVER") != "" {
		log.Fatalln("Syslog server is already defined!")
	} else {
		os.Setenv("SYSLOG_SERVER", syslog)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	args := safeSplit(command)
	app := args[0]
	if !filepath.IsAbs(app) {
		app, err = exec.LookPath(app)
		if err != nil {
			log.Panic(err)
		}
	}
	proc, err := os.StartProcess(app, args[1:], &pa)
	if err != nil {
		log.Panic(err)
	}

	_, err = proc.Wait()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	os.Exit(1)
}

func parseSyslog(channel syslog.LogPartsChannel) {
	for logParts := range channel {
		if str, ok := logParts["content"].(string); ok {
			log.Println(str)
		}
	}
}

func main() {
	flag.Parse()
	if wrapCommand != "" {
		channel := make(syslog.LogPartsChannel)
		syslogServer := syslog.NewServer()
		syslogServer.SetFormat(syslog.RFC3164)
		syslogServer.SetHandler(syslog.NewChannelHandler(channel))
		syslogServer.ListenUDP(syslogBind)
		syslogServer.Boot()
		log.Println("Syslog server started on UDP:", syslogBind)
		log.Println("Starting command:", wrapCommand)
		go parseSyslog(channel)
		run(wrapCommand, syslogBind)
	} else {
		log.Fatalln("Wrap command in syslog server and expose as enviroment variable SYSLOG_SERVER.")
	}
}
