import React from "react";
import Graph from "./graph";

const ResultWrapper = ({ responseData }) => {
    const createNodesAndLinks = (articles) => {
        const newNodes = [];
        const newLinks = [];
        const maxLength = Math.max(...articles.map(level => level.length));
        
        articles.forEach((level, levelIndex) => {
            level.forEach((url, index) => {
                const existingNode = newNodes.find(node => node.url === url);
                if (!existingNode) {
                    const nodeLevel = index === level.length - 1 ? maxLength - 1 : index;
                    newNodes.push({
                        id: newNodes.length, 
                        url: url,
                        x: 0,
                        y: 100 + 80 * nodeLevel,
                        level: nodeLevel 
                    });
                }
            });
        });
    
        for (const article of articles) {
            for (let i = 0; i < article.length - 1; i++) {
                const sourceUrl = article[i];
                const targetUrl = article[i + 1];
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
    
        return { nodes: newNodes, links: newLinks };
    };

    const { nodes, links } = createNodesAndLinks(responseData.articles);
  

    return (
        <div>
            <p>Total Visited: {String(responseData.articlesVisited)}</p>
            <p>Total Searched: {String(responseData.articlesSearched)}</p>
            <p>Time Needed: {String(responseData.timeNeeded)}</p>
            <h2>Articles:</h2>
            <div>
                <h2>Nodes:</h2>
                <ul>
                    {nodes.map((node, index) => (
                        <li key={index}>ID: {node.id}, Url: {node.url}, Level: {node.level}, X : {node.x},Y:{node.y}</li>
                    ))}
                </ul>
                <h2>Links:</h2>
                <ul>
                    {links.map((link, index) => (
                        <li key={index}>Source: {link.source}, Target: {link.target}</li>
                    ))}
                </ul>
            </div>
            <div>
                <Graph node={nodes} link={links} />
            </div>
        </div>
    );
};

export default ResultWrapper;
