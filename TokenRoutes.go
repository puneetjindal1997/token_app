package main

import "token_app/controllers"

// Routes for Admin management
var AdminRoutes = Routes{
	//routes Admin
	Route{"AdminLogin", "POST", "/login", controllers.AdminLogin},
	// routes account
	Route{"Login", "POST", "/user-login", controllers.Login},
}

var UserLoginRoutes = Routes{
	// generate token auth
	Route{"CreateToken", "POST", "/create-token", controllers.CreateToken},
}
