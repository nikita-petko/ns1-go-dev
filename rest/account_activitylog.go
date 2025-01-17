package rest

import (
	"net/http"

	"gopkg.in/nkpetko/ns1-go-dev.v2/rest/model/account"
)

// ActivityLogService handles 'account/activity' endpoint.
type ActivityLogService service

// List returns all activity logs in the account.
//
// NS1 API docs: https://ns1.com/api/#activity-get
func (s *ActivityLogService) List() ([]*account.ActivityLog, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "account/activity", nil)
	if err != nil {
		return nil, nil, err
	}

	kl := []*account.ActivityLog{}
	resp, err := s.client.Do(req, &kl)
	if err != nil {
		return nil, resp, err
	}

	return kl, resp, nil
}
