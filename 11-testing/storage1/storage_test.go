package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// Save and restore original notifyUser.
	savedNotifiyUser := notifyUser
	savedBytesInUse := bytesInUse
	defer func() {
		notifyUser = savedNotifiyUser
		bytesInUse = savedBytesInUse
	}()

	// Install the test's fake notifyUser.
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	bytesInUse = func(username string) int64 {
		return 980000000 // simulate a 980MB-used condition
	}

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s",
			notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}
