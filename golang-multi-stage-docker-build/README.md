# Multi Stage Docker Build

The main purpose of choosing a golang based applciation to demostrate this example is golang is a statically-typed programming language that does not require a runtime in the traditional sense. Unlike dynamically-typed languages like Python, Ruby, and JavaScript, which rely on a runtime environment to execute their code, Go compiles directly to machine code, which can then be executed directly by the operating system.

So the real advantage of multi stage docker build and distro less images can be understand with a drastic decrease in the Image size.


## What Are Multi-Stage Builds?
- Multi-stage builds in Docker allow you to use a single Dockerfile to define multiple stages in the build process
-  Each stage can use a different base image
-   you can copy artifacts between stages. This helps in creating more optimized Docker images.

  ## How It Works:
  - Define Multiple Stages: You use multiple FROM instructions in a single Dockerfile to define different stages of the build process. Each FROM creates a new stage.
  - Build Stage: The initial stage (or stages) is used to compile or build your application. This stage typically uses a larger base image with all the necessary tools for building your app.
  - Final Stage: The final stage uses a different base image, often a smaller one, that contains only the runtime environment needed to run the application. You copy only the necessary files from the build stage into this final image.

## Benefits:
- Smaller Images: The final image is smaller because it excludes build tools and dependencies.
- Cleaner Images: Avoids clutter in the final image, leading to better security and performance.
-  multi-stage builds help in creating more efficient and manageable Docker images by separating the build and runtime environments within a single Dockerfile.

## Why Use Multiple Stages?
- Separate Concerns: Each stage has a distinct role (building, testing, serving).
- Optimization: You can use different base images suited for each stage. For example, using a lightweight image for serving reduces the final image size
- Flexibility: You can add stages for additional tasks like linting, packaging, or security checks.


#### Example Dockerfile with Three Stages
```
# Stage 1: Build the application
FROM node:14 AS build
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build

# Stage 2: Test the application
FROM node:14 AS test
WORKDIR /app
COPY --from=build /app /app
RUN npm test

# Stage 3: Serve the application
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
 ```
  
# Distroless images
- Distroless images are Docker images that contain only the application and its runtime dependencies
- excluding the operating system and unnecessary files
- This approach minimizes the image size and enhances security

## Key Characteristics:
- Minimalist: They include only the essential components needed to run the application. No package managers, shells, or unnecessary utilities are included.
- Security: Fewer components reduce the attack surface, making the image less vulnerable to security threats.
- Efficiency: Smaller image sizes lead to faster deployment and less resource consumption.

## How It Works:
- Base Image: Uses a minimal base image that provides just the runtime environment for your application (e.g., gcr.io/distroless/java for Java applications).
- Application: The image only includes your application and its necessary dependencies, without the underlying OS or additional tools.

  ## Benefits:
  - Reduced Attack Surface: With no package managers or shells, the potential for vulnerabilities is reduced.
  - Smaller Image Size: The final image is lean and focused solely on running the application
  - Distroless images are ideal for production environments where security and efficiency are priorities

## Example Dockerfile:
```
# Build Stage
FROM node:14 AS build
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build

# Final Stage with Distroless Image
FROM gcr.io/distroless/nodejs:14
COPY --from=build /app/build /app/build
CMD ["node", "/app/build/index.js"]

```

- Build Stage: Uses a full Node.js image to build the application.
- Final Stage: Uses a distroless Node.js image, which includes only Node.js runtime and the application.
