-- +goose Up
-- +goose StatementBegin
CREATE TABLE admins (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL,
	email_address TEXT NOT NULL UNIQUE,
	password_digest TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE admin_sessions (
	id UUID PRIMARY KEY,
	admin_id UUID NOT NULL REFERENCES admins(id),
	token TEXT NOT NULL UNIQUE,
	ip_address TEXT NOT NULL,
	user_agent TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE borrowers (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL,
	email_address TEXT NOT NULL UNIQUE,
	password_digest TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE borrower_sessions (
	id UUID PRIMARY KEY,
	borrower_id UUID NOT NULL REFERENCES borrowers(id),
	token TEXT NOT NULL UNIQUE,
	ip_address TEXT NOT NULL,
	user_agent TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE investors (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL,
	email_address TEXT NOT NULL UNIQUE,
	password_digest TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE investor_sessions (
	id UUID PRIMARY KEY,
	investor_id UUID NOT NULL REFERENCES investors(id),
	token TEXT NOT NULL UNIQUE,
	ip_address TEXT NOT NULL,
	user_agent TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE loans (
	id UUID PRIMARY KEY,
	borrower_id UUID NOT NULL REFERENCES borrowers(id),
	principal_amount BIGINT NOT NULL,
	status TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE approvals (
	id UUID PRIMARY KEY,
	loan_id UUID NOT NULL UNIQUE REFERENCES loans(id),
	admin_id UUID NOT NULL REFERENCES admins(id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE investments (
	id UUID PRIMARY KEY,
	loan_id UUID NOT NULL REFERENCES loans(id),
	investor_id UUID NOT NULL REFERENCES investors(id),
	amount BIGINT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE disbursements (
	id UUID PRIMARY KEY,
	loan_id UUID NOT NULL UNIQUE REFERENCES loans(id),
	admin_id UUID NOT NULL REFERENCES admins(id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE disbursements;

DROP TABLE investments;

DROP TABLE approvals;

DROP TABLE loans;

DROP TABLE investor_sessions;

DROP TABLE investors;

DROP TABLE borrower_sessions;

DROP TABLE borrowers;

DROP TABLE admin_sessions;

DROP TABLE admins;
-- +goose StatementEnd
