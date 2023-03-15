
FROM mcr.microsoft.com/mssql/server:2022-latest

# Switch to root user for access to apt-get install



# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app


# Bundle app source
COPY . /usr/src/app



# Grant permissions for the import-data script to be executable
RUN chmod +x /usr/src/app/import-data.sh

# Switch back to mssql user and run the entrypoint script


ENTRYPOINT /bin/bash ./entrypoint.sh


ENV SA_PASSWORD S3cur3P@ssW0rd!
ENV ACCEPT_EULA Y
