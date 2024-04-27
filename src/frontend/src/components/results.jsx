import React from "react";
import Graph from "./graph";

const ResultWrapper = ({ responseData }) => {

    // Function to create nodes and links from matrix of strings
    const createNodesAndLinks = (articles) => {
        const newNodes = [];
        const newLinks = [];
        // Get highest level
        const maxLength = Math.max(...articles.map(level => level.length));
        
        // Iterate through each element
        articles.forEach((level, levelIndex) => {
            level.forEach((url, index) => {
                // If is already a node, don't insert into newNodes
                const existingNode = newNodes.find(node => node.url === url);
                if (!existingNode) {
                    const nodeLevel = index === level.length - 1 ? maxLength - 1 : index;
                    newNodes.push({
                        id: newNodes.length, 
                        url: url,
                        x: 0, // will be updated in the graph
                        y: 100 + 70 * nodeLevel, // Distributing y position evenly
                        level: nodeLevel 
                    });
                }
            });
        });
        
        // Iterate through each articles matrix
        for (const article of articles) {

            // Always leaving the last element behind
            for (let i = 0; i < article.length - 1; i++) {
                const sourceUrl = article[i];
                const targetUrl = article[i + 1];

                // Look for the urls
                const sourceNode = newNodes.find(node => node.url === sourceUrl);
                const targetNode = newNodes.find(node => node.url === targetUrl);

                // Push into the list of links
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


    // Creating nodes and links
    const { nodes, links } = createNodesAndLinks(responseData.articles);
  

    return (
        <div className="font-raleway flex flex-col items-center justify-center">
            <div className="text-neutral-100 text-xl border rounded-md p-5 m-7">
                <p>Found <strong>{String(responseData.articles.length)}</strong> path in <strong>{String(responseData.timeNeeded)} ms </strong>with <strong>{String(responseData.articlesVisited)}</strong> Articles visited and <strong>{String(responseData.articlesSearched)}</strong> Articles searched</p>
            </div>
            <Graph node={nodes} link={links} />
        </div>
    );
};

export default ResultWrapper;