CREATE TABLE IF NOT EXISTS "role_permission" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "role_guid" varchar NOT NULL,
  "permission_guid" varchar NOT NULL
);

INSERT INTO "role_permission" 
    ("role_guid", "permission_guid")
VALUES
    ((SELECT guid FROM roles WHERE name = 'Super Admin'), (SELECT guid FROM permissions WHERE name = 'User Management')),
    ((SELECT guid FROM roles WHERE name = 'Super Admin'), (SELECT guid FROM permissions WHERE name = 'Role Management')),
    ((SELECT guid FROM roles WHERE name = 'Super Admin'), (SELECT guid FROM permissions WHERE name = 'Project Management')),
    ((SELECT guid FROM roles WHERE name = 'Super Admin'), (SELECT guid FROM permissions WHERE name = 'Task Management')),
    ((SELECT guid FROM roles WHERE name = 'Super Admin'), (SELECT guid FROM permissions WHERE name = 'Audit Logs')),
    ((SELECT guid FROM roles WHERE name = 'Admin'), (SELECT guid FROM permissions WHERE name = 'User Management')),
    ((SELECT guid FROM roles WHERE name = 'Admin'), (SELECT guid FROM permissions WHERE name = 'Project Management')),
    ((SELECT guid FROM roles WHERE name = 'Admin'), (SELECT guid FROM permissions WHERE name = 'Role Management')),
    ((SELECT guid FROM roles WHERE name = 'Admin'), (SELECT guid FROM permissions WHERE name = 'Task Management'))
ON CONFLICT DO NOTHING;