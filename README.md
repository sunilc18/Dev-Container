# Install Dev Container as an Developer Environment

## Overview

A Dev Container, or Development Container, is a feature provided by Visual Studio Code that allows developers to define their development environment as code. This is done using a combination of a Dockerfile and a .devcontainer/devcontainer.json file. The Dockerfile specifies the base image and any necessary dependencies, while the devcontainer.json file configures the way VS Code should behave once the environment is running.
The main technology behind Dev Containers is Docker, which provides the containerization. VS Code's Remote - Containers extension is used to interact with the Dev Container.
Dev Containers are particularly useful for teams of developers working on the same project, as they ensure everyone has the same development environment. This helps to eliminate the "it works on my machine" problem. They're also useful for open source projects, as they allow anyone to quickly set up a pre-configured environment.
In addition, Dev Containers can be beneficial for developers working across multiple projects with different environment requirements, as each project can have its own isolated environment.

## Getting Started

### Prerequisites

To use Dev Containers in Visual Studio Code, you'll need the following:

Visual Studio Code: This is the IDE that supports Dev Containers. You can download it from the official website.
Remote - Containers extension: This is a VS Code extension that enables the use of Dev Containers. You can install it from the VS Code marketplace.
Docker: Dev Containers rely on Docker to create and manage containers. You'll need to have Docker installed and running on your machine. You can download Docker from the official website.
A codebase with a .devcontainer folder: This folder, at the root of your project, should contain a devcontainer.json file and optionally a Dockerfile. These files define the Dev Container.
Adequate system resources: Running containers locally requires a fair amount of system resources. Ensure your machine has enough CPU, memory, and disk space to handle the containers you plan to use.
Operating System Support: Docker and VS Code are cross-platform, so you can use Dev Containers on Windows, Mac, and Linux. However, be aware that Docker has specific system requirements for each platform.

### Using the Dev Container

1. Install the [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension in VS Code.
2. Clone this repository.
3. Open the project in VS Code.
4. When prompted to reopen the project in a container, click "Reopen in Container".
5. Wait for the container to build. This may take a few minutes the first time.

## Using the Application

To start the Golang Web Application from this project run --> go run main.go

Once the app is started verify accesing the app from the browser localhost:8080
