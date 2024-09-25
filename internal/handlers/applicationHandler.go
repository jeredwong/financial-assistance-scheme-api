package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/mapper"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/services"
)

type ApplicationHandler struct {
	applicationService services.ApplicationService
	applicantService services.ApplicantService
	schemeService services.SchemeService
}

func NewApplicationHandler(
	applicationService services.ApplicationService, 
	applicantService services.ApplicantService,
	schemeService services.SchemeService) *ApplicationHandler {
	return &ApplicationHandler{
		applicationService: applicationService, 
		applicantService: applicantService,
		schemeService: schemeService,
	}

}

func (h *ApplicationHandler) GetAllApplications(w http.ResponseWriter, r *http.Request) {
	// TODO: convert pagination logic to utility function (DRY)
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

	response, err := h.applicationService.GetAllApplications(paginationQuery)
	if err != nil {
		http.Error(w, "Failed to fetch applications", http.StatusInternalServerError)
		return
	}

	applications, ok := response.Data.([]models.Application)
	if !ok {
		http.Error(w, "Unexpected data format", http.StatusInternalServerError)
		return
	}

	applicationDTOs := make([]dto.ApplicationDTO, len(applications))
	for i, application := range(applications) {
		applicationDTO := mapper.ApplicationModelToDTO(application)

		applicant, err := h.applicantService.GetApplicantById(application.ApplicantId)
		if err != nil {
			http.Error(w, "Failed to fetch applicant", http.StatusInternalServerError)
			return
		}
		applicantDTO := mapper.ApplicantModelToDTO(applicant)
		applicationDTO.Applicant = applicantDTO

		scheme, err := h.schemeService.GetSchemeById(application.SchemeId)
		if err != nil {
			http.Error(w, "Failed to fetch scheme", http.StatusInternalServerError)
			return
		}
		schemeDTO := mapper.SchemeModelToDTO(scheme)
		applicationDTO.Scheme = schemeDTO

		applicationDTOs[i] = applicationDTO
	}
	response.Data = applicationDTOs

	writeJSON(w, response)
}

func (h *ApplicationHandler) CreateApplication(w http.ResponseWriter, r *http.Request) {
	var applicationDTO dto.ApplicationDTO
	if err := json.NewDecoder(r.Body).Decode(&applicationDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	// DTO validation 
	if err := validateApplicationDTO(applicationDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// db transaction 
	application := mapper.ApplicationDTOToModel(applicationDTO)

	err := h.applicationService.CreateApplication(&application)
	if err != nil {
		http.Error(w, "Failed to create application", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	writeJSON(w, applicationDTO)
}

// TODO: move helper functions into utility directory
// helper functions 

// validate application DTO
func validateApplicationDTO(dto dto.ApplicationDTO) error {
	// TODO: validation logic
	return nil
}