package pkg

import (
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

type DocNoc struct {
	Client       *client.Client
	DocNocConfig *DocNocConfig
	Flags        *Flags
}

func NewDocNoc() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("🔥: Can't connect to docker client")
		os.Exit(1)
	}

	fmt.Println("👍: Ready to read file", cli)

}
