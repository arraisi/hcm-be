-- Grant EXECUTE permission to hcmdev user
GRANT EXECUTE ON sp_get_next_hasjrat_running TO [hcmdev];
GO

-- Verify the permission was granted
SELECT 
    dp.name AS principal_name,
    dp.type_desc,
    o.name AS object_name,
    p.permission_name,
    p.state_desc
FROM sys.database_permissions p
JOIN sys.database_principals dp ON p.grantee_principal_id = dp.principal_id
JOIN sys.objects o ON p.major_id = o.object_id
WHERE o.name = 'sp_get_next_hasjrat_running'
  AND dp.name = 'hcmdev';
