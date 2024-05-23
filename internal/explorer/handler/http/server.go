package http

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "remy_explorer/docs"
	"remy_explorer/internal/explorer/handler/http/schemas"
)

// NewHTTPServer creates a new HTTP server and registers all endpoints
func NewHTTPServer(endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	// Swagger UI
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	r.Methods("POST").Path("/files").Handler(httptransport.NewServer(
		endpoints.CreateFile,
		decodeCreateFileRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/files/{id}").Handler(httptransport.NewServer(
		endpoints.GetFileByID,
		decodeGetFileByIDRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/files").Handler(httptransport.NewServer(
		endpoints.GetFilesByParentID,
		decodeGetFilesByParentIDRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/files").Handler(httptransport.NewServer(
		endpoints.UpdateFile,
		decodeUpdateFileRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/files/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteFile,
		decodeDeleteFileRequest,
		encodeResponse,
	))

	return r
}

// commonMiddleware adds common headers and applies to all requests
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// encodeResponse encodes the response into JSON format
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// decodeCreateFileRequest decodes the request body into CreateFileRequest
func decodeCreateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.CreateFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeGetFileByIDRequest decodes the request body into GetFileByIDRequest
func decodeGetFileByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.GetFileByIDRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeGetFilesByParentIDRequest decodes the request body into GetFilesByFolderIDRequest
func decodeGetFilesByParentIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.GetFilesByFolderIDRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeUpdateFileRequest decodes the request body into UpdateFileRequest
func decodeUpdateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.UpdateFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeDeleteFileRequest decodes the request body into DeleteFileRequest
func decodeDeleteFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.DeleteFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
