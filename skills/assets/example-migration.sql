-- migrations/000001_create_core_tables.up.sql
-- Demonstrates naming conventions for tables, columns, indexes, and constraints

-- ============================================================
-- users (plural, snake_case)
-- ============================================================
CREATE TABLE users (
    id            UUID PRIMARY KEY,
    email         VARCHAR(255) UNIQUE NOT NULL,
    first_name    VARCHAR(100) NOT NULL,
    last_name     VARCHAR(100) NOT NULL,
    phone         VARCHAR(20),
    is_active     BOOLEAN NOT NULL DEFAULT TRUE,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX uq_users_email ON users(email);

-- ============================================================
-- pets (plural, foreign key = user_id = singular + _id)
-- ============================================================
CREATE TABLE pets (
    id            UUID PRIMARY KEY,
    user_id       UUID NOT NULL REFERENCES users(id),
    name          VARCHAR(100) NOT NULL,
    species       VARCHAR(50) NOT NULL,
    breed         VARCHAR(100),
    is_neutered   BOOLEAN NOT NULL DEFAULT FALSE,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_pets_user_id ON pets(user_id);

-- ============================================================
-- pet_sitters (multi-word table, snake_case plural)
-- ============================================================
CREATE TABLE pet_sitters (
    id              UUID PRIMARY KEY,
    user_id         UUID NOT NULL REFERENCES users(id),
    bio             TEXT,
    is_verified     BOOLEAN NOT NULL DEFAULT FALSE,
    has_certificate BOOLEAN NOT NULL DEFAULT FALSE,
    rating_count    INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_pet_sitters_user_id ON pet_sitters(user_id);

-- ============================================================
-- services
-- ============================================================
CREATE TABLE services (
    id          UUID PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- pet_sitter_services (join table with extra attributes)
-- ============================================================
CREATE TABLE pet_sitter_services (
    id              UUID PRIMARY KEY,
    pet_sitter_id   UUID NOT NULL REFERENCES pet_sitters(id),
    service_id      UUID NOT NULL REFERENCES services(id),
    price_cents     INTEGER NOT NULL,
    is_available    BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_pet_sitter_services_pet_sitter_id ON pet_sitter_services(pet_sitter_id);
CREATE INDEX idx_pet_sitter_services_service_id ON pet_sitter_services(service_id);
CREATE UNIQUE INDEX uq_pet_sitter_services_sitter_service ON pet_sitter_services(pet_sitter_id, service_id);

-- ============================================================
-- bookings (status as enum type)
-- ============================================================
CREATE TYPE booking_status AS ENUM (
    'PENDING',
    'CONFIRMED',
    'IN_PROGRESS',
    'COMPLETED',
    'CANCELLED'
);

CREATE TABLE bookings (
    id                      UUID PRIMARY KEY,
    user_id                 UUID NOT NULL REFERENCES users(id),
    pet_sitter_service_id   UUID NOT NULL REFERENCES pet_sitter_services(id),
    status                  booking_status NOT NULL DEFAULT 'PENDING',
    start_at                TIMESTAMPTZ NOT NULL,
    end_at                  TIMESTAMPTZ NOT NULL,
    total_cents             INTEGER NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_bookings_pet_sitter_service_id ON bookings(pet_sitter_service_id);
CREATE INDEX idx_bookings_status ON bookings(status);

-- ============================================================
-- payments (monetary values in integer cents)
-- ============================================================
CREATE TYPE payment_status AS ENUM (
    'PENDING',
    'PROCESSING',
    'COMPLETED',
    'FAILED',
    'REFUNDED'
);

CREATE TABLE payments (
    id          UUID PRIMARY KEY,
    booking_id  UUID NOT NULL REFERENCES bookings(id),
    status      payment_status NOT NULL DEFAULT 'PENDING',
    flow_token  VARCHAR(255),
    amount_cents INTEGER NOT NULL,
    fee_cents    INTEGER NOT NULL,
    paid_at     TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payments_booking_id ON payments(booking_id);
CREATE INDEX idx_payments_status ON payments(status);
