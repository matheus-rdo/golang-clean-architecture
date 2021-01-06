package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/matheushr97/golang-clean-architecture/app"
)

// GetAwsSession -- initialize and loads aws session
func GetAwsSession(awsRegion string) *session.Session {
	return session.Must(session.NewSession(
		&aws.Config{
			Endpoint: aws.String(app.ENV.AwsEndpoint),
			Region:   aws.String(app.ENV.AwsRegion),
		}))
}
