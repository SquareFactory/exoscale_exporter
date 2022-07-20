package main

import (
	"net/http"
	"os"

	"github.com/SquareFactory/exoscale_exporter/config"
	"github.com/SquareFactory/exoscale_exporter/exoscale"
	"github.com/SquareFactory/exoscale_exporter/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	collector *exoscale.Collector
)

func main() {
	var configFile string
	var listenAddress string

	log.InitLogger()

	app := &cli.App{
		Name:  "exoscale_exporter",
		Usage: "Fetches statistics from Exoscale API",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config.file",
				Value:       "config.yaml",
				Usage:       "Path to configuration file.",
				Destination: &configFile,
				EnvVars:     []string{"CONFIG_FILE"},
			},
			&cli.StringFlag{
				Name:        "web.listen-address",
				Value:       ":9116",
				Usage:       "Address to listen on for web interface and telemetry.",
				Destination: &listenAddress,
				EnvVars:     []string{"LISTEN_ADDRESS"},
			},
		},
		Action: func(*cli.Context) error {
			config, err := config.LoadFile(configFile)
			if err != nil {
				return err
			}
			computeAPI := exoscale.NewComputeAPI(config.ExoscaleConfig.Key, config.ExoscaleConfig.Secret)
			collector = exoscale.NewCollector(computeAPI)
			go collector.RecordMetrics()

			http.Handle("/metrics", promhttp.Handler())
			log.Logger.Info(
				"Serving http server",
				zap.String("web.listen-address", listenAddress),
				zap.String("config.file", configFile),
			)
			return http.ListenAndServe(listenAddress, nil)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Logger.Fatal("App crashed", zap.Error(err))
	}
}
