package bucketPolicyizer

var defaultVersion = "2012-10-17"

// Policy represents the Bucket policy
type Policy struct {
	Version string
}

// EmptyPolicy creates a valid empty policy
func EmptyPolicy() Policy {
	return Policy{
		Version: defaultVersion,
	}
}
