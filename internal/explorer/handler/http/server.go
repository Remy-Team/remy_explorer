package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"reflect"
	_ "remy_explorer/docs"
	modelerr "remy_explorer/internal/explorer/err"
	"remy_explorer/internal/explorer/handler/http/schemas"
)

// NewHTTPServer initializes and returns a new HTTP server with all routes defined.
func NewHTTPServer(logger log.Logger, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware(logger))

	// Swagger UI
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	// Apply logging middleware to all endpoints
	wrapEndpointsWithLogging(logger, &endpoints)

	// Register file and folder routes
	registerFileRoutes(logger, r, endpoints)
	registerFolderRoutes(logger, r, endpoints)

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

func registerFileRoutes(logger log.Logger, r *mux.Router, endpoints Endpoints) {
	r.Methods("POST").Path("/files").Handler(httptransport.NewServer(
		endpoints.CreateFile,
		decodeCreateFileRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))

	r.Methods("GET").Path("/files/{id}").Handler(httptransport.NewServer(
		endpoints.GetFileByID,
		decodeGetFileByIDRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))

	r.Methods("GET").Path("/files").Handler(httptransport.NewServer(
		endpoints.GetFilesByParentID,
		decodeGetFilesByParentIDRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))

	r.Methods("PUT").Path("/files").Handler(httptransport.NewServer(
		endpoints.UpdateFile,
		decodeUpdateFileRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))

	r.Methods("DELETE").Path("/files/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteFile,
		decodeDeleteFileRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))
}

func registerFolderRoutes(logger log.Logger, r *mux.Router, endpoints Endpoints) {
	r.Methods("POST").Path("/folders").Handler(httptransport.NewServer(
		endpoints.CreateFolder,
		decodeCreateFolderRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))

	r.Methods("GET").Path("/folders/{id}").Handler(httptransport.NewServer(
		endpoints.GetFolderByID,
		decodeGetFolderByIDRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))

	r.Methods("GET").Path("/folders/{parent_id}/subfolders").Handler(httptransport.NewServer(
		endpoints.GetFoldersByParentID,
		decodeGetFoldersByParentIDRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))

	r.Methods("PUT").Path("/folders").Handler(httptransport.NewServer(
		endpoints.UpdateFolder,
		decodeUpdateFolderRequest,
		encodeResponse(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse(logger)),
	))

	r.Methods("DELETE").Path("/folders/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteFolder,
		decodeDeleteFolderRequest,
		encodeResponse(logger),
	))
}

// commonMiddleware adds common HTTP headers to all responses.
func commonMiddleware(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			level.Info(logger).Log(
				"msg", "received request",
				"method", r.Method,
				"url", r.URL.String(),
			)
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
			level.Info(logger).Log(
				"msg", "handled request",
				"method", r.Method,
				"url", r.URL.String(),
			)
		})
	}
}

// encodeResponse encodes the response into JSON format and handles errors.
func encodeResponse(logger log.Logger) httptransport.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		level.Info(logger).Log("msg", "Encode", "response", response)
		if err, ok := response.(error); ok {
			level.Error(logger).Log("msg", "Error received in encoder", "error", err.Error())

			var notFoundErr *modelerr.NotFound
			if errors.As(err, &notFoundErr) {
				http.Error(w, err.Error(), http.StatusNotFound)
				level.Info(logger).Log("msg", "resource not found", "err", err)
				return nil
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			level.Error(logger).Log("msg", "error processing request", "err", err)
			return nil
		}
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(response)
	}
}

func encodeErrorResponse(logger log.Logger) httptransport.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		level.Error(logger).Log("err", err, "msg", "handling error")

		var notFoundErr *modelerr.NotFound
		if errors.As(err, &notFoundErr) {
			http.Error(w, err.Error(), http.StatusNotFound)
			level.Info(logger).Log("msg", "resource not found", "err", err)
			return
		}

		// Обработка других ошибок
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		level.Error(logger).Log("msg", "internal server error", "err", err)
	}
}

func decodeCreateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.CreateFileRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeUpdateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.UpdateFileRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
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

func decodeCreateFolderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.CreateFolderRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetFolderByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}
	return schemas.GetFolderByIDRequest{ID: id}, nil
}

func decodeGetFoldersByParentIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	parentID, ok := vars["parentID"]
	if !ok {
		return nil, errors.New("parentID is missing in parameters")
	}
	return schemas.GetFoldersByParentIDRequest{ParentID: parentID}, nil
}

func decodeUpdateFolderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req schemas.UpdateFolderRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteFolderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}
	return schemas.DeleteFolderRequest{ID: id}, nil
}
