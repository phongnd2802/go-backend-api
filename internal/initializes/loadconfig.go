package initializes

import (
	"fmt"
	"github.com/phongnd2802/go-backend-api/global"
	"github.com/spf13/viper"
	"log"
)

func loadconfig() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
}
