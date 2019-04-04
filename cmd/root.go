package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/liquidslr/storeservice/db"
	"github.com/liquidslr/storeservice/routes"
	"github.com/spf13/cobra"
)

type callBackChan chan struct{}

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "Prints value of key stored in db",
	Run: func(cmd *cobra.Command, args []string) {
		value, err := GetValue(key)
		if err != nil {
			fmt.Println("The key does not exists")
		}

		fmt.Println(value)
	},
}

var cmdPut = &cobra.Command{
	Use:   "put",
	Short: "Create a key value pair",
	Run: func(cmd *cobra.Command, args []string) {
		err := CreateKV(key, value)
		if err != nil {
			fmt.Println("The key value pair could not be created")
		}
	},
}

var cmdServer = &cobra.Command{
	Use:   "start",
	Short: "Start the web server, The default port is 3000",
	Run: func(cmd *cobra.Command, args []string) {
		createDB()
		routes.Server("3000")
	},
}

var cmdWatch = &cobra.Command{
	Use:   "watch",
	Short: "Watch for changes in key value pairs",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		callback := make(callBackChan)

		go checkEvery(ctx, 1*time.Second, callback)
		go func() {
			for {
				select {
				case <-callback:
					Subscribe()
				}
			}
		}()

		for {
			time.Sleep(1 * time.Second)
		}

	},
}

var (
	key   string
	value string
)

func init() {
	cmdGet.Flags().StringVarP(&key, "key", "r", "", "get value of any key")
	cmdPut.Flags().StringVarP(&value, "value", "r", "", "create a value")
	cmdPut.Flags().StringVar(&key, "key", "", "create a key")

	cmdGet.MarkFlagRequired("key")
	cmdPut.MarkFlagRequired("key")
	cmdPut.MarkFlagRequired("value")

}

func createDB() {
	routes.DBClient = &db.BoltDB{}
	routes.DBClient.Initialize()
	routes.DBClient.GetAll()
	fmt.Println("Db instance created")
}

func checkEvery(ctx context.Context, d time.Duration, cb callBackChan) {
	for {
		select {
		case <-ctx.Done():
			// ctx is cancelled
			return
		case <-time.After(d):
			if cb != nil {
				cb <- struct{}{}
			}
		}
	}
}

// Execute is used to call the args and flags
func Execute() {
	var rootCmd = &cobra.Command{Use: "store"}
	rootCmd.AddCommand(cmdGet, cmdPut, cmdServer, cmdWatch)
	rootCmd.Execute()
}
