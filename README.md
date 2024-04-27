# Tubes2_GoLink
> A web based application for finding the nearest route between two Wikipedia articles using the **BFS (Breadth First Search) and IDS (Iterative Deepening Search)** algorithms."

## Table of Contents

- [Technologies Used](#technologies-used)
- [Program Requirements](#program-requirements)
- [Setup](#setup)
- [Algorithms](#algorithms)
- [Screenshots](#screenshots)
- [Authors](#authors)


## Technologies Used

- NextJs - next@14.0.1
- Tailwind CSS - 3.3.0
- Go - version go1.22.2
- GIN - version v1.9.1
- shdcn/ui
- particles.js


## Program Requirements
- Docker installed (optional)
- Go Language Support
- Next JS installed and supported
- npm package manager

## Setup

### To start clone this respository 
```bash
git clone 
```

### Using Docker
> [!IMPORTANT]
> Make sure you have docker installed in your system
1. Navigate to this project's directory
```bash
cd Tubes2_GoLink
cd src
```
2. Run the docker container
```bash
docker compose up
```
3. Once the docker container is up and runner, click the localhost:3000 link
4. To take down, the docker container, place this command on the terminal
```bash
docker compose down
```

### Without using Docker
> [!NOTE]
> You will need to install all of the dependencies first

1. Navigate to this project's directory
```bash
cd Tubes2_GoLink
cd src
```
2. To install the frontend dependencies, run this command
```bash
cd frontend
npm install
```
3. Navigate back into the 'src' folder
```bash
cd ..
```
4. To install the backend dependencies, run this command
```bash
cd backend
go mod download
```
> [!NOTE]
> Once you have installed all of the dependencies, proceed to run the web application
1. To run the frontend, navigate into the frontend directory
```bash
cd frontend
npm run dev
```
2. Make sure you are back into the 'src' folder
```bash
cd ..
```
3. To run the backend, navigate into the backend directory
```bash
cd backend
go run server.go
```

## Algorithms

### BFS (Breadth First Search) Algorithm

The Breadth-First Search (BFS) algorithm is a technique used for traversing or searching tree or graph data structures. It starts from a selected node (often called the "source" or "root" node) and **explores all of its neighboring nodes** at the **present depth level** before moving on to the nodes at the next depth level.

Here is an overview of how BFS works:
1. Start with a **queue data structure** and enqueue the starting node.
2. Dequeue a node from the queue and mark it as visited.
3. Explore all unvisited neighboring nodes of the dequeued node and enqueue them into the queue.
4. Repeat steps 2 and 3 until the goal node is found (the node has been expanded)


### IDS (Iterative Depeening Search) Algorithm
The Iterative Deepening Search (IDS) algorithm is a **combination of Depth-First Search (DFS) and Breadth-First Search (BFS)**. It's used for traversing or searching tree or graph data structures

Here is an overview of how IDS works:
1. Start with a maximum depth limit set to 0.
2. Perform DFS with a depth limit of 0 (i.e., only explore nodes at depth 0).
3. If the goal node is not found at depth 0, increment the maximum depth limit by 1 and repeat step 2.
4. **Continue incrementing the depth limit** and performing DFS until the goal node is found or until all nodes have been explored.


## Screenshots
![localhost_3000-GoogleChrome2024-04-2706-34-4811-ezgif com-speed](https://github.com/novelxv/Tubes2_GoLink/assets/118401646/a00ca9b0-6e22-4875-88cb-96d3e2d0beac)
_Basic Overview of the Website_

![Graph Visualisation](https://github.com/novelxv/Tubes2_GoLink/assets/118401646/20044d30-1dfc-4c15-9092-37c372d0ef95)
_Multiple solutions with graph visualisation interface_


## Authors

| Name                            | GitHub                                           | NIM      |  Contact                     |
| ------------------------------ | ------------------------------------------------- | -------- | ---------------------------- |
| Debrina Veisha Rashika W       | [debrinashika](https://github.com/debrinashika)   | 13522025 | 13522025@std.stei.itb.ac.id  |
| Angelica Kierra Ninta Gurning  | [angiekierra](https://github.com/angiekierra)     | 13522048 | 13522048@std.stei.itb.ac.id  |
| Novelya Putri Ramadhani        | [novelxv](https://github.com/novelxv)             | 13522096 | 13522096@std.stei.itb.ac.id  |
