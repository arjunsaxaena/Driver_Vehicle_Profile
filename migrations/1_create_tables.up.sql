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

CREATE TABLE vehicles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Vehicle Details
    vehicle_number VARCHAR(20) NOT NULL UNIQUE,
    route_number VARCHAR(20),
    total_students_capacity INT NOT NULL CHECK (total_students_capacity > 0),
    seats_available INT NOT NULL DEFAULT total_students_capacity CHECK (seats_available >= 0 AND seats_available <= total_students_capacity),
    
    -- Driver/Helper Information
    driver_helper_id UUID REFERENCES driver_helpers(id) ON DELETE CASCADE,

    -- Document Details
    insurance_number VARCHAR(20) NOT NULL,
    insurance_expiry_date DATE NOT NULL,
    pollution_certificate_number VARCHAR(20) NOT NULL,
    pollution_certificate_expiry_date DATE NOT NULL,
    fitness_certificate_number VARCHAR(20) NOT NULL,
    fitness_certificate_expiry_date DATE NOT NULL,
    vehicle_document_path VARCHAR(255),

    -- Timestamps
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);