package command

import (
	"log"

	"github.com/spf13/cobra"

	application "take-home-test/app"
)

// @title Sports Field Booking API
// @version 1.0
// @description RESTful API for sports field booking system - Take Home Test

// @contact.name API Support
// @contact.email support@sagaratech.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme. Example: "Bearer {token}"

// @host localhost:3005
// @BasePath /api

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
}

var cmdRoot = &cobra.Command{
	Use:   "take-home-test",                                                                 // ✅ UBAH DARI "master-service" KE "take-home-test"
	Short: "Sports Field Booking API",                                                       // ✅ UPDATE SHORT DESCRIPTION
	Long:  `RESTful API for sports field booking system - Take Home Test Backend Developer`, // ✅ UPDATE LONG DESCRIPTION
	Run: func(cmd *cobra.Command, args []string) {
		app := application.New()
		err := app.Init()
		if err != nil {
			log.Fatalf("Error in initializing the application: %+v", err)
			return
		}

		err = app.Run()
		if err != nil {
			log.Fatalf("Error in running the application: %+v", err)
			return
		}
	},
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		log.Fatalf("Error in executing the root command: %+v", err)
	}
}
