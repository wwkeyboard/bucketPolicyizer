package bucketPolicyizer

import (
	"fmt"
	"regexp"
	"testing"
)

// Many of these test cases are taken from
// http://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html

func TestVersion(t *testing.T) {
	policy := EmptyPolicy()

	if policy.Version != "2012-10-17" {
		t.Error("version not set correctly")
	}
}

func TestReadOnlyFromAnonymous(t *testing.T) {
	policy := EmptyPolicy()
	action := GetObjectAction
	resource := "arn:aws:s3::exampleBucket/*"
	s := Statement{
		Sid:       "AddCannedAcl",
		Effect:    "Allow",
		Principal: "*",
		Action:    []string{action},
		Resource:  []string{resource},
	}

	policy.Statement = []Statement{s}

	p, err := CompilePolicy(policy)
	if err != nil {
		t.Error("couldn't compile policy", err)
	}

	sidTest := regexp.MustCompile(`AddCannedAcl`)
	if !sidTest.MatchString(p) {
		fmt.Println(p)
		t.Error("couldn't match Sid")
	}
}
