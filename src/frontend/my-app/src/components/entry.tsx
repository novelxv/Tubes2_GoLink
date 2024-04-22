import React, { useState } from 'react';
import Image from 'next/image';
import { Input } from '@/components/ui/input';
import { Switch } from "@/components/ui/switch"
import { Button } from './ui/button';
import axios from 'axios';
import Loading from './loading';

const Entry = () => {
    const [startLink, setStartLink] = useState('');
    const [endLink, setEndLink] = useState('');
    const [useToggle, setUseToggle] = useState(false);
    const [responseData, setResponseData] = useState<any>(null);
    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setLoading(true);
        try {
            const response = await axios.post('http://localhost:8080/api/input', {
                startLink: startLink,
                endLink: endLink,
                useToggle: useToggle
            });

            setResponseData(response.data);
            // console.log('Server response: ', response.data);
        } catch (error) {
            console.error('Error sending the data', error);
        } finally {
            setLoading(false);
        }
    };

    const handleSwitchChange = () => {
        setUseToggle(!useToggle);
    };

    return (
        <div className='bg-neutral-800 flex flex-col items-center justify-center min-h-screen relative'>
            <div className="absolute top-0 left-0 w-[571px] h-[442px] bg-emerald-400 rounded-full blur-[200px] -z-5"></div>
            <div className="absolute top-1/4 right-1/4 w-[600px] h-[363px] bg-violet-700 rounded-full blur-[150px] -z-6"></div>
            <div className="absolute top-1/2 right-1/2 w-[469px] h-[363px] bg-rose-400 rounded-full blur-[200px] -z-7"></div>
            <div className='z-0'>
                <Image src="/images/logo.png" alt='logo' width={872} height={165} />
                <form onSubmit={handleSubmit}>
                    <div className='flex flex-col items-center justify-center font-raleway text-neutral-100 p-7'>
                        <p className='flex-auto text-2xl xl:text-3xl'>Find the <b>shortest path</b> from</p>
                        <div className='flex flex-col xl:flex-row xl:gap-1  items-center' >
                            <div className='p-5'>
                                <Input 
                                    className='font-raleway text-neutral-400' 
                                    placeholder='Start Link'
                                    value={startLink}
                                    onChange={(e) => setStartLink(e.target.value)} 
                                />
                            </div>
                            <p className='text-lg'>to</p>
                            <div className='p-5'>
                                <Input 
                                    className='font-raleway text-neutral-400' 
                                    placeholder='End Link'
                                    value={endLink}
                                    onChange={(e) => setEndLink(e.target.value)} 
                                />
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
                        <Button type='submit' variant="default">Search Now</Button>
                    </div>
                </form>
                {loading && (
                    <div>
                        <p className='text-neutral-100 text-xl'>Loading...</p>
                    </div>
                )}
                {responseData && (
                    <div>
                        <p>Start Link: {responseData.startLink}</p>
                        <p>End Link: {responseData.endLink}</p>
                        <p>Use Toggle: {String(responseData.useToggle)}</p>
                    </div>
                )}
                {/* <div className='z-0'>
                    <Loading/>
                </div> */}
                
            </div>
        </div>
    );
};

export default Entry;
