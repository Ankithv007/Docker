# What Are Docker Volumes?
- A Docker volume is a type of storage that Docker manages on your behalf.
- Volumes allow you to store and persist data generated or used by your containers.
- The key point is that volumes exist outside of the container's filesystem, so when a container is deleted, the data in the volume remains.

 ### Why Use Docker Volumes?
- Persistent storage: When containers stop or are removed, any data stored directly inside them is lost. Volumes ensure that your data remains intact, even if you remove the container.
- Sharing data: Volumes allow multiple containers to access the same data. This is useful for shared services, like databases, where multiple containers need to read/write to the same storage.
- Docker-managed: Docker handles where and how the volume is stored on the host machine, making it easier to manage compared to bind mounts.

  ### Where Are Volumes Stored?
  ##### On a Linux system, Docker stores volumes at a default location (usually /var/lib/docker/volumes/). However, you don't need to know or interact with this path—Docker manages it for you.

### Example of Creating and Using Docker Volumes
#### Create a Volume:
```
docker volume create my_volume
```
#### Run a database container using the volume:
```
docker run -d -v db_data:/var/lib/mysql mysql
```
```
docker run -d -v my_volume:/path/in/container my_image
```
- my_volume: The volume you created.
- /path/in/container: The path inside the container where the volume is mounted.
- my_image: The Docker image you are using.
- Now, anything written to /path/in/container inside the container will be saved in my_volume, and the data will persist even if the container is deleted.

### Checking Volumes:(List all volumes)
```
docker volume ls
```
### Inspect a specific volume:
```
docker volume inspect my_volume
```
### Removing a Volume:(before we should stop the conatiner if it is in running  stage)
```
docker volume rm my_volume
```
- Be careful, though! Once a volume is removed, all the data stored in it is lost.

   ## Benefits of Docker Volumes:
 - Easier to manage: Docker takes care of where the volume is stored, and you don’t need to worry about filesystem paths.
 - Better isolation: The data in volumes is isolated from the host’s filesystem, making it safer and easier to manage in multi-container setups.
 - More flexibility: Volumes are easier to share between multiple containers, and they can be mounted into containers at runtime.
 - Backup and restore: Volumes are easy to back up and move between different systems.


### Visualizing Volumes:
###### Imagine you have a file server running in a Docker container that stores user files. You don’t want to lose all those files if the server crashes, so you create a Docker volume to store those files

- Without volumes: If the server container is deleted, all user files are lost.
- With volumes: The user files remain stored in the Docker volume, and you can start a new container that still has access to the same data.

# What is a Bind Mount?
#### A bind mount is when you take a specific folder or file from your computer (the host)(any local it may S3 bucket or ec2) and make it available inside your Docker container.
####  It means that the container can use and modify the files on your computer directly.
#### So, whatever changes happen inside the container are reflected in the files on your computer, and vice versa.

## Example of Bind Mount:
- Let’s say you are working on a web application, and you want to test your code inside a container. You don't want to copy your code into the container every time you change something.
- Instead, you can use a bind mount to link your project folder on your computer to the container.


  #### On your computer (host):
- You have a project folder: /home/ankith/my-web-app/
- This folder contains your web application files.

#### Inside the container:
- You want the application to access this folder so that any changes in the code will be reflected immediately.

 #### Run the container with a bind mount:
 ```
docker run -d -v /home/ankith/my-web-app:/usr/src/app my_image
```
- Here, /home/ankith/my-web-app is the host folder (your project folder).
- /usr/src/app is the container folder where the files will appear inside the container.
- my_image is the Docker image you are running.


### Why Use Bind Mounts?
- Live code updates: Perfect for development. If you edit a file on your computer, it automatically updates inside the container.
- No need to rebuild: You don’t need to rebuild your Docker image every time you make a code change. The changes reflect instantly in the container


### A bind mount is a way to "link" a folder from your computer to a folder inside a Docker container, allowing both to use the same files.
