package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/mapper"
	"github.com/jeredwong/financial-scheme-manager/internal/services"

	"github.com/google/uuid"
)

type ApplicantHandler struct {
	applicantService services.ApplicantService
}

func NewApplicantService(applicantService services.ApplicantService) *ApplicantHandler {
	return &ApplicantHandler{applicantService: applicantService}
}

func (h *ApplicantHandler) GetApplicant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid applicant ID", http.StatusBadRequest)
		return
	}

	applicant, err := h.applicantService.GetApplicantById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, applicant)
}

func (h *ApplicantHandler) ListApplicants(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1  {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || page < 1  {
		page = 10
	}

	applicants, err := h.applicantService.ListApplicants(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) 
		return
	}

	writeJSON(w, applicants)
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

	// convert DTO to model
	applicant := mapper.ApplicantDTOToModel(applicantDTO)

	err := h.applicantService.CreateApplicant(&applicant)
	if err != nil {
		http.Error(w, "Failed to create applicant", http.StatusInternalServerError)
		return
	}

	// TODO: add household members 

	// convert model back to DTO
	responseDTO := mapper.ApplicantModelToDTO(applicant)
	
	w.WriteHeader(http.StatusCreated)
	writeJSON(w, responseDTO)
}

// helper functions
// write formatted JSON responsoe
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
	// TODO: 1 validation logic 
	return nil
}