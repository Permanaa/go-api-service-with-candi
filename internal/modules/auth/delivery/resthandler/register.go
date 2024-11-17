package resthandler

import (
	"encoding/json"
	"go-api-service-with-candi/internal/modules/auth/domain"
	"io"
	"net/http"

	"github.com/golangid/candi/wrapper"
)

func (h *RestHandler) postRegister(rw http.ResponseWriter, req *http.Request) {
	var payload domain.RegisterRequest

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("auth/register", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload").JSON(rw)
		return
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	resp, err := h.uc.Auth().Register(req.Context(), &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success", resp).JSON(rw)
}
