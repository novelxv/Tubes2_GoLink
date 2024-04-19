import React from 'react';
import Image from 'next/image';
import InputField from './inputfield';

const Entry = () => {
    return (
        <div className='bg-neutral-800 flex items-center justify-center h-screen'>
            <Image src="/images/logo.png" alt='logo' width={872} height={165} />
            <div className='font-raleway text-neutral-100'>
                <p className='text-3xl'>Find the <b>shortest path</b> from</p>
                <div>
                    <InputField/>
                    <p className=''>to</p>
                    <InputField/>
                </div>
                <div>
                    <p>using</p>
                    
                </div>
            </div>
        </div>
    );
};

export default Entry;