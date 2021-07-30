package test

type TestCase struct {
	Create                        bool
	Bucket                        string
	Acl                           string
	Tags                          map[string]string
	BlockPublicAcls               bool
	BlockPublicPolicy             bool
	IgnorePublicAcls              bool
	RestrictPublicBuckets         bool
	VersioningEnabled             bool
	ServerSideEncryptionEnabled   bool
	ServerSideEncryptionAlgorithm string
	LoggingEnabled                bool
	LoggingBucket                 string
}
