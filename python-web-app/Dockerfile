FROM ubuntu:latest

# Set the working directory
WORKDIR /app

# Copy requirements and application code
COPY requirements.txt /app/
COPY devops /app/

# Install Python, pip, and create a virtual environment
RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    pip3 install virtualenv && \
    virtualenv venv && \
    /venv/bin/pip install -r requirements.txt

# Use the virtual environment for running the application
ENV PATH="/app/venv/bin:$PATH"

# Set the entrypoint and default command
ENTRYPOINT ["python3"]
CMD ["manage.py", "runserver", "0.0.0.0:8000"]
