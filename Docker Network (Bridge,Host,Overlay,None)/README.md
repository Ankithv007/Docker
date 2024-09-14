# Types of Docker Networks:
- Bridge Network (default for standalone containers):
- Host Network:
- None:
- Overlay Network:

  ### Bridge Network (default for standalone containers):
  - Use case: Best for when you want containers to communicate with each other on the same host but isolate them from the outside world.
  - Behavior: Containers on the same bridge network can talk to each other using container names as hostnames.
 
  ### What is Host Network?
  #### In Docker, each container typically has its own virtual network interface, meaning it’s isolated from the host machine's network. However, with the host network mode, the container does not get its own network interface. Instead, it shares the network stack with the host machine (the computer where Docker is running).

 #### No Network Isolation:
- The container shares the same network interface and IP address as the host machine.
- Whatever ports the container exposes, they are directly exposed on the host machine.
#### Ports Behavior:
- Normally, when you run a container, you need to map its internal ports to the host's ports (e.g., docker run -p 8080:8080). In host network mode, you don’t need to do this.
- If your application inside the container runs on port 8080, it will be available directly on the host’s port 8080, without any port mapping.
#### (create one and show that how it work if any doubt go to Docker book and refer there )

## What is the None Network Mode?
 #### In Docker, by default, each container is connected to a network, whether it's the bridge network or any custom network. However, with the none network mode, the container is completely isolated from any network—this means:

- The container won’t have any network interface.
- It cannot communicate with other containers or the host machine.
- It has no access to the internet or any external network.

### Key Points:
###  No Communication:
- The container cannot send or receive any network traffic.
- There is no way for other containers or external systems to connect to this container.
### No Network Interface:
- Normally, each container gets a network interface (e.g., eth0) to handle communication, but in none mode, this interface is not created.
- The container is isolated from any kind of network.
```
  docker run --network none my-app
```
- The container starts, but it has no network capability.
- It cannot talk to any other containers, the host, or the internet.

  ## What is an Overlay Network? (for more we will do it in k8s (just to know we can achive form this also )
  #### An Overlay Network allows containers on different Docker hosts to communicate with each other as if they were on the same local network. This is useful for applications that are distributed across multiple servers.

 ###  Key Points:
 ### Multi-Host Communication:
- Containers on different physical or virtual machines (hosts) can talk to each other.
- The network abstracts away the complexity of routing traffic between these different hosts.
### Use in Orchestration:
- Docker Swarm and Kubernetes use Overlay Networks to connect containers that are part of the same service or application, even if they are deployed on different hosts.
- This network mode supports scaling out applications across a cluster of machines.
  ### Benefits:
- Seamless Communication: Containers can use their names to resolve and communicate with each other, regardless of the host they are on.
- Scalability: Easily scale services across multiple hosts without worrying about network complexity.
- Security: Overlay networks can provide encrypted communication between containers.
  #### An Overlay Network abstracts the network details and allows containers across multiple Docker hosts to communicate as if they were on the same local network. This is crucial for container orchestration platforms like Docker Swarm and Kubernetes, which manage containerized applications across a cluster of machines.
