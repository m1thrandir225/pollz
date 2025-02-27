CREATE TABLE votes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_id UUID REFERENCES poll_options(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id),
    voted_at TIMESTAMPTZ DEFAULT 'now()',
    ip_address TEXT,
    user_agent TEXT
);