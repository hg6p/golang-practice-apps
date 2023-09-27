# Use the official PostgreSQL image from Docker Hub
FROM postgres

# Set environment variables for PostgreSQL
ENV POSTGRES_USER davor
ENV POSTGRES_PASSWORD 1234
ENV POSTGRES_DB urlshort

# Copy custom initialization scripts into the container
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432
