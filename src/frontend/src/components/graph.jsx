import { useRef, useEffect } from 'react';
import * as d3 from 'd3';

const Graph = () => {
  const svgRef = useRef(null);

  useEffect(() => {
    const data = {
      nodes: [
        { url: 'https://en.wikipedia.org/wiki/Joko_Widodo_2014_presidential_campaign', id: 0, x: 100, y: 100, level: 0 },
        { url: 'https://en.wikipedia.org/wiki/Fauzi_Bowo', id: 1, x: 50, y: 200, level: 1 },
        { url: 'https://en.wikipedia.org/wiki/Javanese_people', id: 2, x: 75, y: 300, level: 2 },
        { url: 'https://en.wikipedia.org/wiki/List_of_languages_by_number_of_native_speakers', id: 3, x: 100, y: 400, level: 3 },
        { url: 'https://en.wikipedia.org/wiki/Mandarin_Chinese', id: 4, x: 100, y: 500, level: 4 },
        { url: 'https://en.wikipedia.org/wiki/Historical_capitals_of_China', id: 5, x: 100, y: 200, level: 1 },
        { url: 'https://en.wikipedia.org/wiki/Anyang', id: 6, x: 150, y: 200, level: 1 },
        { url: 'https://en.wikipedia.org/wiki/Stone_Age', id: 7, x: 150, y: 300, level: 2 },
        { url: 'https://en.wikipedia.org/wiki/Stone_Age', id: 8, x: 100, y: 600, level: 5 }
      ],
      links: [
        { source: 0, target: 1 },
        { source: 1, target: 2 },
        { source: 2, target: 3 },
        { source: 3, target: 4 },
        { source: 0, target: 5 },
        { source: 5, target: 4 },
        { source: 0, target: 6 },
        { source: 6, target: 7 },
        { source: 7, target: 4 },
        { source: 4, target: 8 }

      ],
    };

    const width = 800;
    const height = 800;

    
    // Main color
    const startColor = '#FB7185' // rose-400
    const endColor = '#1CD0A1' // emerald-400

    // Color per level
    const colorScale = d3.scaleOrdinal()
        .domain(data.nodes.map(node => node.level))
        .range(['#6027CC', '#f6cc6e', '#47dcfc']) // violet, blue,yellow


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
        .text(d => formatWikipediaUrl(d.url))
        .on("click", (event, d) => {handleClick(d)});

    // Function to get node position based on ID
    function getNodePosition(nodeId) {
        const node = data.nodes.find(node => node.id === nodeId);
        return { x: node.x, y: node.y };
    }
    }, []);

    return (
        <svg ref={svgRef}></svg>
    );
};

export default Graph;
