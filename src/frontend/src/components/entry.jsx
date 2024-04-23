import React, { useState, useRef } from 'react';
import Image from 'next/image';
import { Input } from '@/components/ui/input';
import { Switch } from "@/components/ui/switch"
import { Button } from './ui/button';
import axios from 'axios';
import Loading from './loading';
import ResultWrapper from './results';

const Entry = () => {
    const [startLink, setStartLink] = useState('');
    const [endLink, setEndLink] = useState('');
    const [useToggle, setUseToggle] = useState(false);
    const [responseData, setResponseData] = useState(null);
    const [loading, setLoading] = useState(false);
    const [focusOnStart, setFocusOnStart] = useState(true);
    const [startLinkSuggestions, setStartLinkSuggestions] = useState([]);
    const [endLinkSuggestions, setEndLinkSuggestions] = useState([]);
    const endLinkRef = useRef(null);
    const startLinkRef = useRef(null);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        try {
            const response = await axios.post('http://localhost:8080/api/input', {
                startLink: startLink,
                endLink: endLink,
                useToggle: useToggle
            });

            const responseData  = response.data;
            setResponseData(responseData);
        } catch (error) {
            console.error('Error sending the data', error);
        } finally {
            setLoading(false);
        }
    };

    const handleSwitchChange = () => {
        setUseToggle(!useToggle);
    };

    const switchText = () => {
        const temp = startLink;
        setStartLink(endLink);
        setEndLink(temp);
        if (focusOnStart) {
            startLinkRef.current.focus();
        } else {
            endLinkRef.current.focus();
        }
        setFocusOnStart(!focusOnStart);
    };

    const fetchSuggestions = async (input, setSuggestions, limit = 6) => {
        try {
            const response = await fetch(
                `https://en.wikipedia.org/w/api.php?action=opensearch&limit=10&format=json&search=${input}&origin=*`
            );
            const data = await response.json();
            const suggestions = data[1] || [];
            setSuggestions(suggestions.slice(0,limit));
        } catch (error) {
            console.error('Error fetching suggestions:', error);
        }
    };

    const handleStartLinkChange = (e) => {
        const value = e.target.value;
        setStartLink(value);
        fetchSuggestions(value, setStartLinkSuggestions,6);
    };

    const handleEndLinkChange = (e) => {
        const value = e.target.value;
        setEndLink(value);
        fetchSuggestions(value, setEndLinkSuggestions,6);
    };

    const handleStartLinkSuggestionClick = (suggestion) => {
        setStartLink(suggestion);
        setStartLinkSuggestions([]);
    };

    const handleEndLinkSuggestionClick = (suggestion) => {
        setEndLink(suggestion);
        setEndLinkSuggestions([]);
    };

    return (
        <div className='bg-neutral-800 flex flex-col items-center justify-center min-h-screen relative'>
            <div className="absolute top-0 left-0 w-[571px] h-[442px] bg-emerald-400 rounded-full blur-[200px] -z-5"></div>
            <div className="absolute top-1/4 right-1/4 w-[600px] h-[363px] bg-violet-700 rounded-full blur-[150px] -z-6"></div>
            <div className="absolute top-1/2 right-1/2 w-[469px] h-[363px] bg-rose-400 rounded-full blur-[200px] -z-7"></div>
            <div className='z-0'>
                <Image src="/logo.png" alt='logo' width={872} height={165} />
                <form onSubmit={handleSubmit}>
                    <div className='flex flex-col items-center justify-center font-raleway text-neutral-100 p-7'>
                        <p className='flex-auto text-2xl xl:text-3xl'>Find the <b>shortest path</b> from</p>
                        <div className='flex flex-col xl:flex-row xl:gap-1  items-center' >
                            <div className='p-5 relative'>
                                <Input
                                    ref={startLinkRef}
                                    className='font-raleway text-neutral-800' 
                                    placeholder='Start Link'
                                    value={startLink}
                                    onChange={handleStartLinkChange} 
                                />
                                <div className='absolute top-full bg-white w-72 shadow-md rounded-md z-10'>
                                    {startLinkSuggestions.map((suggestion, index) => (
                                        <div
                                            key={index}
                                            className='p-2 cursor-pointer font-raleway text-neutral-800/50 hover:bg-yellow-400/20 hover:text-neutral-800/70'
                                            onClick={() => handleStartLinkSuggestionClick(suggestion)}
                                        >
                                            {suggestion}
                                        </div>
                                    ))}
                                </div>
                            </div>
                            <Image src="/switch.svg" alt="switch" onClick={switchText} className='hover:animate-pulse' width={25} height={25}/>
                            <div className='p-5 relative'>
                                <Input 
                                    ref={endLinkRef}
                                    className='font-raleway text-neutral-800 ' 
                                    placeholder='End Link'
                                    value={endLink}
                                    onChange={handleEndLinkChange} 
                                />
                                <div className='absolute top-full bg-white w-72 shadow-md rounded-md z-10'>
                                    {endLinkSuggestions.map((suggestion, index) => (
                                        <div
                                            key={index}
                                            className='p-2 cursor-pointer font-raleway text-neutral-800/50 hover:bg-yellow-400/20 hover:text-neutral-800/70'
                                            onClick={() => handleEndLinkSuggestionClick(suggestion)}
                                        >
                                            {suggestion}
                                        </div>
                                    ))}
                                </div>
                            </div>
                        </div>
                        <div className='flex flex-col gap-3'>
                            <p className='text-center text-xl'>using</p>
                            <div className="flex flex-row gap-3 pb-5">
                                <p className="text-neutral-100 text-lg"><b>IDS</b></p>
                                <Switch
                                    checked={useToggle}
                                    onCheckedChange={handleSwitchChange}
                                />
                                <p className="text-neutral-100 text-lg"><b>BFS</b></p>
                            </div>
                        </div>
                        <Button type='submit' variant="default" className="hover:translate-y-2 transition-transform duration-300">Search Now</Button>
                    </div>
                </form>
                {loading && (
                    <Loading/>
                )}
                {responseData && (
                    <ResultWrapper
                        responseData={responseData}
                    />
                )}
            </div>
        </div>
    );
};

export default Entry;
