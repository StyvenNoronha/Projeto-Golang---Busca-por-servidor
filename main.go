package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main(){
	/*
	so executa na linha de comando
	go run main.go ip --host example.com.br


	so executar na linha de comando 
	go run main.go server  --host google.com
	*/

	//Para executar so usar linha-de-comando.exe ip host example
	fmt.Println("Inicio")

	aplicacao:= app.Gerar()
	erro := aplicacao.Run(os.Args)

	if erro != nil{
		log.Fatal(erro)
	}
}