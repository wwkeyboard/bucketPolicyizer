package bucketPolicyizer

import "testing"

func TestVersion(t *testing.T) {
	policy := EmptyPolicy()

	if policy.Version != "2012-10-17" {
		t.Error("version not set correctly")
	}
}
