package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retorna a aplicação na linha de comando
	/*
	so executa na linha de comando
	go run main.go ip --host example.com.br


	so executar na linha de comando 
	go run main.go server  --host google.com
	*/


func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e Nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca por Ip endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "server",
			Usage:  "Buscar por servidores na internet",
			Flags:  flags,
			Action: buscarServer,
		},
	}

	return app
}

// Busca por ips na net

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
	

func buscarServer(c *cli.Context) {
	server:=c.String("host")
	serves,erro:= net.LookupNS(server)//NS name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range serves {
		fmt.Println(server.Host)
	}
}