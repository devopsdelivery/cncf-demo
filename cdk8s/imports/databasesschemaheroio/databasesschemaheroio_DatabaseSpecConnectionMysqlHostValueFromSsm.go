// databasesschemaheroio
package databasesschemaheroio


type DatabaseSpecConnectionMysqlHostValueFromSsm struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AccessKeyId *DatabaseSpecConnectionMysqlHostValueFromSsmAccessKeyId `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	SecretAccessKey *DatabaseSpecConnectionMysqlHostValueFromSsmSecretAccessKey `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	WithDecryption *bool `field:"optional" json:"withDecryption" yaml:"withDecryption"`
}

