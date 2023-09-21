//NOME DO ARQUIVO
package app

import "github.com/urfave/cli"

//Gerar vai retornar a linha de comando
func Gerar() *cli.App {
 app := cli.NewApp();
 app.Name = "Aplicação de linha de comando"
 app.Usage = "Busca ipd de servidor na internet"

 return app
}

