package log_test

import (
	"testing"

	"transactor-server/pkg/infra/log"

	"github.com/stretchr/testify/require"
)

func TestNewLog(t *testing.T) {
	assert := require.New(t)
	logger, _ := log.New(true, nil, nil)
	assert.NotNil(logger)
}
