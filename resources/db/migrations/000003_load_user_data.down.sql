-- Delete data from user_role_mapping
DELETE FROM user_role_mapping
WHERE user_id IN ('1e7d07a7-896e-41a7-bb47-8ccedb9c9fc3', '5f74d85e-87d1-4c1c-b0c1-ec54d1d23f57')
  AND role_id IN ('c2e7d07a-896e-41a7-bb47-8ccedb9c9fc3', 'd3f74d85-87d1-4c1c-b0c1-ec54d1d23f57');

-- Delete data from role
DELETE FROM role
WHERE id IN ('c2e7d07a-896e-41a7-bb47-8ccedb9c9fc3', 'd3f74d85-87d1-4c1c-b0c1-ec54d1d23f57');

-- Delete data from user_identity
DELETE FROM user_identity
WHERE id IN ('1e7d07a7-896e-41a7-bb47-8ccedb9c9fc3', '5f74d85e-87d1-4c1c-b0c1-ec54d1d23f57');
