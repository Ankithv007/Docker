#### Log in to Docker Hub: Run the following command and enter your Docker Hub credentials when prompted:
```
docker login -u
```
```
docker login
```
#### Tag the Image: Make sure your image is tagged correctly with your Docker Hub username. For example:
```
docker tag my-django-app:latest <your-dockerhub-username>/<img you created -( my-django-app)>:latest
```
#### if have not done while you build the iamge (so do it while building only its look cool also )
```
docker build -t <your-dockerhub-username>/<img you created with version or latest( my-django-app)> .
```
#### Push the Image: Now, try pushing the image again with the correct tag:
```
docker push <your-dockerhub-username>/my-django-app:latest
```
