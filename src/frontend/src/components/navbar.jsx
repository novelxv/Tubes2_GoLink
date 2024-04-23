import React from 'react';
import Link from 'next/link';

const Navbar = () => {
    return (
        <nav className="p-4 font-raleway absolute top-0 right-0"> 
            <div className="flex">
                <div className="flex space-x-4">
                    <Link href="/" className="text-white hover:text-neutral-100 hover:underline">Author</Link>
                    <Link href="/about" className="text-white hover:text-neutral-100 hover:underline">How to Use</Link>
                    <a href="https://github.com/novelxv/Tubes2_GoLink" target="_blank" className="text-white hover:text-neutral-100 hover:underline">Github</a>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;
