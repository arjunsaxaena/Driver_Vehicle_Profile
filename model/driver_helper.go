package model

import (
	"time"

	"github.com/google/uuid"
)

type DriverHelper struct {
	ID                             uuid.UUID  `db:"id"`
	UserType                       string     `db:"user_type"`
	FirstName                      string     `db:"first_name"`
	LastName                       string     `db:"last_name"`
	MobileNumber                   string     `db:"mobile_number"`
	AadharNumber                   string     `db:"aadhar_number"`
	LicenseNumber                  string     `db:"license_number"`
	LicenseExpiryDate              *time.Time `db:"license_expiry_date"`
	LicenseDocumentPath            string     `db:"license_document_path"`
	PoliceVerification             string     `db:"police_verification"`
	PoliceVerificationDate         *time.Time `db:"police_verification_date"`
	PoliceVerificationDocumentPath string     `db:"police_verification_document_path"`
	AdditionalDocumentsPath        string     `db:"additional_documents_path"`
	BloodGroup                     string     `db:"blood_group"`
	EmergencyContactName           string     `db:"emergency_contact_name"`
	EmergencyContactNumber         string     `db:"emergency_contact_number"`
	EmergencyContactRelation       string     `db:"emergency_contact_relation"`
	CreatedAt                      time.Time  `db:"created_at"`
	UpdatedAt                      time.Time  `db:"updated_at"`
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
	DriverHelperByLicenseNumber(license string) (DriverHelper, error)
}
