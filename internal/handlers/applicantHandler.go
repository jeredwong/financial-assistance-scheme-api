package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/mapper"
	"github.com/jeredwong/financial-scheme-manager/internal/services"

	// "github.com/google/uuid"
	// "github.com/gorilla/mux"
)

type ApplicantHandler struct {
	applicantService services.ApplicantService
}

func NewApplicantService(applicantService services.ApplicantService) *ApplicantHandler {
	return &ApplicantHandler{applicantService: applicantService}
}

// func (h *ApplicantHandler) GetApplicant(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := uuid.Parse(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid applicant ID", http.StatusBadRequest)
// 		return
// 	}

// 	applicant, err := h.applicantService.GetApplicantById(id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}

// 	writeJSON(w, applicant)
// }

func (h *ApplicantHandler) ListApplicants(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = constants.DefaultPage
	}
	pageSize, err := strconv.Atoi(query.Get("page_size"))
	if err != nil || pageSize < 1  {
		pageSize = constants.DefaultPageSize
	}
	if pageSize > constants.MaxPageSize {
		pageSize = constants.MaxPageSize
	}

	paginationQuery := dto.PaginationQuery{
		Page:     page,
        PageSize: pageSize,
	}

	response, err := h.applicantService.ListApplicants(paginationQuery)
	if err != nil {
		http.Error(w, "Failed to fetch applicants", http.StatusInternalServerError) 
		return
	}

	writeJSON(w, response)
}

func (h *ApplicantHandler) CreateApplicant(w http.ResponseWriter, r *http.Request) {
	var applicantDTO dto.ApplicantDTO
	if err := json.NewDecoder(r.Body).Decode(&applicantDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	// DTO validation 
	if err := validateApplicantDTO(applicantDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	applicant := mapper.ApplicantDTOToModel(applicantDTO)

	err := h.applicantService.CreateApplicant(&applicant)
	if err != nil {
		http.Error(w, "Failed to create applicant", http.StatusInternalServerError)
		return
	}

	// TODO: add household members 

	responseDTO := mapper.ApplicantModelToDTO(applicant)
	
	w.WriteHeader(http.StatusCreated)
	writeJSON(w, responseDTO)
}

// helper functions
// format JSON response
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ") // 2 spaces used for indentation
	if err := encoder.Encode(v); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// validate applicant DTO
func validateApplicantDTO(dto dto.ApplicantDTO) error {
	// TODO: validation logic 
	return nil
}