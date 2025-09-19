-- +goose Up
-- +goose StatementBegin
CREATE TABLE admins (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email_address TEXT UNIQUE NOT NULL,
    password_digest TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE admin_sessions (
    id UUID PRIMARY KEY,
    admin_id UUID REFERENCES admins(id),
    token TEXT UNIQUE NOT NULL,
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE borrowers (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email_address TEXT UNIQUE NOT NULL,
    password_digest TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE borrower_sessions (
    id UUID PRIMARY KEY,
    borrower_id UUID REFERENCES borrowers(id),
    token TEXT UNIQUE NOT NULL,
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE investors (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email_address TEXT UNIQUE NOT NULL,
    password_digest TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE investor_sessions (
    id UUID PRIMARY KEY,
    investor_id UUID REFERENCES investors(id),
    token TEXT UNIQUE NOT NULL,
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE investor_sessions;

DROP TABLE investors;

DROP TABLE borrower_sessions;

DROP TABLE borrowers;

DROP TABLE admin_sessions;

DROP TABLE admins;
-- +goose StatementEnd
