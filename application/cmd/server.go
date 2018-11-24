package cmd

import (
	"github.com/duck8823/duci/application"
	"github.com/duck8823/duci/application/semaphore"
	"github.com/duck8823/duci/infrastructure/logger"
	"github.com/duck8823/duci/presentation/router"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var (
	serverCmd = createCmd("server", "Start server", runServer)
	logo      = `
 ___    __ __    __  ____ 
|   \  |  |  |  /  ]|    |
|    \ |  |  | /  /  |  | 
|  D  ||  |  |/  /   |  | 
|     ||  :  /   \_  |  | 
|     ||     \     | |  | 
|_____| \__,_|\____||____|
`
)

func runServer(cmd *cobra.Command, _ []string) {
	readConfiguration(cmd)

	if err := semaphore.Make(); err != nil {
		logger.Errorf(uuid.New(), "Failed to initialize a semaphore.\n%+v", err)
		os.Exit(1)
		return
	}

	rtr, err := router.New()
	if err != nil {
		logger.Errorf(uuid.New(), "Failed to initialize controllers.\n%+v", err)
		os.Exit(1)
		return
	}

	println(logo)
	if err := http.ListenAndServe(application.Config.Addr(), rtr); err != nil {
		logger.Errorf(uuid.New(), "Failed to run server.\n%+v", err)
		os.Exit(1)
		return
	}
}
