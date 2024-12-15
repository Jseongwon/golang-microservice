FROM alpine:latest

# Set up working directory
WORKDIR /app

# Copy the built application binary
COPY authApp /app/authApp

# Set the entry point for the container
CMD [ "/app/authApp" ]