package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"net/http"
	"reflect"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "remy_explorer/docs"
	"remy_explorer/internal/explorer/handler/http/schemas"
)

// NewHTTPServer initializes and returns a new HTTP server with all routes defined.
func NewHTTPServer(logger log.Logger, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	// Swagger UI
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	// Apply logging middleware to all endpoints
	wrapEndpointsWithLogging(logger, &endpoints)

	// Register file and folder routes
	registerFileRoutes(r, endpoints)
	registerFolderRoutes(r, endpoints)

	return r
}

// wrapEndpointsWithLogging applies logging middleware to all fields in the Endpoints struct using reflection.
func wrapEndpointsWithLogging(logger log.Logger, endpoints interface{}) {
	loggingMiddleware := makeLoggingMiddleware(logger)
	endpointsVal := reflect.ValueOf(endpoints).Elem()

	for i := 0; i < endpointsVal.NumField(); i++ {
		field := endpointsVal.Field(i)
		if field.CanInterface() {
			ep, ok := field.Interface().(endpoint.Endpoint)
			if ok {
				wrapped := loggingMiddleware(ep)
				if field.CanSet() {
					field.Set(reflect.ValueOf(wrapped))
				}
			}
		}
	}
}

func makeLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			level.Info(logger).Log("msg", "calling endpoint", "request", request)
			response, err = next(ctx, request)
			level.Info(logger).Log("msg", "called endpoint", "response", response, "err", err)
			return
		}
	}
}

func registerFileRoutes(r *mux.Router, endpoints Endpoints) {
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
}

func registerFolderRoutes(r *mux.Router, endpoints Endpoints) {
	r.Methods("POST").Path("/folders").Handler(httptransport.NewServer(
		endpoints.CreateFolder,
		decodeRequest(new(schemas.CreateFolderRequest)),
		encodeResponse,
	))

	r.Methods("GET").Path("/folders/{id}").Handler(httptransport.NewServer(
		endpoints.GetFolderByID,
		decodeRequest(new(schemas.GetFolderByIDRequest)),
		encodeResponse,
	))

	r.Methods("GET").Path("/folders").Handler(httptransport.NewServer(
		endpoints.GetFoldersByParentID,
		decodeRequest(new(schemas.GetFoldersByParentIDRequest)),
		encodeResponse,
	))

	r.Methods("PUT").Path("/folders").Handler(httptransport.NewServer(
		endpoints.UpdateFolder,
		decodeRequest(new(schemas.UpdateFolderRequest)),
		encodeResponse,
	))

	r.Methods("DELETE").Path("/folders/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteFolder,
		decodeRequest(new(schemas.DeleteFolderRequest)),
		encodeResponse,
	))
}

// commonMiddleware adds common HTTP headers to all responses.
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Log("msg", "calling endpoint", "request", request)
			response, err = next(ctx, request)
			logger.Log("msg", "called endpoint", "response", response, "err", err)
			return
		}
	}
}

// encodeResponse encodes the response into JSON format and handles errors.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.CreateFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.UpdateFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetFileByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.GetFileByIDRequest
	vars := mux.Vars(r)
	if id, ok := vars["id"]; ok {
		req.ID = id
	} else {
		return nil, errors.New("id is missing in parameters")
	}
	return req, nil
}

func decodeGetFilesByParentIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	parentID, ok := vars["parentID"]
	if !ok {
		return nil, errors.New("parentID is missing in parameters")
	}
	return schemas.GetFilesByFolderIDRequest{FolderID: parentID}, nil
}

func decodeDeleteFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}
	return schemas.DeleteFileRequest{ID: id}, nil
}
