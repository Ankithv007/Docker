## 1 Docker Version
- Command:` docker --version `
- Use: Shows the installed version of Docker.
## 2. Pull an Image
- Command:` docker pull <image_name>`
- Example:` docker pull nginx ` 
- Use: Downloads the specified image from Docker Hub or another Docker registry.
## 3. List Docker Images
- Command: `docker images `
- Use: Lists all the Docker images that are stored locally.
## 4. Run a Container
- Command: `docker run <image_name>`
- Example: ` docker run nginx `
- Use: Runs a container from a specified image. If the image is not available locally, Docker will pull it first.
## 5. Run a Container with Port Mapping
- Command: `docker run -d -p <host_port>:<container_port> <image_name>`
- Example: `docker run -d -p 8080:80 nginx`
- Use: Runs a container in detached mode (-d), mapping the host port 8080 to the container port 80.
## 6. List Running Containers
- Command:` docker ps`
- Use: Lists all running containers.
## 7. List All Containers (including stopped)
- Command:` docker ps -a`
- Use: Lists all containers, including the ones that have stopped.
## 8. Stop a Running Container
- Command: `docker stop <container_id>`
- Example: `docker stop 1a2b3c4d5e6f`
- Use: Stops a running container.
## 9. Start a Stopped Container
- Command:` docker start <container_id>`
- Use: Starts a container that was previously stopped.
## 10. Remove a Container
- Command: `docker rm <container_id>`
- Use: Removes a stopped container. You cannot remove running containers unless you stop them first.
## 11. Remove an Image
- Command: `docker rmi <image_id>`
- Use: Deletes a local Docker image.
## 12. View Logs of a Container
- Command:` docker logs <container_id>`
- Use: Shows the logs of a running or stopped container.
## 13. Run a Container in Interactive Mode
- Command:` docker run -it <image_name> /bin/bash`
- Example: `docker run -it ubuntu /bin/bash`
- Use: Runs a container in interactive mode, allowing you to interact with the container via a shell (/bin/bash).
## 14. Execute a Command in a Running Container
- Command: `docker exec -it <container_id> <command>`
- Example:` docker exec -it 1a2b3c4d5e6f /bin/bash`
- Use: Executes a command inside a running container (e.g., opening a shell in a running container).
## 15. Build an Image from a Dockerfile
- Command:`docker build -t <image_name>` .
- Example:` docker build -t my-app` .
- Use: Builds a Docker image from a Dockerfile in the current directory and tags it with <image_name>.
## 16. Push an Image to Docker Hub
- Command: `docker push <username>/<image_name>`
- Example:` docker push ankithbv007/my-app`
- Use: Uploads the Docker image to Docker Hub (or another registry).
## 17. View Container Details
- Command: `docker inspect <container_id>`
- Use: Displays low-level information about a container or image in JSON format.
## 18. Remove Unused Containers, Networks, Images, and Volumes
- Command: `docker system prune`
- Use: Cleans up unused containers, networks, images, and volumes to free up space.
## 19. View Resource Usage
- Command:` docker stats`
- Use: Displays a live stream of resource usage statistics for running containers.
## 20. Tag an Image
- Command: `docker tag <image_name>:<tag> <new_image_name>:<new_tag>`
- Example:` docker tag my-app:latest ankithbv007/my-app:v1`
- Use: Adds a new tag to an existing Docker image.
