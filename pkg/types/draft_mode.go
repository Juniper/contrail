package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/errutil"
)

type draftModeStateGetter interface {
	GetDraftModeState() string
}

// DraftModeStateChecker checks if request contains draftModeState property
type DraftModeStateChecker interface {
	CheckDraftModeState(context.Context, draftModeStateGetter) error
}

func checkDraftModeState(ctx context.Context, dms draftModeStateGetter) error {
	if IsInternalRequest(ctx) {
		return nil
	}

	if dms.GetDraftModeState() != "" {
		return errutil.ErrorBadRequest(
			"security resource property 'draft_mode_state' is only readable",
		)
	}

	return nil
}
