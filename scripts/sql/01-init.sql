-- Initialize HCM Database
-- This script will run when the SQL Server container starts

-- Wait for SQL Server to be fully ready
WAITFOR DELAY '00:00:05'

-- Create the main database if it doesn't exist
IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = 'hcmdb')
BEGIN
    PRINT 'Creating database hcmdb...'
    CREATE DATABASE hcmdb
    PRINT 'Database hcmdb created successfully!'
END
ELSE
BEGIN
    PRINT 'Database hcmdb already exists'
END
GO

-- Switch to the hcmdb database
USE hcmdb
GO

-- Create users table (example based on your domain structure)
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'users')
BEGIN
    PRINT 'Creating users table...'
    CREATE TABLE users (
        id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        email NVARCHAR(255) NOT NULL,
        name NVARCHAR(255) NOT NULL,
        created_at DATETIME2 DEFAULT GETDATE(),
        updated_at DATETIME2 DEFAULT GETDATE(),
        CONSTRAINT UK_users_email UNIQUE (email)
    )
    PRINT 'Users table created successfully!'
END
ELSE
BEGIN
    PRINT 'Users table already exists'
END
GO

-- Create index on email for better performance
IF NOT EXISTS (SELECT * FROM sys.indexes WHERE name = 'idx_users_email' AND object_id = OBJECT_ID('users'))
BEGIN
    PRINT 'Creating index on users.email...'
    CREATE INDEX idx_users_email ON users(email)
    PRINT 'Index created successfully!'
END
ELSE
BEGIN
    PRINT 'Index on users.email already exists'
END
GO

-- Insert sample data (optional)
IF NOT EXISTS (SELECT 1 FROM users)
BEGIN
    PRINT 'Inserting sample data...'
    INSERT INTO users (email, name) VALUES 
        ('admin@example.com', 'System Administrator'),
        ('user@example.com', 'Test User')
    PRINT 'Sample data inserted successfully!'
END
ELSE
BEGIN
    PRINT 'Sample data already exists'
END
GO

PRINT 'Database initialization completed successfully!'
PRINT 'You can now connect to database: hcmdb'