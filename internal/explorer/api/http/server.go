package http

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"remy_explorer/internal/explorer/api/http/schemas"
)

// NewHTTPServer returns a new HTTP server
func NewHTTPServer(endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

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

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.CreateFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeGetFileByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.GetFileByIDRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeGetFilesByParentIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.GetFilesByFolderIDRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeUpdateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.UpdateFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeDeleteFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.DeleteFileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
