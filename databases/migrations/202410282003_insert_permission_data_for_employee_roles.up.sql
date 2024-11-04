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
    'EmployeeRole',
    'List',
    'Employee Role',
    'List',
    'GET',
    '/admin/v1/employee-roles',
    'employee_roles',
    1,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'EmployeeRole',
    'View',
    'Employee Role',
    'View',
    'GET',
    '/admin/v1/employee-roles/:id',
    'employee_roles',
    2,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Employee Role',
    'Create',
    'EmployeeRole',
    'Create',
    'POST',
    '/admin/v1/employee-roles',
    'employee_roles',
    3,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'EmployeeRole',
    'Edit',
    'Employee Role',
    'Edit',
    'PUT',
    '/admin/v1/employee-roles/:id',
    'employee_roles',
    4,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'EmployeeRole',
    'UpdateStatus',
    'Employee Role',
    'Update Status',
    'PUT',
    '/admin/v1/employee-roles/:id/status',
    'employee_roles',
    5,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  );