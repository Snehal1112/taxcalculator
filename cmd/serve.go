/*
Copyright © 2020 Snehal Dangroshiya <snehaldangroshiya@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/snehal1112/taxcal/api"
	"github.com/snehal1112/taxcal/app"
	"github.com/snehal1112/taxcal/transport"
	"os"

	"github.com/spf13/cobra"
)

var defaultDBURI = "mongodb://0.0.0.0:27017/?retryWrites=false"
var defaultListenAddr = "127.0.0.1:8774"
var bootstrapConfig = &Config{}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Coderland server.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(cmd, args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	cfg := bootstrapConfig

	serveCmd.Flags().StringVar(&cfg.Listen, "listen", envOrDefault("CODER_LISTEN", defaultListenAddr),  "TCP listen address")
	serveCmd.Flags().StringVar(&cfg.DbURI, "coder_db_uri", envOrDefault("CODER_DB_URI", defaultDBURI), "Data base connection url.")
	serveCmd.Flags().StringVar(&cfg.DbUser, "db_user", os.Getenv("DB_USER"), "Data base user name.")
	serveCmd.Flags().StringVar(&cfg.DbPassword, "db_pass", os.Getenv("DB_PASS"), "Data base user password.")
	serveCmd.Flags().StringVar(&cfg.DbName, "db_name", os.Getenv("DB_NAME"), "Data base name.")

	serveCmd.Flags().Bool("log-timestamp", true, "Prefix each log line with timestamp")
	serveCmd.Flags().String("log-level", "info", "Log level (one of panic, fatal, error, warn, info or debug)")
}

func serve(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	logTimestamp, _ := cmd.Flags().GetBool("log-timestamp")
	logLevel, _ := cmd.Flags().GetString("log-level")

	logger, err := newLogger(!logTimestamp, logLevel)
	if err != nil {
		return fmt.Errorf("failed to create logger: %v", err)
	}
	logger.Infoln("serve start")

	cfg := bootstrapConfig

	tr := transport.NewTransport(
		transport.WithDBName(cfg.DbName),
		transport.WithDBURI(cfg.DbURI),
		transport.WithDBPass(cfg.DbPassword),
		transport.WithDBUser(cfg.DbUser),
		transport.WithLogger(logger))

	apps := app.NewApp(
		app.WithCtx(ctx),
		app.WithLogger(logger),
		app.WithTransport(tr),
	)

	api.Init(apps, apps.Srv.Router)
	apps.StartServer(app.WithListen(cfg.Listen))
	return nil
}