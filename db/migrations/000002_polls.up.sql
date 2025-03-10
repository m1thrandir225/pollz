CREATE TABLE polls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description TEXT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    created_by UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL ,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
)