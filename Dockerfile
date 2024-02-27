# Use the official PostgreSQL image as the base
FROM postgres

# Set environment variables
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=sample
ENV POSTGRES_DB=rssagg

# Expose PostgreSQL port
EXPOSE 5432
