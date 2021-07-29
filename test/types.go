package test

type PlanTestCase struct {
	TestCase
}

type ApplyTestCase struct {
	TestCase
}

type TestCase struct {
	Create                            bool
	Bucket                            string
	Acl                               string
	Tags                              map[string]string
	BlockPublicAcls                   bool
	BlockPublicPolicy                 bool
	IgnorePublicAcls                  bool
	RestrictPublicBuckets             bool
	VersioningEnabled                 bool
	ServerSideEncryptionEnabled       bool
	ServerSideEncryptionConfiguration struct { Rule struct { ApplyServerSideEncryptionByDefault struct {KMSMasterKeyID string; SSEAlgorithm string}}}
	LoggingConfiguration              struct { TargetBucket string; TargetPrefix string}
}
