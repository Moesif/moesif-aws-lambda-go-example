package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"context"
	moesifawslambda "github.com/moesif/moesif-aws-lambda-go"
	moesifOptions "github.com/moesif/moesif-aws-lambda-go-example/moesif_options"
)

func HandleLambdaEvent(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: 200,
		Headers: map[string] string {
			"RspHeader1":     "RspHeaderValue1",
			"Content-Type":   "application/json",
			"Content-Length": "1000",
		},
	   }, nil
}

func main() {
	lambda.Start(moesifawslambda.MoesifLogger(HandleLambdaEvent, moesifOptions.MoesifOptions()))
}