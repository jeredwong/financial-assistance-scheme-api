package handlers

import (
	"net/http"
	"strconv"

	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/mapper"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/services"
)

type SchemeHandler struct {
	schemeService services.SchemeService
	schemeCriteriaService services.SchemeCriteriaService
	benefitService services.BenefitService
}

func NewSchemeHandler(
	schemeService services.SchemeService,
	schemeCriteriaService services.SchemeCriteriaService,
	benefitService services.BenefitService) *SchemeHandler {
	return &SchemeHandler{
		schemeService: schemeService, 
		schemeCriteriaService: schemeCriteriaService,
		benefitService: benefitService,
	}
}

func (h *SchemeHandler) GetAllSchemes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = constants.DefaultPage
	}
	pageSize, err := strconv.Atoi(query.Get("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = constants.DefaultPageSize
	}
	if pageSize > constants.MaxPageSize {
		pageSize = constants.MaxPageSize
	}

	paginationQuery := dto.PaginationQuery{
		Page: page,
		PageSize: pageSize,
	}

	response, err := h.schemeService.GetAllSchemes(paginationQuery)

	if err != nil {
		http.Error(w, "Failed to fetch schemes", http.StatusInternalServerError)
		return
	}

	schemes, ok := response.Data.([]models.Scheme)
	if !ok {
		http.Error(w, "Unexpected data format", http.StatusInternalServerError)
		return
	}

	schemeDTOs := make([]dto.SchemeDTO, len(schemes))
	for i, scheme := range(schemes) {
		schemeDTO := mapper.SchemeModelToDTO(scheme)

		criteria, err := h.schemeCriteriaService.GetSchemeCriteriaBySchemeId(scheme.ID)
		if err != nil {
			http.Error(w, "Failed to fetch scheme criteria", http.StatusInternalServerError)
			return
		}

		criteriaDTOs := mapper.SchemeCriteriaModelsToDTOs(criteria)
		schemeDTO.Criteria = criteriaDTOs

		for j, criteriaDTO := range(schemeDTO.Criteria) {

			benefits, err := h.benefitService.GetBenefitsByCriteriaId(criteriaDTO.Id)
			if err != nil {
				http.Error(w, "Failed to fetch scheme benefits", http.StatusInternalServerError)
				return
			}
			criteriaDTO.Benefits = mapper.BenefitModelsToDTOs(benefits)

			criteriaDTOs[j] = criteriaDTO

		}

		schemeDTOs[i] = schemeDTO
	}
	response.Data = schemeDTOs

	w.WriteHeader(http.StatusCreated)
	writeJSON(w, response)
}

// func (h *SchemeHandler) GetEligibleSchemes(w http.ResponseWriter, r *http.Request) {
// 	query := 
// }