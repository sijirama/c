import React from 'react';
import SVG from '/public/all.svg';
import SVG2 from '/public/home2.svg';
import Image from 'next/image';

/*
 <div className=" ">
<Image src={SVG2} alt="My SVG" className='mx-auto' />
</div>
<div className=" ">
<Image src={SVG} alt="My SVG" className='mx-auto' />
</div>

    */

export default function HomePageRight() {
    return (
        <section className={`hidden md:w-1/2 overflow-y-hidden max-h-full md:flex items-center justify-end text-black relative`}>
            <div className=" flex absolute">
                <div className=" ">
                    <Image src={SVG} alt="My SVG" className='mx-auto' />
                </div>
            </div>
        </section>
    );
}

