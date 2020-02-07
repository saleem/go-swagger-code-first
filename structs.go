package main

import "net/http"

// User struct.
// swagger:model user
type User struct {
	// ID of user
	//
	// required: true
	ID int64 `json:"id"`
  
	// nickname of user.
	//
	// required: true
	Nickname string `json:"nickname"`
  }

// A RequestStruct model.
//
// This is used for operations that want a user as body of the request
// swagger:parameters createUser
type RequestStruct struct {
// The order to submit.
  //
  // in: body
  // required: true
  User *User `json:"user"`
}

// A ResponseStruct response model.
//
// This is used for returning a response with a single user as body.
//
// swagger:response responseStruct
type ResponseStruct struct {
	// in: body
	Payload *User `json:"user"`
}

// A ValidationError is a swagger response to represent error.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
	  Code    int32  `json:"code"`
	  Message string `json:"message"`
	  Field   string `json:"field"`
	} `json:"body"`
}

// An UserQueryParams model.
//
// swagger:parameters listOrder
type UserQueryParams struct {
	// List user limitations.
	//
	// required: true
	// minimum: 1
	// maximum: 10
	Limit int32 `json:"limit"`
  }

// GetUsers swagger:route GET /users users listUsers
//
// Handler returning list of users.
//
// List of users
//
// Responses:
//        200: []user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: implement code to get users. 
	// Not important for swagger example.
  }

// CreateUser swagger:route POST /users users createUser
//
// Handler to create a user.
//
// Responses:
//        200: responseStruct
//        422: validationError
func CreateUser(w http.ResponseWriter, req *http.Request) {
	// TODO: implement code to create a user. 
	// Not important for swagger example.
}
