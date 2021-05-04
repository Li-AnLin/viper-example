package main

import (
	"fmt"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serverCmd = &cobra.Command{
		Use: "go run ./",
		Run: func(cmd *cobra.Command, args []string) {
			vPort := viper.GetInt("port")
			fmt.Println("vport", vPort)

			var conf config
			if err := viper.Unmarshal(&conf); err != nil {
				panic(err)
			}
			fmt.Printf("config: %+v\n", conf)

			fmt.Println("database port", viper.Get("database.port"))
		},
	}
)

type config struct {
	Port     int64
	Database dbConfig
}

type dbConfig struct {
	Port int64
	Name string
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	var port int64
	flags := serverCmd.Flags()
	flags.Int64VarP(&port, "port", "p", 80, "listening port.")
	viper.BindPFlags(flags)

	if err := serverCmd.Execute(); err != nil {
		panic(err)
	}
}
