package graphqlbackend

import (
	"context"

	"github.com/tetrafolium/sourcegraph-cloned/cmd/frontend/internal/siteid"
	"github.com/tetrafolium/sourcegraph-cloned/internal/actor"
	"github.com/tetrafolium/sourcegraph-cloned/internal/db"
	"github.com/tetrafolium/sourcegraph-cloned/internal/errcode"
	"github.com/tetrafolium/sourcegraph-cloned/internal/hubspot/hubspotutil"
)

type trialRequestForHubSpot struct {
	Email  *string `url:"email"`
	SiteID string  `url:"site_id"`
}

// RequestTrial makes a submission to the request trial form.
func (r *schemaResolver) RequestTrial(ctx context.Context, args *struct {
	Email string
}) (*EmptyResponse, error) {
	email := args.Email

	// If user is authenticated, use their uid and overwrite the optional email field.
	if actor := actor.FromContext(ctx); actor.IsAuthenticated() {
		e, _, err := db.UserEmails.GetPrimaryEmail(ctx, actor.UID)
		if err != nil && !errcode.IsNotFound(err) {
			return nil, err
		}
		if e != "" {
			email = e
		}
	}

	// Submit form to HubSpot
	if err := hubspotutil.Client().SubmitForm(hubspotutil.TrialFormID, &trialRequestForHubSpot{
		Email:  &email,
		SiteID: siteid.Get(),
	}); err != nil {
		return nil, err
	}

	return &EmptyResponse{}, nil
}
