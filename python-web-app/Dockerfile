FROM ubuntu:latest

# Install Python and pip
RUN apt-get update && \
    apt-get install -y python3 python3-pip python3-venv

# Create and activate a virtual environment
RUN python3 -m venv /venv
ENV PATH="/venv/bin:$PATH"

# Copy application files
WORKDIR /app
COPY requirements.txt /app/
COPY devops /app/

# Install Python packages within the virtual environment
RUN pip install --upgrade pip && \
    pip install -r requirements.txt

# Set the default command
ENTRYPOINT ["python3"]
CMD ["manage.py", "runserver", "0.0.0.0:8000"]
