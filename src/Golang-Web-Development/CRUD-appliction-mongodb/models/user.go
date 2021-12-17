package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Firstname string		      `json:"firstname" bson:"firstname,omitempty"`
	Lastname  string		      `json:"lastname" bson:"lastname,omitempty"`
	EmailID   string		      `json:"email_id" bson:"email_id,omitempty"`
	ContactNo string		      `json:"contact_no" bson:"contact_no,omitempty"`
	Dob       string              `json:"dob" bson:"dob,omitempty"` // Date of bitrh
}