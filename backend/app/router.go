package app

import "user_auth_golang/backend/controller/users"

func mapUrls() {
	router.POST("/api/register", users.Register)
}
