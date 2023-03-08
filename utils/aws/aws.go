package external

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"terra-mapper.com/pkg/config"
)

func AwsNewSession() (mySession *session.Session, err error) {
	mySession, err = session.NewSession(
		&aws.Config{
			Region:      aws.String(config.Global.AwsRegion),
			Credentials: credentials.NewStaticCredentials(config.Global.AwsAccessKeyID, config.Global.AwsSecretAccessKey, ""),
		})
	return
}
