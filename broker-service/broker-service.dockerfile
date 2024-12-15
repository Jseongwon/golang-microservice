FROM alpine:latest

# Set up working directory
WORKDIR /app

# Copy the built application binary
COPY brokerApp /app/brokerApp

# Expose the application port
EXPOSE 8080

# Set the entry point for the container
CMD [ "/app/brokerApp" ]