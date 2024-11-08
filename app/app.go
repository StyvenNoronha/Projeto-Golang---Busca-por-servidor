package app

import "github.com/urfave/cli"
//Gear vai retorna a aplicação na linha de comando 
func Gerar() *cli.App{
	app:= cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e Nomes de servidor na internet"

	return app
}