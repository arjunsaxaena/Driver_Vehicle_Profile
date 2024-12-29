package controllers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"

	"github.com/arjunsaxaena/driver_vehicle_profile/model"
)

type DBDriverHelperStore struct {
	db *sqlx.DB
}

func NewDBDriverHelperStore(db *sqlx.DB) *DBDriverHelperStore {
	return &DBDriverHelperStore{db: db}
}

func (s *DBDriverHelperStore) DriverHelperByID(id uuid.UUID) (model.DriverHelper, error) {
	var dh model.DriverHelper
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("driver_helpers").Where(sb.Equal("id", id))

	query, args := sb.Build()
	err := s.db.Get(&dh, query, args...)
	return dh, err
}

func (s *DBDriverHelperStore) DriverHelpers() ([]model.DriverHelper, error) {
	var dhs []model.DriverHelper
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("driver_helpers")

	query, args := sb.Build()
	err := s.db.Select(&dhs, query, args...)
	return dhs, err
}

func (s *DBDriverHelperStore) Drivers() ([]model.DriverHelper, error) {
	var drivers []model.DriverHelper
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("driver_helpers").Where(sb.Equal("user_type", "Driver"))

	query, args := sb.Build()
	err := s.db.Select(&drivers, query, args...)
	return drivers, err
}

func (s *DBDriverHelperStore) Helpers() ([]model.DriverHelper, error) {
	var helpers []model.DriverHelper
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("driver_helpers").Where(sb.Equal("user_type", "Helper"))

	query, args := sb.Build()
	err := s.db.Select(&helpers, query, args...)
	return helpers, err
}

func (s *DBDriverHelperStore) VerifiedDriverHelpers() ([]model.DriverHelper, error) {
	var dhs []model.DriverHelper
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("driver_helpers").Where(sb.Equal("police_verification", "Yes"))

	query, args := sb.Build()
	err := s.db.Select(&dhs, query, args...)
	return dhs, err
}

func (s *DBDriverHelperStore) PendingVerificationDriverHelpers() ([]model.DriverHelper, error) {
	var dhs []model.DriverHelper
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("driver_helpers").Where(sb.Equal("police_verification", "No"))

	query, args := sb.Build()
	err := s.db.Select(&dhs, query, args...)
	return dhs, err
}

func (s *DBDriverHelperStore) CreateDriverHelper(dh *model.DriverHelper) error {
	if dh.ID == uuid.Nil {
		dh.ID = uuid.New()
	}

	sb := sqlbuilder.NewInsertBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.InsertInto("driver_helpers").
		Cols("id", "user_type", "first_name", "last_name", "mobile_number", "aadhar_number",
			"license_number", "license_expiry_date", "license_document_path", "police_verification",
			"police_verification_date", "police_verification_document_path", "additional_documents_path",
			"blood_group", "emergency_contact_name", "emergency_contact_number", "emergency_contact_relation").
		Values(dh.ID, dh.UserType, dh.FirstName, dh.LastName, dh.MobileNumber, dh.AadharNumber,
			dh.LicenseNumber, dh.LicenseExpiryDate, dh.LicenseDocumentPath, dh.PoliceVerification,
			dh.PoliceVerificationDate, dh.PoliceVerificationDocumentPath, dh.AdditionalDocumentsPath,
			dh.BloodGroup, dh.EmergencyContactName, dh.EmergencyContactNumber, dh.EmergencyContactRelation)

	query, args := sb.Build()
	_, err := s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert driver/helper: %w", err)
	}

	return nil
}

func (s *DBDriverHelperStore) UpdateDriverHelper(dh *model.DriverHelper) error {
	sb := sqlbuilder.NewUpdateBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Update("driver_helpers").Set(
		sb.Assign("user_type", dh.UserType),
		sb.Assign("first_name", dh.FirstName),
		sb.Assign("last_name", dh.LastName),
		sb.Assign("mobile_number", dh.MobileNumber),
		sb.Assign("aadhar_number", dh.AadharNumber),
		sb.Assign("license_number", dh.LicenseNumber),
		sb.Assign("license_expiry_date", dh.LicenseExpiryDate),
		sb.Assign("license_document_path", dh.LicenseDocumentPath),
		sb.Assign("police_verification", dh.PoliceVerification),
		sb.Assign("police_verification_date", dh.PoliceVerificationDate),
		sb.Assign("police_verification_document_path", dh.PoliceVerificationDocumentPath),
		sb.Assign("additional_documents_path", dh.AdditionalDocumentsPath),
		sb.Assign("blood_group", dh.BloodGroup),
		sb.Assign("emergency_contact_name", dh.EmergencyContactName),
		sb.Assign("emergency_contact_number", dh.EmergencyContactNumber),
		sb.Assign("emergency_contact_relation", dh.EmergencyContactRelation),
		sb.Assign("updated_at", time.Now()),
	).Where(sb.Equal("id", dh.ID))

	query, args := sb.Build()
	_, err := s.db.Exec(query, args...)
	return err
}

func (s *DBDriverHelperStore) DeleteDriverHelper(id uuid.UUID) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.DeleteFrom("driver_helpers").Where(sb.Equal("id", id))

	query, args := sb.Build()
	_, err := s.db.Exec(query, args...)
	return err
}

func (s *DBDriverHelperStore) DriverHelperByMobileNumber(mobile string) (model.DriverHelper, error) {
	var dh model.DriverHelper
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("driver_helpers").Where(sb.Equal("mobile_number", mobile))

	query, args := sb.Build()
	err := s.db.Get(&dh, query, args...)
	return dh, err
}

func (s *DBDriverHelperStore) DriverHelperByLicenseNumber(license string) (model.DriverHelper, error) {
	var dh model.DriverHelper
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("driver_helpers").Where(sb.Equal("license_number", license))

	query, args := sb.Build()
	err := s.db.Get(&dh, query, args...)
	return dh, err
}
