package intent_test

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/common"
)

func TestMain(m *testing.M) {
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.Set("log_level", "debug")

	common.SetLogLevel()
	os.Exit(m.Run())
}
