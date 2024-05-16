package http

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"remy_explorer/internal/explorer/handler/http/schemas"
)

//	@title			Remy Explorer API
//	@version		1.0
//	@description	This is the API documentation for Remy Explorer.
//	@termsOfService	http://remy_explorer.com/terms/

//	@contact.name	API Support
//	@contact.url	http://www.remy_explorer.com/support
//	@contact.email	support@remy_explorer.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// NewHTTPServer @host localhost:8080
//
//	@BasePath	/api/v1
//
// NewHTTPServer creates a new HTTP server and registers all endpoints
func NewHTTPServer(endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	// Swagger UI
	r.Methods("GET").Path("/swagger").Handler(httpSwagger.WrapHandler)

	// Redirect to Swagger UI
	r.Methods("GET").Path("/").Handler(http.RedirectHandler("/swagger/index.html", http.StatusTemporaryRedirect))

	//	@Summary		Create a new file
	//	@Description	Create a new file in the system
	//	@Tags			files
	//	@Accept			json
	//	@Produce		json
	//	@Param			body	body		schemas.CreateFileRequest	true	"Create File Request"
	//	@Success		200		{object}	schemas.CreateFileResponse
	//	@Failure		400		{object} 	schemas.ErrorResponse
	//	@Failure		500		{object} 	schemas.ErrorResponse
	//	@Router			/files [post]
	r.Methods("POST").Path("/files").Handler(httptransport.NewServer(
		endpoints.CreateFile,
		decodeCreateFileRequest,
		encodeResponse,
	))

	//	@Summary		Get file by ID
	//	@Description	Retrieve a file's details by its ID
	//	@Tags			files
	//	@Accept			json
	//	@Produce		json
	//	@Param			id	path		string	true	"File ID"
	//	@Success		200	{object}	schemas.GetFileByIDResponse
	//	@Failure		404	{object} 	schemas.ErrorResponse
	//	@Failure		500	{object} 	schemas.ErrorResponse
	//	@Router			/files/{id} [get]
	r.Methods("GET").Path("/files/{id}").Handler(httptransport.NewServer(
		endpoints.GetFileByID,
		decodeGetFileByIDRequest,
		encodeResponse,
	))

	//	@Summary		Get files by folder ID
	//	@Description	Retrieve a list of files within a specific folder
	//	@Tags			files
	//	@Accept			json
	//	@Produce		json
	//	@Param			folderID	query		string	true	"Folder ID"
	//	@Success		200			{array}		schemas.GetFilesByFolderIDResponse
	//	@Failure		404			{object} 	schemas.ErrorResponse
	//	@Failure		500			{object} 	schemas.ErrorResponse
	//	@Router			/files [get]
	r.Methods("GET").Path("/files").Handler(httptransport.NewServer(
		endpoints.GetFilesByParentID,
		decodeGetFilesByParentIDRequest,
		encodeResponse,
	))

	//	@Summary		Update a file
	//	@Description	Update the details of an existing file
	//	@Tags			files
	//	@Accept			json
	//	@Produce		json
	//	@Param			body	body		schemas.UpdateFileRequest	true	"Update File Request"
	//	@Success		200		{object}	schemas.UpdateFileResponse
	//	@Failure		400		{object} 	schemas.ErrorResponse
	//	@Failure		404		{object} 	schemas.ErrorResponse
	//	@Failure		500		{object} 	schemas.ErrorResponse
	//	@Router			/files [put]
	r.Methods("PUT").Path("/files").Handler(httptransport.NewServer(
		endpoints.UpdateFile,
		decodeUpdateFileRequest,
		encodeResponse,
	))

	//	@Summary		Delete a file
	//	@Description	Delete a file by its ID
	//	@Tags			files
	//	@Accept			json
	//	@Produce		json
	//	@Param			id	path		string	true	"File ID"
	//	@Success		200	{object}	schemas.DeleteFileResponse
	//	@Failure		404	{object} 	schemas.ErrorResponse
	//	@Failure		500	{object} 	schemas.ErrorResponse
	//	@Router			/files/{id} [delete]
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
