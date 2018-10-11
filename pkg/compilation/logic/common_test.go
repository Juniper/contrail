package logic

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/compilation/dependencies"
)

func parseReactions(t *testing.T) dependencies.Reactions {
	reactions, err := dependencies.ParseReactions([]byte(ReactionsYAML), "intent-compiler")
	require.NoError(t, err)

	return reactions
}
