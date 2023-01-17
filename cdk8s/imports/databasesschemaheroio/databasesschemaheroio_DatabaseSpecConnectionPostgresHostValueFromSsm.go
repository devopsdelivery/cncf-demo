// databasesschemaheroio
package databasesschemaheroio


type DatabaseSpecConnectionPostgresHostValueFromSsm struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AccessKeyId *DatabaseSpecConnectionPostgresHostValueFromSsmAccessKeyId `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	SecretAccessKey *DatabaseSpecConnectionPostgresHostValueFromSsmSecretAccessKey `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	WithDecryption *bool `field:"optional" json:"withDecryption" yaml:"withDecryption"`
}

