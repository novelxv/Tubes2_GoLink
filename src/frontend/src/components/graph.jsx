import { useRef, useEffect } from 'react';
import * as d3 from 'd3';

const Graph = ({node,link}) => {
    const svgRef = useRef(null);

    useEffect(() => {
        const data = { nodes: node, links: link };


        const width = 1000;
        const height = 800;


        const countNodesPerLevel = (nodes) => {
            const maxLevel = Math.max(...nodes.map(node => node.level));
            
            const nodesPerLevel = Array.from({ length: maxLevel + 1 }, () => 0);
            
            nodes.forEach(node => {
                nodesPerLevel[node.level]++;
            });
        
            return nodesPerLevel;
        };
    
        
        const calculateXPositions = (numElementsPerLevel, width) => {
            const matrix = [];
            for (let i = 0; i < numElementsPerLevel.length; i++) {
                const positions = [];
                const nodeWidth = width / (numElementsPerLevel[i] + 1);
                for (let j = 0; j < numElementsPerLevel[i]; j++) {
                    positions.push(nodeWidth * (j + 1));
                }
                matrix.push(positions);
            }
            return matrix;
        };
        
        const nodesPerLevel = countNodesPerLevel(node);
        const positionsMatrix = calculateXPositions(nodesPerLevel, width);
        
        const updateNodeXPositions = (nodes, positionsMatrix) => {
            // Update x positions for each node
            nodes.forEach(node => {
                const positions = positionsMatrix[node.level];
                if (positions && positions.length > 0) {
                    // Set x position to the first position in the array
                    node.x = positions[0];
                    // Remove the first position from the array
                    positions.splice(0, 1);
                }
            });
        };
        
    
        updateNodeXPositions(node,positionsMatrix);
        // Main color
        const startColor = '#FB7185' // rose-400
        const endColor = '#1CD0A1' // emerald-400

        // Color per level
        const colorScale = d3.scaleOrdinal()
            .domain(data.nodes.map(node => node.level))
            .range(['#9060ff', '#f6cc6e', '#47dcfc']) // violet, blue,yellow


        const svg = d3.select(svgRef.current)
        .attr('width', width)
        .attr('height', height);    

        // Split links
        function formatWikipediaUrl(url) {
            const lastSegment = url.split('/').pop();
            const formattedText = lastSegment.split('_').join(' ');
            return formattedText;
        }

        function handleClick(d) {
            const url = d.url; // Access the 'url' property directly
            if (url) {
                window.open(url, '_blank'); // Open the URL in a new tab
            }
            console.log(url)
        }

        // Draw links
        svg.selectAll('line')
            .data(data.links)
            .enter()
            .append('line')
            .attr('x1', d => getNodePosition(d.source).x)
            .attr('y1', d => getNodePosition(d.source).y)
            .attr('x2', d => getNodePosition(d.target).x)
            .attr('y2', d => getNodePosition(d.target).y)
            .style('stroke', 'black')
            .style('stroke-width', 2);

        // Draw nodes
        svg.selectAll('circle')
            .data(data.nodes)
            .enter()
            .append('circle')
            .attr('cx', d => d.x)
            .attr('cy', d => d.y)
            .attr('r', 15)
            .style('fill', d => {
                if (d.level === 0) return startColor; // start
                if (d.level === Math.max(...data.nodes.map(node => node.level))) return endColor; // end
                return colorScale(d.level % 3); // others
            })
            .on("click", (event, d) => {
                console.log("Clicked label data:", d); // Log the entire data object associated with the clicked element
            });


        // Add labels to nodes
        svg.selectAll('text')
            .data(data.nodes)
            .enter()
            .append('text')
            .attr('x', d => d.x)
            .attr('y', d => d.y + 3)
            .attr('text-anchor', 'middle')
            .style('font-size', '12px')
            .text(d => formatWikipediaUrl(d.url))
            .on("click", (event, d) => {handleClick(d)});

        // Function to get node position based on ID
        function getNodePosition(nodeId) {
            const node = data.nodes.find(node => node.id === nodeId);
            return { x: node.x, y: node.y };
        }
    }, [node,link]);

    return (
        <svg ref={svgRef}></svg>
    );
};

export default Graph;
