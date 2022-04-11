package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	nekoRooms "github.com/FunkiCraft/mr-324245r24598245u82485024"
	"github.com/FunkiCraft/mr-324245r24598245u82485024/internal/config"
)

func init() {
	command := &cobra.Command{
		Use:   "serve",
		Short: "serve neko-rooms server",
		Long:  `serve neko-rooms server`,
		Run:   nekoRooms.Service.ServeCommand,
	}

	configs := []config.Config{
		nekoRooms.Service.Configs.Server,
		nekoRooms.Service.Configs.API,
		nekoRooms.Service.Configs.Room,
	}

	cobra.OnInitialize(func() {
		for _, cfg := range configs {
			cfg.Set()
		}
		nekoRooms.Service.Preflight()
	})

	for _, cfg := range configs {
		if err := cfg.Init(command); err != nil {
			log.Panic().Err(err).Msg("unable to run serve command")
		}
	}

	root.AddCommand(command)
}
