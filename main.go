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
	go run main.go ip --host pudim.com.br

	se colocar sem a flag --host. Volta o ip do google
	*/
	fmt.Println("Inicio")

	aplicacao:= app.Gerar()
	erro := aplicacao.Run(os.Args)

	if erro != nil{
		log.Fatal(erro)
	}
}