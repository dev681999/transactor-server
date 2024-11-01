package config_test

import (
	"testing"

	"transactor-server/pkg/infra/config"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

type database struct {
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
}

type server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type testConfig struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
}

func TestNewConfig(t *testing.T) {
	assert := require.New(t)

	c := &testConfig{}

	godotenv.Load("test.env")

	err := config.New(c, config.FromFile("./test.yaml"), config.FromENV("APP"))
	assert.Nil(err)

	assert.Equal("localhost", c.Server.Host)
	assert.Equal("5000", c.Server.Port)

	assert.Equal("test", c.Database.Username)
	assert.Equal("test", c.Database.Password)

	t.Logf("%+v", c)
}
