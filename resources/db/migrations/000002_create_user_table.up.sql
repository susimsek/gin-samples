-- Up Migration: Create user_identity, role, and user_role_mapping tables with indexes and relationships

-- Create user_identity table
CREATE TABLE IF NOT EXISTS user_identity (
    id TEXT PRIMARY KEY, -- Unique identifier for the user
    username TEXT NOT NULL UNIQUE, -- Username of the user
    password TEXT NOT NULL, -- Password (hashed)
    email TEXT NOT NULL UNIQUE, -- Email of the user
    first_name TEXT, -- First name of the user
    last_name TEXT, -- Last name of the user
    enabled BOOLEAN NOT NULL, -- Is the user enabled?
    created_at DATETIME NOT NULL, -- Creation timestamp
    created_by TEXT NOT NULL, -- Creator
    updated_at DATETIME, -- Last update timestamp
    updated_by TEXT -- Last updater
);

-- Create indexes for user_identity
CREATE INDEX IF NOT EXISTS idx_user_identity_username ON user_identity (username); -- Fast username search
CREATE INDEX IF NOT EXISTS idx_user_identity_email ON user_identity (email); -- Fast email search

-- Create role table
CREATE TABLE IF NOT EXISTS role (
    id TEXT PRIMARY KEY, -- Unique identifier for the role
    name TEXT NOT NULL UNIQUE, -- Name of the role
    description TEXT, -- Description of the role
    created_at DATETIME NOT NULL, -- Creation timestamp
    created_by TEXT NOT NULL, -- Creator
    updated_at DATETIME, -- Last update timestamp
    updated_by TEXT -- Last updater
);

-- Create indexes for role
CREATE INDEX IF NOT EXISTS idx_role_name ON role (name); -- Fast role name search

-- Create user_role_mapping table
CREATE TABLE IF NOT EXISTS user_role_mapping (
    user_id TEXT NOT NULL, -- Foreign key to user_identity
    role_id TEXT NOT NULL, -- Foreign key to role
    created_at DATETIME NOT NULL, -- Creation timestamp
    created_by TEXT NOT NULL, -- Creator
    updated_at DATETIME, -- Last update timestamp
    updated_by TEXT, -- Last updater
    PRIMARY KEY (user_id, role_id), -- Composite primary key
    FOREIGN KEY (user_id) REFERENCES user_identity (id) ON DELETE CASCADE, -- Inline foreign key to user_identity
    FOREIGN KEY (role_id) REFERENCES role (id) ON DELETE CASCADE -- Inline foreign key to role
);
