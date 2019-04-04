package cmd

import (
	"fmt"

	"github.com/liquidslr/storeservice/db"
	"github.com/liquidslr/storeservice/routes"
	"github.com/spf13/cobra"
)

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
	fmt.Println("Db instance created")
}

// Execute is used to call the args and flags
func Execute() {
	var rootCmd = &cobra.Command{Use: "store"}
	rootCmd.AddCommand(cmdGet, cmdPut, cmdServer)
	rootCmd.Execute()
}
