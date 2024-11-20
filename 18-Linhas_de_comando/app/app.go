package app

//Sempre referencia o nome do pacote dps da ultima "/", no caso o "cli"
import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de compando pronto para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Aplicação de linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.devbook.com.br",
		},
	}

	/*
		Como seria o comando:

		Para buscar IPS:
			go run main.go ip --host {ALGUM_SITE}
				ou sem parametro (ai pegando o default que é o www.devbook.com.br)
			go run main.go ip
		Para buscar Servidores:
			go run main.go servidores --host {ALGUM_SITE}
				ou sem parametro (ai pegando o default que é o www.devbook.com.br)
			go run main.go servidores

	*/
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Buscar o nome do servidor na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println("IP:>>> ", ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //server name

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println("Server:>>> ", servidor)
	}
}
