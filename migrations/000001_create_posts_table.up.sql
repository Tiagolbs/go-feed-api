CREATE TABLE IF NOT EXISTS posts (
	id bigserial PRIMARY KEY,
	content text NOT NULL,
	created_at TIMESTAMP with time zone NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP with time zone NOT NULL DEFAULT NOW()
)