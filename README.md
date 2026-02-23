# Domain App

A Go server that shows custom text based on the request domain.

## Requirements

- Docker and Docker Compose

## Getting Started

1. Start the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

2. The application will start on port 8080. It will automatically run database migrations on startup.

3. To test it, you'll need to add entries to the `sites` table in the database.

   Example:
   ```sql
   INSERT INTO sites (domain, custom_text) VALUES ('example.com', 'Welcome to Example.com');
   INSERT INTO sites (domain, custom_text) VALUES ('sub.example.com', 'This is a subdomain');
   ```

4. You can then test it using `curl` or by editing your `/etc/hosts`:
   ```bash
   curl -H "Host: example.com" http://localhost:8080
   curl -H "Host: unknown.com" http://localhost:8080
   ```

## Configuration

The application is configured via environment variables:

- `DATABASE_URL`: PostgreSQL connection string.
- `PORT`: Port the server listens on (default: 8080).
