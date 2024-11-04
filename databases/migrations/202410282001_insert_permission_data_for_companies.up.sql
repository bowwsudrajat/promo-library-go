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
    'Company',
    'List',
    'Company',
    'List',
    'GET',
    '/admin/v1/company',
    'companies',
    1,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Company',
    'View',
    'Company',
    'View',
    'GET',
    '/admin/v1/company/:id',
    'companies',
    2,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Company',
    'Create',
    'Company',
    'Create',
    'POST',
    '/admin/v1/company',
    'companies',
    3,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Company',
    'Edit',
    'Company',
    'Edit',
    'PUT',
    '/admin/v1/company/:id',
    'companies',
    4,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  ), (
    'WebsiteAdmin',
    'Company',
    'UpdateStatus',
    'Company',
    'Update Status',
    'PUT',
    '/admin/v1/company/:id/status',
    'companies',
    5,
    CURRENT_TIMESTAMP,
    '0',
    CURRENT_TIMESTAMP,
    '0'
  );