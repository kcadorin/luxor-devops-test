# Use an official Python runtime as the base image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Copy the Python script and configuration file into the container
COPY server.py config.json ./

# Install dependencies
RUN pip install flask requests

# Expose the port the HTTP server listens on
EXPOSE 8000

# Run the HTTP server when the container starts
CMD ["python", "server.py"]
