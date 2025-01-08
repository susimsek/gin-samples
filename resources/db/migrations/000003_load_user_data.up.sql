-- Insert data into user_identity
INSERT OR IGNORE INTO user_identity (id, username, password, email, first_name, last_name, enabled, created_at, created_by, updated_at, updated_by)
VALUES
  ('1e7d07a7-896e-41a7-bb47-8ccedb9c9fc3', 'admin', '$2a$10$45h4TdLTwTCtLIRThucXLuPOMtALeRErlNU5Ch2GkwZIWojh7mTOe', 'admin@example.com', 'Admin', 'User', 1, '2023-07-13 10:00:00.533433', 'system', NULL, NULL),
  ('5f74d85e-87d1-4c1c-b0c1-ec54d1d23f57', 'user', '$2a$10$45h4TdLTwTCtLIRThucXLuPOMtALeRErlNU5Ch2GkwZIWojh7mTOe', 'user@example.com', 'User', 'User', 1, '2023-07-13 10:00:00.533433', 'system', NULL, NULL);

-- Insert data into role
INSERT OR IGNORE INTO role (id, name, description, created_at, created_by, updated_at, updated_by)
VALUES
  ('c2e7d07a-896e-41a7-bb47-8ccedb9c9fc3', 'ROLE_ADMIN', 'Administrator role', '2023-07-13 10:00:00.533433', 'system', NULL, NULL),
  ('d3f74d85-87d1-4c1c-b0c1-ec54d1d23f57', 'ROLE_USER', 'User role', '2023-07-13 10:00:00.533433', 'system', NULL, NULL);

-- Insert data into user_role_mapping
INSERT OR IGNORE INTO user_role_mapping (user_id, role_id, created_at, created_by, updated_at, updated_by)
VALUES
  ('1e7d07a7-896e-41a7-bb47-8ccedb9c9fc3', 'c2e7d07a-896e-41a7-bb47-8ccedb9c9fc3', '2023-07-13 10:00:00.533433', 'system', NULL, NULL),
  ('5f74d85e-87d1-4c1c-b0c1-ec54d1d23f57', 'd3f74d85-87d1-4c1c-b0c1-ec54d1d23f57', '2023-07-13 10:00:00.533433', 'system', NULL, NULL);
