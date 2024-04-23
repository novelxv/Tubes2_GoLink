import { Target } from "lucide-react";
import React from "react";

const ResultWrapper = ({responseData}) => {

    const getNodeLinks = () => {
        const articles = responseData.articles;
        const nodes = articles.map((article) => ({ name: article }));
        const links = articles.slice(0,-1).map((article,index)=> ({
            source : index,
            target : index +1
        }));

        return {nodes,links};
    };
    
    const { nodes, links } = getNodeLinks();

    return (
        <div>
            <p>Total Visited: {String(responseData.articlesVisited)}</p>
            <p>Total Searched: {String(responseData.articlesSearched)}</p>
            <p>Use Time Needed: {String(responseData.useToggle)}</p>
            <h2>Articles:</h2>
            <div>
                <h2>Nodes:</h2>
                <ul>
                    {nodes.map((node, index) => (
                    <li key={index}>{node.name}</li>
                    ))}
                </ul>
                <h2>Links:</h2>
                <ul>
                    {links.map((link, index) => (
                    <li key={index}>Source: {link.source}, Target: {link.target}</li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default ResultWrapper;