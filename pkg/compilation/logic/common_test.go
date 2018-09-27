package logic

import (
	"testing"

	"github.com/Juniper/contrail/pkg/compilation/dependencies"
	"github.com/stretchr/testify/require"
)

func parseReactions(t *testing.T) dependencies.Reactions {
	reactions, err := dependencies.ParseReactions("../../../tools/dependencies.yml", "intent-compiler")
	require.NoError(t, err)

	return reactions
}
