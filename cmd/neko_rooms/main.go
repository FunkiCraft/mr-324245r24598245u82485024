package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	nekoRooms "github.com/m1k1o/neko-rooms"
	"github.com/m1k1o/neko-rooms/cmd"
	"github.com/FunkiCraft/mr-324245r24598245u82485024/internal/utils"
)

func main() {
	fmt.Print(utils.Colorf(nekoRooms.Header, "server", nekoRooms.Service.Version))
	if err := cmd.Execute(); err != nil {
		log.Panic().Err(err).Msg("failed to execute command")
	}
}
