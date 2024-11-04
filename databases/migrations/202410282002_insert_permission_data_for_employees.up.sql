INSERT INTO
  `permissions`(
    package,
    module_name,
    action_name,
    display_module_name,
    display_action_name,
    http_method,
    route,
    table_name,
    sequence_number_detail,
    created_at,
    created_by,
    updated_at,
    updated_by
  )
VALUES
  (
    'WebsiteAdmin',
    'Employee',
    'List',
    'Employee',
    'List',
    'GET',
    '/admin/v1/employees',
    'employees',
    1,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Employee',
    'View',
    'Employee',
    'View',
    'GET',
    '/admin/v1/employees/:id',
    'employees',
    2,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Employee',
    'Create',
    'Employee',
    'Create',
    'POST',
    '/admin/v1/employees',
    'employees',
    3,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Employee',
    'Edit',
    'Employee',
    'Edit',
    'PUT',
    '/admin/v1/employees/:id',
    'employees',
    4,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Employee',
    'EditPassword',
    'Employee',
    'Edit Password',
    'PUT',
    '/admin/v1/employees/:id/password',
    'employees',
    5,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Employee',
    'ResetPassword',
    'Employee',
    'Reset Password',
    'PUT',
    '/admin/v1/employees/:id/reset-password',
    'employees',
    6,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Employee',
    'UpdateStatus',
    'Employee',
    'Update Status',
    'PUT',
    '/admin/v1/employees/:id/status',
    'employees',
    7,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  );