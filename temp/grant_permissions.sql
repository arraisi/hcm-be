-- Grant EXECUTE permission on the stored procedure
-- Replace 'your_username' with the actual database user that your application uses

-- Option 1: Grant to a specific user
GRANT EXECUTE ON sp_get_next_hasjrat_running TO [your_username];

-- Option 2: Grant to a role (if your user is in a role)
-- GRANT EXECUTE ON sp_get_next_hasjrat_running TO [role_name];

-- Option 3: Grant to the public role (use with caution in production)
-- GRANT EXECUTE ON sp_get_next_hasjrat_running TO PUBLIC;

-- To check current permissions:
SELECT 
    dp.name AS principal_name,
    dp.type_desc,
    o.name AS object_name,
    p.permission_name,
    p.state_desc
FROM sys.database_permissions p
JOIN sys.database_principals dp ON p.grantee_principal_id = dp.principal_id
JOIN sys.objects o ON p.major_id = o.object_id
WHERE o.name = 'sp_get_next_hasjrat_running';

-- To see what user you're currently logged in as:
SELECT SUSER_SNAME() AS login_name, USER_NAME() AS database_user;
