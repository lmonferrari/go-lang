package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta pra ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicaçäo de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "deeplearningbook.com.br",
				},
			},
			Action: func(c *cli.Context) {
				host := c.String("host")

				ips, erro := net.LookupIP(host)

				if erro != nil {
					log.Fatal(erro)
				}

				for _, ip := range ips {
					fmt.Println(ip)
				}
			},
		},
	}

	return app
}
