"use client"
import React from 'react'
import { Button } from './ui/button'
import { IoSearch } from "react-icons/io5";
import { Inter } from 'next/font/google';
import { useRouter } from 'next/navigation';

export default function HomePageLeft() {
    const router = useRouter();
    
    const onClick = () => {
        router.push("/all")
    }

    return (
        <section className={`w-full md:w-1/2 h-full flex items-center justify-start text-black  `}>
            <div className=' w-5/6 md:w-2/3 mx-auto md:mx-0 py-2.5 md:py-4 space-y-4 md:space-y-6 text-center md:text-left '>
                <p className='text-3xl md:text-4xl lg:text-6xl -tracking-wide font-bold'>Find a place you will love to live!</p>
                <p className='font-xs text-sm md:text-xl'>See through the lenses of people who have
                    lived or visited the neighbourhood you might
                    have in mind.</p>
                <div className='bg-white p-1.5 md:p-3 flex items-center justify-center gap-2 rounded-lg border-[0.1px] border-zinc-300'>
                    <IoSearch className='text-xl' />
                    <input className='w-full outline-none focus:outline-none' type="text" placeholder="" />
                </div>
                <Button onClick={onClick} className='px-7 bg-[#3366FF] rounded-sm'>SEARCH</Button>
            </div>
        </section>
    )
}
