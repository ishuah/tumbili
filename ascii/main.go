package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/common-nighthawk/go-figure"
)

type AsciiRequest struct {
	Text string `json:"text"`
}

type AsciiResponse struct {
	Text   string `json:"text"`
	Art    string `json:"art"`
	Length int    `json:"length"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var asciiRequest AsciiRequest
	err := json.Unmarshal([]byte(request.Body), &asciiRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
	}
	myFigure := figure.NewFigure(asciiRequest.Text, "big", true)
	asciiResponse := AsciiResponse{
		Text:   asciiRequest.Text,
		Art:    myFigure.String(),
		Length: len(myFigure.String()),
	}
	response, err := json.Marshal(asciiResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
