package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Anondo/graphql-and-go/config"
	"github.com/Anondo/graphql-and-go/routes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCMD represents the serve command
var serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "Starts the http server...",
	Run:   serve,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		appCfg := config.App()

		portStr := strconv.Itoa(appCfg.Port)
		listener, err := net.Listen("tcp", ":"+portStr)
		if err != nil {
			return fmt.Errorf("port %s is not available", portStr)
		}
		_ = listener.Close()

		return nil
	},
}

func init() {

	serveCMD.PersistentFlags().IntP("p", "p", 8080, "port on which the server will listen for http")

	err := viper.BindPFlag("app.port", serveCMD.PersistentFlags().Lookup("p"))
	if err != nil {
		log.Fatal(err.Error())
	}
	RootCMD.AddCommand(serveCMD)
}

func serve(cmd *cobra.Command, args []string) {

	appCfg := config.App()
	portStr := strconv.Itoa(appCfg.Port)

	e := routes.Router()

	server := &http.Server{
		ReadTimeout:  appCfg.ReadTimeout,
		WriteTimeout: appCfg.WriteTimeout,
		IdleTimeout:  appCfg.IdleTimeout,
		Addr:         ":" + portStr,
		Handler:      e,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Println("Listening on port:" + fmt.Sprintf(":%d", appCfg.Port))
	log.Println("Graphql Playground running on http://localhost:8080/")
	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = server.Shutdown(ctx)

	log.Println("Server shutdowns gracefully")
}
