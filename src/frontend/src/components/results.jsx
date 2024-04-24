import React from "react";
import Graph from "./graph";

const ResultWrapper = ({ responseData }) => {
    // Initialize an empty array to store nodes
    const newNodes = [];
    // Initialize an empty array to store links
    const newLinks = [];
    const maxLength = Math.max(...responseData.articles.map(level => level.length));
    
    // Create nodes
    responseData.articles.forEach((level, levelIndex) => {
        level.forEach((url, index) => {
            // Check if the node already exists
            const existingNode = newNodes.find(node => node.url === url);
            if (!existingNode) {
                // If the node doesn't exist, create a new node
                const nodeLevel = index === level.length - 1 ? maxLength - 1 : index;
                newNodes.push({
                    id: newNodes.length, 
                    url: url,
                    x: 100, // Initialize x-coordinate to 0
                    y: 100 + 100 * nodeLevel,
                    level: nodeLevel // Store the level information
                });
            }
        });
    });

    // Create links
    for (const connection of responseData.articles) {
        for (let i = 0; i < connection.length - 1; i++) {
            const sourceUrl = connection[i];
            const targetUrl = connection[i + 1];
            const sourceNode = newNodes.find(node => node.url === sourceUrl);
            const targetNode = newNodes.find(node => node.url === targetUrl);
            if (sourceNode && targetNode) {
                newLinks.push({
                    source: sourceNode.id,
                    target: targetNode.id
                });
            }
        }
    }

    return (
        <div>
            <p>Total Visited: {String(responseData.articlesVisited)}</p>
            <p>Total Searched: {String(responseData.articlesSearched)}</p>
            <p>Use Time Needed: {String(responseData.timeNeeded)}</p>
            <h2>Articles:</h2>
            <div>
                <h2>Nodes:</h2>
                <ul>
                    {newNodes.map((node, index) => (
                        <li key={index}>ID: {node.id}, Url: {node.url}, Level: {node.level}, X : {node.x},Y:{node.y}</li>
                    ))}
                </ul>
                <h2>Links:</h2>
                <ul>
                    {newLinks.map((link, index) => (
                        <li key={index}>Source: {link.source}, Target: {link.target}</li>
                    ))}
                </ul>
            </div>
            <div>
                <Graph node={newNodes} link={newLinks} />
            </div>
        </div>
    );
};

export default ResultWrapper;
