package httpreq

import (
	"net/http"
	"net/url"
)

// IsOK check response status code is 200
func IsOK(statusCode int) bool {
	return statusCode == http.StatusOK
}

// IsSuccessful check response status code is in 200 - 300
func IsSuccessful(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < 300
}

// IsRedirect check response status code is in [301, 302, 303, 307]
func IsRedirect(statusCode int) bool {
	return statusCode == http.StatusMovedPermanently ||
		statusCode == http.StatusFound ||
		statusCode == http.StatusSeeOther ||
		statusCode == http.StatusTemporaryRedirect
}

// IsForbidden is this response forbidden(403)
func IsForbidden(statusCode int) bool {
	return statusCode == http.StatusForbidden
}

// IsNotFound is this response not found(404)
func IsNotFound(statusCode int) bool {
	return statusCode == http.StatusNotFound
}

// IsClientError check response is client error (400 - 500)
func IsClientError(statusCode int) bool {
	return statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError
}

// IsServerError check response is server error (500 - 600)
func IsServerError(statusCode int) bool {
	return statusCode >= http.StatusInternalServerError && statusCode <= 600
}

// AddHeadersToRequest adds the key, value pairs from the given http.Header to the
// request. Values for existing keys are appended to the keys values.
func AddHeadersToRequest(req *http.Request, header http.Header) {
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}

func AddQueriesToURL() {

}

func ToQueryValues(data interface{}) url.Values {
	// use url.Values directly if we have it
	if uv, ok := data.(url.Values); ok {
		return uv
	}

	uv := make(url.Values)
	if strMp, ok := data.(map[string]string); ok {
		for k, v := range strMp {
			uv.Add(k, v)
		}
	}

	return uv
}
