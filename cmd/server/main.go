package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jeredwong/financial-scheme-manager/internal/handlers"
	"github.com/jeredwong/financial-scheme-manager/internal/middleware"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
	"github.com/jeredwong/financial-scheme-manager/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// connect to db (GORM)
	dsn := "host=localhost user=jered dbname=financial_scheme_manager port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Applicant{}, &models.HouseholdMember{})

	// instantiate repositories, services, handlers 
	applicantRepo := repository.NewGormApplicantRepository(db)
	householdMemberRepo := repository.NewGormHouseholdMemberRepository(db)
	schemeRepo := repository.NewGormSchemeRepository(db)
	schemeCriteriaRepo := repository.NewGormSchemeCriteriaRepository(db)
	benefitRepo := repository.NewGormBenefitRepository(db)
	applicationRepo := repository.NewGormApplicationRepository(db)
	
	applicantService := services.NewApplicantService(applicantRepo)
	householdMemberService := services.NewHouseholdMemberService(householdMemberRepo)
	schemeService := services.NewSchemeService(schemeRepo)
	schemeCriteriaService := services.NewSchemeCriteriaService(schemeCriteriaRepo)
	benefitService := services.NewBenefitService(benefitRepo)
	applicationService := services.NewApplicationService(applicationRepo)

	applicantHandler := handlers.NewApplicantHandler(applicantService, householdMemberService)
	schemeHandler := handlers.NewSchemeHandler(schemeService, schemeCriteriaService, benefitService)
	applicationHandler := handlers.NewApplicationHandler(applicationService, applicantService, schemeService)

	// set up router 
	r := mux.NewRouter()

	// routes 
	r.HandleFunc("/api/health", handlers.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/api/applicants", applicantHandler.GetAllApplicants).Methods("GET")
	r.HandleFunc("/api/applicants", applicantHandler.CreateApplicant).Methods("POST")
	r.HandleFunc("/api/schemes", schemeHandler.GetAllSchemes).Methods("GET")
	// r.HandleFunc("/api/schemes/eligible", schemeHandler.GetEligibleSchemes).Methods("GET")
	r.HandleFunc("/api/applications", applicationHandler.GetAllApplications).Methods("GET")
	r.HandleFunc("/api/applications", applicationHandler.CreateApplication).Methods("POST")

	r.Use(middleware.LoggingMiddleware)

	srv := &http.Server{
		Addr:			":8080",
		Handler:		r,
		ReadTimeout:	15 * time.Second,
		WriteTimeout:	15 * time.Second,
		IdleTimeout:	60 * time.Second,
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
