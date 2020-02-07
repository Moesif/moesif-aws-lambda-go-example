package moesif_options

import (
	"strings"
	"github.com/moesif/moesifapi-go/models"
	"github.com/aws/aws-lambda-go/events"
)

// Mask Event Model
func maskEventModel(eventModel models.EventModel) models.EventModel {
	return eventModel
}

// Set User Id
func identifyUser(request events.APIGatewayProxyRequest, response events.APIGatewayProxyResponse) string{
	return "golangapiuser"
}

// Set Company Id
func identifyCompany(request events.APIGatewayProxyRequest, response events.APIGatewayProxyResponse) string{
	return "golangapicompany"
}

// Set Session Token
func getSessionToken(request events.APIGatewayProxyRequest, response events.APIGatewayProxyResponse) string{
	return "XXXXXXXXXXXXXXXX"
}

// Skip Event
func shouldSkip(request events.APIGatewayProxyRequest, response events.APIGatewayProxyResponse) bool{
	return strings.Contains(request.Path, "incoming")
}

// Set Metadata
func getMetadata(request events.APIGatewayProxyRequest, response events.APIGatewayProxyResponse) map[string]interface{} {
	
	var innerNestedFields = map[string] interface{} {
		"nestedInner": "test",
	}

	var nestedFields = map[string] interface{} {
		"inner":  innerNestedFields,
	}
	
	var metadata = map[string]interface{} {
		"foo" : "bar",
		"user": "golangapiuser",
		"test": nestedFields,
	}
	return metadata
}

func MoesifOptions() map[string]interface{} {
	var moesifOptions = map[string]interface{} {
		"Api_Version": "1.0.0",
		"Get_Metadata": getMetadata,
		"Should_Skip": shouldSkip,
		"Identify_User": identifyUser,
		"Identify_Company": identifyCompany,
		"Get_Session_Token": getSessionToken,
		"Mask_Event_Model": maskEventModel,
		"Debug": true,
		"Log_Body": true,
	}
	return moesifOptions
}
