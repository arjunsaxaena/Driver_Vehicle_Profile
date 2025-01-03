package model

import (
	"time"

	"github.com/google/uuid"
)

type DriverHelper struct {
	ID                             uuid.UUID  `db:"id" json:"id"`
	UserType                       string     `db:"user_type" json:"user_type"`
	FirstName                      string     `db:"first_name" json:"first_name"`
	LastName                       string     `db:"last_name" json:"last_name"`
	MobileNumber                   string     `db:"mobile_number" json:"mobile_number"`
	AadharNumber                   string     `db:"aadhar_number" json:"aadhar_number"`
	LicenseNumber                  string     `db:"license_number" json:"license_number"`
	LicenseExpiryDate              time.Time  `db:"license_expiry_date" json:"license_expiry_date"`
	LicenseDocumentPath            string     `db:"license_document_path" json:"license_document_path"`
	PoliceVerification             string     `db:"police_verification" json:"police_verification"`
	PoliceVerificationDate         *time.Time `db:"police_verification_date" json:"police_verification_date"`
	PoliceVerificationDocumentPath string     `db:"police_verification_document_path" json:"police_verification_document_path"`
	AdditionalDocumentsPath        string     `db:"additional_documents_path" json:"additional_documents_path"`
	BloodGroup                     string     `db:"blood_group" json:"blood_group"`
	EmergencyContactName           string     `db:"emergency_contact_name" json:"emergency_contact_name"`
	EmergencyContactNumber         string     `db:"emergency_contact_number" json:"emergency_contact_number"`
	EmergencyContactRelation       string     `db:"emergency_contact_relation" json:"emergency_contact_relation"`
	CreatedAt                      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt                      time.Time  `db:"updated_at" json:"updated_at"`
}

type Vehicle struct {
	ID                             uuid.UUID `db:"id" json:"id"`
	VehicleNumber                  string    `db:"vehicle_number" json:"vehicle_number"`
	RouteNumber                    string    `db:"route_number" json:"route_number"`
	TotalStudentsCapacity          int       `db:"total_students_capacity" json:"total_students_capacity"`
	SeatsAvailable                 int       `db:"seats_available" json:"seats_available"`
	DriverHelperID                 uuid.UUID `db:"driver_helper_id" json:"driver_helper_id"`
	InsuranceNumber                string    `db:"insurance_number" json:"insurance_number"`
	InsuranceExpiryDate            time.Time `db:"insurance_expiry_date" json:"insurance_expiry_date"`
	PollutionCertificateNumber     string    `db:"pollution_certificate_number" json:"pollution_certificate_number"`
	PollutionCertificateExpiryDate time.Time `db:"pollution_certificate_expiry_date" json:"pollution_certificate_expiry_date"`
	FitnessCertificateNumber       string    `db:"fitness_certificate_number" json:"fitness_certificate_number"`
	FitnessCertificateExpiryDate   time.Time `db:"fitness_certificate_expiry_date" json:"fitness_certificate_expiry_date"`
	VehicleDocumentPath            string    `db:"vehicle_document_path" json:"vehicle_document_path"`
	CreatedAt                      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt                      time.Time `db:"updated_at" json:"updated_at"`
}

type DriverHelperStore interface {
	DriverHelperByID(id uuid.UUID) (DriverHelper, error)
	DriverHelpers() ([]DriverHelper, error)
	Drivers() ([]DriverHelper, error)
	Helpers() ([]DriverHelper, error)
	VerifiedDriverHelpers() ([]DriverHelper, error)
	PendingVerificationDriverHelpers() ([]DriverHelper, error)
	CreateDriverHelper(dh *DriverHelper) error
	UpdateDriverHelper(dh *DriverHelper) error
	DeleteDriverHelper(id uuid.UUID) error
	DriverHelperByMobileNumber(mobile string) (DriverHelper, error)
}

type VehicleStore interface {
	CreateVehicle(v *Vehicle) error
	UpdateVehicle(v *Vehicle) error
	DeleteVehicle(id uuid.UUID) error
	Vehicles() ([]Vehicle, error)
	VehicleByID(id uuid.UUID) (Vehicle, error)
	VehiclesByDriverHelperID(driverHelperID uuid.UUID) ([]Vehicle, error)
	VehiclesByRouteNumber(routeNumber string) ([]Vehicle, error)
	ExpiredCertificatesVehicles() ([]Vehicle, error)
}
