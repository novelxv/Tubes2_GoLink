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
                        y: 100 + 70 * nodeLevel,
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
        <div className="font-raleway flex flex-col items-center justify-center">
            <div className="text-neutral-100 text-xl border rounded-md p-5 m-7">
                <p>Found <strong>{String(responseData.articles.length)}</strong> path in <strong>{String(responseData.timeNeeded)} ms </strong>with <strong>{String(responseData.articlesVisited)}</strong> Articles visited and <strong>{String(responseData.articlesSearched)}</strong> Articles searched</p>
            </div>
            <Graph node={nodes} link={links} />
        </div>
    );
};

export default ResultWrapper;