import React from 'react';
import Link from 'next/link';
import Image from 'next/image';
import { Drawer, DrawerContent, DrawerTrigger } from "@/components/ui/drawer";

const Navbar = () => {
    return (
        <nav className="p-4 font-raleway absolute top-0 right-0"> 
            <div className="flex">
                <div className="flex space-x-4">
                    <Drawer>
                        <DrawerTrigger asChild>
                            <Link href="/" className="text-white hover:text-neutral-100 hover:underline">Author</Link>
                        </DrawerTrigger>
                        <DrawerContent>
                            <div className='font-raleway text-neutral-800'>
                                <div className='flex flex-row p-5 items-center gap-10 justify-center'>
                                    <div className='flex flex-col items-center justify-center'>
                                        <div className='rounded-full overflow-hidden'>
                                            <Image src="/download.jpg" alt='logo' width={100} height={100} />
                                        </div>
                                        <p className='font-semibold text-l xl:text-xl '>Debrina Veisha Rashika W</p>
                                        <p className='text-m xl:text-l'>13522025</p>
                                    </div>
                                    <div className='flex flex-col items-center justify-center'>
                                        <div className='rounded-full overflow-hidden'>
                                            <Image src="/download.jpg" alt='logo' width={100} height={100} />
                                        </div>
                                        <p className='font-semibold flex-auto text-l xl:text-xl p-1'>Angelica Kierra Ninta Gurning</p>
                                        <p className='flex-auto text-m xl:text-l'>13522048</p>    
                                    </div>
                                    <div className='flex flex-col items-center justify-center'>
                                        <div className='rounded-full overflow-hidden'>
                                            <Image src="/download.jpg" alt='logo' width={100} height={100} />
                                        </div>
                                        <p className='font-semibold text-l xl:text-xl p-1'>Novelya Putri Ramadhani</p>
                                        <p className='text-m xl:text-l'>13522096</p>
                                    </div>
                                </div>
                            </div>
                        </DrawerContent>
                    </Drawer>
                    <Link href="/about" className="text-white hover:text-neutral-100 hover:underline">How to Use</Link>
                    <a href="https://github.com/novelxv/Tubes2_GoLink" target="_blank" className="text-white hover:text-neutral-100 hover:underline">Github</a>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;
