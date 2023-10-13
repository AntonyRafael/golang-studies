package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Gerar - Retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

// BuscarIps - Busca os IPs de um endereço na internet
func buscarIps(c *cli.Context) {
	// Pega a flag host no contexto do comando ip
	host := c.String("host")
	ips, erro := net.LookupIP(host)

	if erro != nil {
		fmt.Println(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// BuscarServidores - Busca os servidores de um endereço na internet
func buscarServidores(c *cli.Context) {
	// Pega a flag host no contexto do comando servidores
	host := c.String("host")
	servidores, erro := net.LookupNS(host)

	if erro != nil {
		fmt.Println(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
