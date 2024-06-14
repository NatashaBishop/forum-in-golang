# Forum

---

Authors: [Natalia Korba](https://github.com/NatashaBishop), Michael Young

Solved during fellowship in 01Founders coding school in London, Sept 2023

---

# [Task description and audit questions](https://github.com/01-edu/public/tree/master/subjects/forum)

---

# Demo 
[weWillHostSoon.com](https://weWillHostSoon.com/)

---
The code in this repo is meant to be a starting point for anyone building a simple forum application with Golang and SQL

## Prerequisites
You need to have [Go](https://golang.org/dl/) installed on your computer. The
version used to test the code in the tutorial is **1.19**.
## Natively (dev)
requirements: Golang v. 1.19, Node.js, Docker

## Get started

- Clone or download this repository to your filesystem.

```bash
$ git clone https://github.com/NatashaBishop/forum-in-golang
```

- `cd` into the project directory
---

# How to run?
Docker compose: 

    Build the docker image.
        docker build -t forum .
    Start an app container
        docker run -dp 127.0.0.1:8080:8080 forum
    Connect to database
        docker exec -it <DOCKER_ID> sqlite3 forum.db

The -d flag (short for --detach) runs the container in the background. The -p flag (short for --publish) creates a port mapping between the host and the container. The -p flag takes a string value in the format of HOST:CONTAINER, where HOST is the address on the host, and CONTAINER is the port on the container. The command publishes the container's port 8080 to 127.0.0.1:8080 (localhost:8080) on the host. Without the port mapping, you wouldn't be able to access the application from the host.
DATABASE

#### Run command line script: 
