CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE user_type_enum AS ENUM ('Driver', 'Helper');
CREATE TYPE police_verification_enum AS ENUM ('Yes', 'No');
CREATE TYPE blood_group_enum AS ENUM ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-');

CREATE TABLE driver_helpers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Personal Details
    user_type user_type_enum NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    mobile_number VARCHAR(15) NOT NULL,
    aadhar_number VARCHAR(12),

    -- Document Details
    license_number VARCHAR(20) NOT NULL,
    license_expiry_date DATE NOT NULL,
    license_document_path VARCHAR(255),
    police_verification police_verification_enum DEFAULT 'No',
    police_verification_date DATE,
    police_verification_document_path VARCHAR(255),
    additional_documents_path VARCHAR(255),

    -- Other Information
    blood_group blood_group_enum,

    -- Emergency Contact
    emergency_contact_name VARCHAR(50),
    emergency_contact_number VARCHAR(15),
    emergency_contact_relation VARCHAR(50),

    -- Timestamps
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
