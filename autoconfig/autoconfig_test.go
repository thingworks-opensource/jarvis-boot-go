package autoconfig

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"thingworks.net/thingworks/jarvis-boot/autoconfig/config"
)

func TestDefaultConfig(t *testing.T) {
	location := "./config_test.yaml"
	appConfig := config.Init(config.AppArgs{ConfigLocation: &location})
	assert.NotNil(t, appConfig)
	assert.NotNil(t, appConfig.Mongodb)
	assert.NotNil(t, appConfig.App)
	assert.Equal(t, "localhost", appConfig.Mongodb.Host)
	assert.Equal(t, "27017", appConfig.Mongodb.Port)
	assert.Equal(t, 9090, appConfig.App.Port)
	assert.Equal(t, "Test", appConfig.App.Name)
}
