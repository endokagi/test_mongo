package models

type User struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
}

type Response struct {
	Constants   Constants `json:"constants"`
	Information User      `json:"information"`
}

func MessageResponse(message Constants, user User) Response {

	var response Response
	response.Constants = message
	response.Information = user

	return response
}
