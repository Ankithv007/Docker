# Step 1: Use an official OpenJDK runtime as a parent image
FROM openjdk:11-jre-slim

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the current directory contents into the container at /app
COPY . /app

# Step 4: Install any necessary dependencies (in this case, curl)
RUN apt-get update && apt-get install -y curl

# Step 5: Set environment variables
ENV APP_ENV=production
ENV APP_VERSION=1.0

# Step 6: Expose the port your application will run on
EXPOSE 8080

# Step 7: Define a volume to persist data
VOLUME ["/app/data"]

# Step 8: Add a tar file containing additional resources
ADD resources.tar.gz /app/resources/

# Step 9: Set up a health check to monitor the container's health
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

# Step 10: Switch to a non-root user
USER 1001

# Step 11: Set the entrypoint to always run the Java application
ENTRYPOINT ["java", "-jar", "app.jar"]

# Step 12: Use CMD to provide default arguments that can be overridden at runtime
CMD ["--server.port=8080"]

 - FROM openjdk:11-jre-slim: This uses the OpenJDK 11 slim image as the base.
 - WORKDIR /app: Sets /app as the working directory.
 - COPY . /app: Copies the local directory content to /app in the container.
 - RUN apt-get update && apt-get install -y curl: Updates the package list and installs curl.
 - ENV APP_ENV=production and ENV APP_VERSION=1.0: Sets environment variables for the application.
 - EXPOSE 8080: Exposes port 8080 for the application.
 - VOLUME ["/app/data"]: Creates a volume for persisting data in the /app/data directory.
 - ADD resources.tar.gz /app/resources/: Adds and extracts the resources.tar.gz file into the /app/resources/ directory.
 - HEALTHCHECK: Defines a health check that will call the /health endpoint on port 8080 every 30 seconds.
 - USER 1001: Switches the user to UID 1001 (non-root user) for security purposes.
 - ENTRYPOINT ["java", "-jar", "app.jar"]: Specifies the Java JAR file as the entry point.
 - CMD ["--server.port=8080"]: Provides default command arguments to set the server port to 8080, which can be overridden if necessary.


###  FROM:
- `FROM <base-image>`
- Purpose: This defines the base image that your Docker container will be built on top of.
- Example: FROM openjdk:11-jre-slim
- Explanation: This tells Docker to use the OpenJDK version 11 with a slimmed-down JRE as the base for the container. The base image provides the environment that your application needs to run, such as the operating system and any pre-installed libraries.

### WORKDIR:
- `WORKDIR /app`
- Purpose: This sets the working directory inside the container.
- Explanation: Docker will switch to this directory before executing the rest of the instructions in the Dockerfile. In this case, it's setting the working directory to /app inside the container. Any files copied or commands run will take place within this directory.

###  COPY:
- `COPY <source> <destination>`
- Purpose: This copies files or directories from your local file system (source) into the Docker image (destination).
- Example: COPY . /app
- Explanation: It copies everything in the current directory (denoted by .) from the local system into the /app directory in the container.

  
### RUN:
  - `RUN <command> `
  - Purpose: This allows you to run any shell commands needed to set up your application environment (such as installing packages).
  - Example: RUN apt-get update && apt-get install -y curl
  - Explanation: This will update the package list inside the container and install curl. These changes will be included in the image.

### CMD
- `CMD ["executable", "param1", "param2"`
- Purpose: This specifies the default command to run when the container starts.
- Example: CMD ["java", "-jar", "app.jar"]
- Explanation: This tells Docker to run the java -jar app.jar command when the container starts. If you provide a different command at runtime, the CMD will be overridden.

###  ENTRYPOINT:
- `ENTRYPOINT ["executable", "param1"]`
- Purpose: This is similar to CMD, but it is not overridden by command-line arguments passed to the container. It defines the main application that should always run.
- Example: ENTRYPOINT ["java", "-jar", "app.jar"]
- Explanation: This makes sure that the java -jar app.jar is always executed when the container starts, even if additional arguments are provided at runtime

### EXPOSE:
- `EXPOSE <port>`
- Purpose: This informs Docker that the container will listen on the specified network port at runtime.
- Example: EXPOSE 8080
- Explanation: This exposes port 8080 to allow communication between the container and the outside world (but does not automatically publish it to your host).

 ### ENV:
 - `ENV <key>=<value>`
 - Purpose: This sets environment variables inside the container.
 - Example: ENV APP_ENV=production
 - Explanation: This defines an environment variable APP_ENV with the value production that will be available to your application when it runs.

### ADD:
 - ` ADD <source> <destination> `
 - Purpose: This is similar to COPY, but it can also retrieve files from remote URLs and automatically decompress tar files.
 - Example: ADD my-archive.tar.gz /app/
 - Explanation: This will copy the my-archive.tar.gz file into the /app/ directory inside the container and automatically extract it.

 ### VOLUME:
 - `VOLUME ["/path/to/directory"]`
 - Purpose: This is used to define a mount point for a volume. Volumes are used to persist data outside of the container's lifecycle.
 - Example: VOLUME ["/data"]
 - Explanation: This creates a volume at /data in the container, allowing data stored there to persist across container restarts or even after the container is destroyed.

###  USER:
- `USER <username or UID>`
- Purpose: This switches the user the container will use when running commands.
- Example: USER 1001
- Explanation: This tells Docker to run the container using the user with UID 1001, which enhances security by avoiding running as root.

###  HEALTHCHECK:
- `HEALTHCHECK --interval=<time> --timeout=<time> --retries=<num> CMD <command>`
- Purpose: This defines a command to check whether the container is still healthy.
- Example: HEALTHCHECK --interval=30s --timeout=3s CMD curl -f http://localhost/ || exit 1
- Explanation: Docker will check every 30 seconds whether the container is healthy by sending a request to http://localhost/. If it fails, the container will be marked as unhealthy.


