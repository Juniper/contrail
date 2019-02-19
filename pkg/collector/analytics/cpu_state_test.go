package analytics

import (
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/version"
	"github.com/stretchr/testify/assert"
)

func TestCPUState(t *testing.T) {
	c := &mockCollector{}
	t.Run("ModuleCpuStateTrace", func(t *testing.T) {
		c.Send(ModuleCPUStateTrace("localhost:9091"))
		assert.NotNil(t, c.message)
		assert.Equal(t, c.message.SandeshType, typeModuleCPUStateTrace)
		m, ok := c.message.Payload.(*payloadModuleCPUStateTrace)
		assert.True(t, ok)
		assert.Equal(t, m.Name, "localhost")
		assert.Equal(t, m.Table, "ObjectConfigNode")
		assert.Equal(t, m.BuildInfo.BuildVersion, version.Version)
		_, err := time.Parse("2006-01-02 15:04:05.999999999", m.BuildInfo.BuildTime)
		assert.NoError(t, err)
		assert.Equal(t, m.BuildInfo.BuildUser, version.BuildUser)
		assert.Equal(t, m.BuildInfo.BuildHostname, version.BuildHostname)
		assert.Equal(t, m.BuildInfo.BuildID, version.BuildID)
		assert.Equal(t, m.BuildInfo.BuildNumber, version.BuildNumber)
	})
}
