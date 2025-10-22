# SQL Server Database Setup

This project is configured to use Microsoft SQL Server as the database backend. All configuration is managed through `internal/config/config.yaml`.

## Quick Start

1. **Start SQL Server:**

   ```bash
   docker-compose up sqlserver -d
   ```

2. **Initialize the database (first time only):**

   ```bash
   docker-compose up sqlserver-init
   ```

   Or manually:

   ```bash
   docker exec hcm-sqlserver /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P "YourStrong@Passw0rd" -C -i /scripts/01-init.sql
   ```

3. **Verify the database is running:**

   ```bash
   docker-compose logs sqlserver
   ```

4. **Test the connection:**

   ```bash
   docker exec hcm-sqlserver /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P "YourStrong@Passw0rd" -C -d hcmdb -Q "SELECT COUNT(*) FROM users"
   ```

5. **Update your application config:**
   Edit `internal/config/config.yaml`:
   ```yaml
   database:
     driver: sqlserver
     dsn: 'sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=hcmdb&encrypt=disable&trustServerCertificate=true'
   ```

## Configuration Options

### Configuration File

All database configuration is managed in `internal/config/config.yaml`:

```yaml
database:
  driver: sqlserver
  dsn: 'sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=hcmdb&encrypt=disable&trustServerCertificate=true'
```

### Connection String Format

```
sqlserver://username:password@host:port?database=dbname&options
```

Common options:

- `encrypt=disable`: Disable encryption (development only)
- `trustServerCertificate=true`: Trust self-signed certificates
- `encrypt=true`: Enable encryption (production)

## Database Management

### Connect with SQL Tools

You can connect to the database using:

- **SQL Server Management Studio (SSMS)**
- **Azure Data Studio**
- **VS Code SQL Server Extension**

Connection details:

- Server: `localhost,1433`
- Authentication: SQL Server Authentication
- Username: `sa`
- Password: `YourStrong@Passw0rd` (as configured in config.yaml)

### Database Schema

The database is automatically initialized with:

- `hcmdb` database
- `users` table with basic structure
- Sample data (optional)

### Backup and Restore

**Backup:**

```bash
docker exec hcm-sqlserver /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P "YourStrong@Passw0rd" -C -Q "BACKUP DATABASE [hcmdb] TO DISK = N'/var/opt/mssql/data/hcmdb.bak'"
```

**Restore:**

```bash
docker exec hcm-sqlserver /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P "YourStrong@Passw0rd" -C -Q "RESTORE DATABASE [hcmdb] FROM DISK = N'/var/opt/mssql/data/hcmdb.bak'"
```

## Development Commands

**Start all services:**

```bash
docker-compose up -d
```

**Stop all services:**

```bash
docker-compose down
```

**View database logs:**

```bash
docker-compose logs -f sqlserver
```

**Reset database (removes all data):**

```bash
docker-compose down -v
docker-compose up sqlserver -d
```

## Production Considerations

1. **Change default passwords** in production
2. **Enable encryption** by setting `encrypt=true` in DSN
3. **Use proper SSL certificates** instead of `trustServerCertificate=true`
4. **Configure backup strategies**
5. **Set up proper monitoring**
6. **Review security settings**

## Troubleshooting

### Common Issues

1. **Container fails to start:**

   - Check password complexity requirements
   - Ensure port 1433 is not in use
   - Check available disk space

2. **Connection refused:**

   - Verify container is running: `docker ps`
   - Check health status: `docker-compose ps`
   - Review logs: `docker-compose logs sqlserver`

3. **Authentication failed:**
   - Verify SA password in config.yaml
   - Check connection string format
   - Ensure database exists

### Health Check

The SQL Server container includes a health check that verifies:

- Service is responding
- Authentication is working
- Database is accessible

Check health status:

```bash
docker-compose ps
```
