package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	moesifawslambda "github.com/moesif/moesif-aws-lambda-go"
	moesifOptions "github.com/moesif/moesif-aws-lambda-go-example/moesif_options"
	moesifOptionsV2 "github.com/moesif/moesif-aws-lambda-go-example/moesif_options_v2"
)

func HandleLambdaEvent(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	moesifawslambda.StartCaptureOutgoing(moesifOptions.MoesifOptions())

	// Outgoing API call to third parties like Github / Stripe or to your own dependencies
	_, err := http.Get("https://api.github.com")

	// Check for any errors while sending outgoing request
	if err != nil {
		log.Printf("Error while sending outgoing request : %s.\n", err.Error())
	}

	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: 200,
		Headers: map[string]string{
			"RspHeader1":   "RspHeaderValue1",
			"Content-Type": "application/json",
		},
	}, nil
}

func HandleLambdaEventV2HTTP(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	moesifawslambda.StartCaptureOutgoing(moesifOptionsV2.MoesifOptions())

	// Outgoing API call to third parties like Github / Stripe or to your own dependencies
	_, err := http.Get("https://api.github.com")

	// Check for any errors while sending outgoing request
	if err != nil {
		log.Printf("Error while sending outgoing request : %s.\n", err.Error())
	}

	return events.APIGatewayV2HTTPResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"RspHeader1":   "TheYearofDesktopLinux",
				"Content-Type": "application/json",
			},
			MultiValueHeaders: map[string][]string{
				"X-Forwarded-For":   {"127.0.0.1, 127.0.0.2"},
				"X-Forwarded-Port":  {"443"},
				"X-Forwarded-Proto": {"https"},
			},
			Body:            request.Body,
			IsBase64Encoded: false,
		},
		nil
}

func main() {
	// For API Gateway Payload Format Version 1.0
	// lambda.Start(moesifawslambda.MoesifLogger(HandleLambdaEvent, moesifOptions.MoesifOptions()))

	// For API Gateway Payload Format Version 2.0
	lambda.Start(moesifawslambda.MoesifLogger(HandleLambdaEventV2HTTP, moesifOptionsV2.MoesifOptions()))
}
