package yaml

import (
	"fmt"

	"github.com/ryicoh/clean-arch/internal/infrastructure/conf"
	"github.com/spf13/viper"
)

func New() conf.Config {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(".")
	vp.AddConfigPath("../")
	vp.AddConfigPath("../../")
	vp.AddConfigPath("../../../")
	vp.AddConfigPath("../../../../")
	vp.AddConfigPath("../../../../../")
	err := vp.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error: %s", err))
	}

	c := &conf.ConfigStruct{}
	_ = vp.Unmarshal(c)
	return c
}
