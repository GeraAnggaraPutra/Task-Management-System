CREATE TYPE "action" AS ENUM (
  'Read',
  'Create',
  'Update',
  'Delete'
);

CREATE TABLE IF NOT EXISTS "role_permission_action" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "role_permission_guid" varchar NOT NULL,
  "action" action NOT NULL,
  "is_checked" bool
);

INSERT INTO "role_permission_action" 
    ("role_permission_guid", "action", "is_checked")
VALUES
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), 'Create', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), 'Update', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), 'Delete', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Role Management')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Role Management')), 'Create', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Role Management')), 'Update', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Role Management')), 'Delete', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Project Management')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Project Management')), 'Create', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Project Management')), 'Update', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Project Management')), 'Delete', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Task Management')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Task Management')), 'Create', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Task Management')), 'Update', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Task Management')), 'Delete', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Audit Logs')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Audit Logs')), 'Create', false),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Audit Logs')), 'Update', false),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Super Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Audit Logs')), 'Delete', false),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), 'Create', false),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), 'Update', false),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'User Management')), 'Delete', false),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Project Management')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Project Management')), 'Create', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Project Management')), 'Update', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Project Management')), 'Delete', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Task Management')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Task Management')), 'Create', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Task Management')), 'Update', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Task Management')), 'Delete', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Role Management')), 'Read', true),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Role Management')), 'Create', false),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Role Management')), 'Update', false),
    ((SELECT guid FROM role_permission WHERE role_guid = (SELECT guid FROM roles WHERE name = 'Admin') AND permission_guid = (SELECT guid FROM permissions WHERE name = 'Role Management')), 'Delete', false)
ON CONFLICT DO NOTHING;