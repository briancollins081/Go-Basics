package handlers

import (
	"net/http"
	"rest-api/user"
)

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()

	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}
