package controller

const (
	// API represents the group of API.
	API = "/api"
	// APICategories represents the group of category management API.
	APICategories = API + "/categories"
)

const (
	// APIUser represents the group of auth management API.
	APIUser = API + "/auth"
	// APIUserLoginStatus represents the API to get the status of logged in user.
	APIUserLoginStatus = APIUser + "/loginStatus"
	// APIUserLoginUser represents the API to get the logged in user.
	APIUserLoginUser = APIUser + "/loginUser"
	// APIUserLogin represents the API to login by session authentication.
	APIUserLogin = APIUser + "/login"
	// APIUserLogout represents the API to logout.
	APIUserLogout = APIUser + "/logout"
)

const (
	// APIHealth represents the API to get the status of this application.
	APIHealth = API + "/health"
)
