CREATE TABLE votes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_id uuid REFERENCES poll_options (id) ON DELETE CASCADE NOT NULL,
    voted_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL
);
