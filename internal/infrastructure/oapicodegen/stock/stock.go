// Package stock provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package stock

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse struct {
	Message string `json:"message"`
}

// NewStockItem defines model for NewStockItem.
type NewStockItem = StockItemName

// StockItemId defines model for StockItemId.
type StockItemId struct {
	Id openapi_types.UUID `json:"id" validate:"required"`
}

// StockItemName defines model for StockItemName.
type StockItemName struct {
	Name string `json:"name" validate:"required,lt=100"`
}

// BadRequest defines model for BadRequest.
type BadRequest = BadRequestResponse

// Created defines model for Created.
type Created = StockItemId

// PostStockItemJSONRequestBody defines body for PostStockItem for application/json ContentType.
type PostStockItemJSONRequestBody = NewStockItem

// PutStockItemJSONRequestBody defines body for PutStockItem for application/json ContentType.
type PutStockItemJSONRequestBody = NewStockItem

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create Stock Item
	// (POST /stock/items)
	PostStockItem(ctx echo.Context) error
	// Delete Stock Item
	// (DELETE /stock/items/{stockitemId})
	DeleteStockItem(ctx echo.Context, stockitemId openapi_types.UUID) error
	// Update Stock Item
	// (PUT /stock/items/{stockitemId})
	PutStockItem(ctx echo.Context, stockitemId openapi_types.UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) PostStockItem(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostStockItem(ctx)
	return err
}

// DeleteStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteStockItem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "stockitemId" -------------
	var stockitemId openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "stockitemId", runtime.ParamLocationPath, ctx.Param("stockitemId"), &stockitemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter stockitemId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteStockItem(ctx, stockitemId)
	return err
}

// PutStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) PutStockItem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "stockitemId" -------------
	var stockitemId openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "stockitemId", runtime.ParamLocationPath, ctx.Param("stockitemId"), &stockitemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter stockitemId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutStockItem(ctx, stockitemId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/stock/items", wrapper.PostStockItem)
	router.DELETE(baseURL+"/stock/items/:stockitemId", wrapper.DeleteStockItem)
	router.PUT(baseURL+"/stock/items/:stockitemId", wrapper.PutStockItem)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xUTW/bMAz9Kwa3oxI7a3cxsMO6DyAo0BYtdipyEGwm1WZ9VKKzBIH++yA5cZw4a1Is",
	"QHezKIp8fo+PKyi0NFqhIgf5Ciw6o5XDeLji5T0+1+gonAqtCFX85MZUouAktEp/Oq1CzBVPKHn4em9x",
	"Cjm8S7el0+bWpduS9+tO4L1nUKIrrDChIuShcbLp7Bl8scgJy7OBeCBd/BoTynF5qPumnWcwVoRW8eoB",
	"7RztN2u1DcV38zdJSZOVNGmewY2m77pWZf/JjaakufIMbq/7CbfXEICtEe+K0TKXr8BYbdCSaAST6Byf",
	"xQtaGoQcHFmhZrGWxeda2EDjY5s4CSjxd0tI5LWqbqeQP57I4Q2XCD4U6tLagyZibKqt5AQ51LUoge2h",
	"ZLAYaG7EoNAlzlANcEGWD4jPYok5r0TJKTxo/8Xv/5koYQdLhNdDo9ZRyRdC1hLyUZb9OxpW0adRlvVB",
	"xXaTGBZqqqM+gqrwNgJNAtLk890YGMzRumYGRsNsmIUB0QYVNwJyuBhmwwtgYDg9RRSpC+9TQSjj2ejG",
	"rIcmOtn2gljURvMEteBOO9pOQYMeHV3pcnk22+0M2h5HZGuMgc76+ZCN/layzUs7Zr3MsuP5nZ3mGXw8",
	"5cmhJRDNWUvJ7fIgv+G+q026igcR3eEbhSok7Gv1NcZf0qrJ6KpluOUSCa2LvhWhThgRYOtJh0532Oed",
	"dfR72aDeT3oanUBgWGZBnsvjqe3KPJ84fUI9A1MfsMkPUx6zSU1vyPv/YMlXyP1qN77VhPR19977PwEA",
	"AP//a3TP8xsJAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}