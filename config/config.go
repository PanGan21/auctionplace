package config

import (
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// environmentPrefix prefix used to avoid environment variable names collisions
const environmentPrefix = "SW"

var (
	// Filename configuration file name.
	Filename string
	// App configuration struct
	App AppConfig

	// environmentVarList list of environment variables read by the app. The name should match with a struct field.
	// The dots will be replaced by underscores, it will be capitalized and the environmentPrefix will be added
	// 		i.e.: blockchain.pk => SW_BLOCKCHAIN_PK
	environmentVarList = []string{
		"blockchain.pk",
	}
)

type AppConfig struct {
	Blockchain BlockchainConfig
	Contract   ContractConfig
}

type BlockchainConfig struct {
	Address    string `mapstructure:"address"`
	Ws         string `mapstructure:"ws"`
	PrivateKey string `mapstructure:"pk"`
	Timeout    string `mapstructure:"timeout"`
	TimeoutIn  time.Duration
}

// ContractConfig struct
type ContractConfig struct {
	Address   string `mapstructure:"address"`
	GasLimit  int64  `mapstructure:"gas_limit"`
	GasPrice  int64  `mapstructure:"gas_price"`
	WeiFounds int64  `mapstructure:"default_wei_founds"`
}

// Setup bind command flags and environment variables
func Setup(cmd *cobra.Command, _ []string) error {
	v := viper.New()
	v.SetConfigFile(Filename)
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	v.SetEnvPrefix(environmentPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		_ = v.BindPFlag(flag.Name, cmd.Flags().Lookup(flag.Name))
	})

	for _, env := range environmentVarList {
		_ = v.BindPFlag(env, cmd.Flags().Lookup(env))
	}

	err = v.Unmarshal(&App)
	if err != nil {
		return err
	}
	App.Blockchain.TimeoutIn, err = time.ParseDuration(App.Blockchain.Timeout)
	if err != nil {
		return err
	}
	return nil
}
