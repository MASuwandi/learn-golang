package middleware

import (
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"net/http"
)

// Struct with Kontrak Handler
// Middleware harus dalam bentuk handler
type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

// Base on Security Schema/Category Auth/name: "X-API-Key" in apispec.
func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	// Check "X-API-Key" with value "SECRET"
	// if "SECRET" == request.Header.Get("X-API-Key") // Yoda Condition
	if request.Header.Get("X-API-Key") == "SECRET" {
		// approve
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// reject
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		
		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}

}