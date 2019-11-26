package main

import (
	"fmt"
	"time"
	"treinamentoBcGo/logger"
	"treinamentoBcGo/person"
)

func main() {
	fmt.Println("Bem-vindo")

	var jogador = person.NewPerson("Jogador", 18, 10)

	fmt.Println(jogador.Age())
	jogador.Run()

	var l logger.Logger = logger.NewJSON("treinamento")
	l.Info("xablau", map[string]interface{}{
		"timeout": 5 * time.Second,
		"err":     "nenhum",
	})
}
