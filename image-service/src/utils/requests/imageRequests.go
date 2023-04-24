package requests

import (
	"backend/src/utils/images"
	"net/http"
	"strconv"
	"strings"
)

func GetImageRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	dirs := strings.Split(path, "/")
	val, err := strconv.Atoi(dirs[1])

	if err != nil {
		makeErrorResponse(w, "can't parse image id", http.StatusBadRequest)
		return
	}

	buffer, err := images.GetImageBytes(val)

	if err != nil {
		makeErrorResponse(w, "can't get image with id", 404)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(buffer)
}
