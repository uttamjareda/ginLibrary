package db

import "ginLibrary/models"

// Users is our mock database of users
var Users = []models.User{
	{Username: "user1", UserType: "regular", Password: "user123"},
	{Username: "user2", UserType: "regular", Password: "user123"},
	{Username: "user3", UserType: "regular", Password: "user123"},
	{Username: "user4", UserType: "regular", Password: "user123"},
	{Username: "user5", UserType: "regular", Password: "user123"},
	{Username: "user6", UserType: "regular", Password: "user123"},
	{Username: "admin", UserType: "admin", Password: "admin123"},
}
