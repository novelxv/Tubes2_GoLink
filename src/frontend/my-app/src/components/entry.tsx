import React from 'react';
import Image from 'next/image';
import InputField from './inputField';
import Toggle from './toggle';
import { Button } from './ui/button';

const Entry = () => {
    return (
        <div className='bg-neutral-800 flex flex-col items-center justify-center h-screen '>
            <div className="absolute top-0 left-0 w-[571px] h-[442px] bg-emerald-400 rounded-full blur-[200px] -z-1"></div>
            <div className="absolute top-1/4 right-1/4 w-[600px] h-[363px] bg-violet-700 rounded-full blur-[150px] -z-2"></div>
            <div className="absolute top-1/2 right-1/2 w-[469px] h-[363px] bg-rose-400 rounded-full blur-[200px] -z-2"></div>
            <div className='z-0'>
                <Image src="/images/logo.png" alt='logo' width={872} height={165} />
                <div className='flex flex-col items-center justify-center font-raleway text-neutral-100 p-7'>
                    <p className='flex-auto text-2xl xl:text-3xl'>Find the <b>shortest path</b> from</p>
                    <div className='flex flex-col xl:flex-row xl:gap-1  items-center' >
                        <InputField placeholder='Start Link'/>
                        <p className='text-lg'>to</p>
                        <InputField placeholder='End Link'/>
                    </div>
                    <div className='flex flex-col gap-3'>
                        <p className='text-center text-xl'>using</p>
                        <Toggle/>
                    </div>
                    <Button variant="default">Search Now</Button>
                </div>
            </div>
        </div>
    );
};

export default Entry;