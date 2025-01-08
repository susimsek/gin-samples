-- Down Migration: Drop all tables in reverse order

-- Drop user_role_mapping table and constraints
DROP TABLE IF EXISTS user_role_mapping;

-- Drop role table
DROP TABLE IF EXISTS role;

-- Drop user_identity table
DROP TABLE IF EXISTS user_identity;
