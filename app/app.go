package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gear vai retorna a aplicação na linha de comando

func buscarIps(c *cli.Context){
	host := c.String("host")

	ips,erro:= net.LookupIP(host)
	if erro != nil{
		log.Fatal(erro)
	}

	for _,ip := range ips{
		fmt.Println(ip)
	}
}


func Gerar() *cli.App{
	app:= cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e Nomes de servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:"ip",
			Usage: "Busca por Ip endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "google.com",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

