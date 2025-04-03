package moesif_options_v2

import (
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/moesif/moesifapi-go/models"
)

// Mask Event Model
func maskEventModel(eventModel models.EventModel) models.EventModel {
	return eventModel
}

// Set User Id
func identifyUser(request events.APIGatewayV2HTTPRequest, response events.APIGatewayV2HTTPResponse) string {
	return "golangapiuserV2"
}

// Set Company Id
func identifyCompany(request events.APIGatewayV2HTTPRequest, response events.APIGatewayV2HTTPResponse) string {
	return "golangapicompanyV2"
}

// Set Session Token
func getSessionToken(request events.APIGatewayV2HTTPRequest, response events.APIGatewayV2HTTPResponse) string {
	return "XXXXXXXXXXXXXXXX"
}

// Skip Event
func shouldSkip(request events.APIGatewayV2HTTPRequest, response events.APIGatewayV2HTTPResponse) bool {
	return strings.Contains(request.RawPath, "incoming")
}

// Set Metadata
func getMetadata(request events.APIGatewayV2HTTPRequest, response events.APIGatewayV2HTTPResponse) map[string]interface{} {

	var innerNestedFields = map[string]interface{}{
		"nestedInner": "test",
	}

	var nestedFields = map[string]interface{}{
		"inner": innerNestedFields,
	}

	var metadata = map[string]interface{}{
		"foo":  "bar",
		"user": "golangapiuserV2",
		"test": nestedFields,
	}
	return metadata
}

// Skip Outgoing Event
func shouldSkipOutgoing(request *http.Request, response *http.Response) bool {
	return strings.Contains(request.URL.String(), "outgoing")
}

// Set Outgoing Event User Id
func identifyUserOutgoing(request *http.Request, response *http.Response) string {
	return "golangapiuserOutgoingV2"
}

// Set Outgoing Event Company Id
func identifyCompanyOutgoing(request *http.Request, response *http.Response) string {
	return "golangapicompanyOutgoingV2"
}

// Set Outgoing Event Session Token
func getSessionTokenOutgoing(request *http.Request, response *http.Response) string {
	return "token is blah blah blah"
}

// Mask Outgoing Event Model
func maskEventModelOutgoing(eventModel models.EventModel) models.EventModel {
	return eventModel
}

// Set Outoing Event Metadata
func getMetadataOutgoing(request *http.Request, response *http.Response) map[string]interface{} {

	var innerNestedFields = map[string]interface{}{
		"nestedInner": "test",
	}

	var nestedFields = map[string]interface{}{
		"inner": innerNestedFields,
	}

	var metadata = map[string]interface{}{
		"foo":  "bar",
		"user": "golangapiuserV2",
		"test": nestedFields,
	}
	return metadata
}

func MoesifOptions() map[string]interface{} {
	var moesifOptions = map[string]interface{}{
		"Api_Version":                "1.0.0",
		"Get_Metadata":               getMetadata,
		"Should_Skip":                shouldSkip,
		"Identify_User":              identifyUser,
		"Identify_Company":           identifyCompany,
		"Get_Session_Token":          getSessionToken,
		"Mask_Event_Model":           maskEventModel,
		"Debug":                      true,
		"Log_Body":                   true,
		"Log_Body_Outgoing":          true,
		"Should_Skip_Outgoing":       shouldSkipOutgoing,
		"Identify_User_Outgoing":     identifyUserOutgoing,
		"Identify_Company_Outgoing":  identifyCompanyOutgoing,
		"Get_Metadata_Outgoing":      getMetadataOutgoing,
		"Get_Session_Token_Outgoing": getSessionTokenOutgoing,
		"Mask_Event_Model_Outgoing":  maskEventModelOutgoing,
	}
	return moesifOptions
}
