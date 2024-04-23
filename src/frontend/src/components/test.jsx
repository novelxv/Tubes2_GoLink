import React, { useEffect, useState } from 'react';

const Test = () => {
  const [nodes, setNodes] = useState([]);
  const [links, setLinks] = useState([]);

  useEffect(() => {
    const data = [
      ['A', 'B', 'C', 'D', 'E'],
      ['A', 'G', 'E'],
      ['A', 'L', 'E']
    ];

    // Initialize an empty array to store nodes
    const newNodes = [];
    // Initialize an empty array to store links
    const newLinks = [];

    // Create nodes
    data.forEach((level, levelIndex) => {
      level.forEach((name, index) => {
        // Check if the node already exists
        const existingNode = newNodes.find(node => node.name === name);
        if (!existingNode) {
          // If the node doesn't exist, create a new node
          newNodes.push({
            id: newNodes.length, // Assign a unique ID to each node
            name: name,
            x: 0, // Initialize x-coordinate to 0
            y: 0, // Initialize y-coordinate to 0
            level: index // Store the level information
          });
        }
      });
    });

    // Create links
    for (const connection of data) {
      for (let i = 0; i < connection.length - 1; i++) {
        const sourceName = connection[i];
        const targetName = connection[i + 1];
        const sourceNode = newNodes.find(node => node.name === sourceName);
        const targetNode = newNodes.find(node => node.name === targetName);
        newLinks.push({
          source: sourceNode.id,
          target: targetNode.id
        });
      }
    }

    // Set the nodes and links state
    setNodes(newNodes);
    setLinks(newLinks);
  }, []);

  return (
    <div>
      <h2>Nodes:</h2>
      <ul>
        {nodes.map(node => (
          <li key={node.id}>
            Name: {node.name}, ID: {node.id}, X: {node.x}, Y: {node.y}, Level: {node.level}
          </li>
        ))}
      </ul>
      <h2>Links:</h2>
      <ul>
        {links.map((link, index) => (
          <li key={index}>
            Source: {link.source}, Target: {link.target}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Test;
