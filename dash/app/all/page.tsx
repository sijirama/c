"use client"
import Header from '@/components/Header'
import ListOfReviews from '@/components/ListOfReviews'
import { Button } from '@/components/ui/button'
import { facilitiesToConsider } from '@/lib/facilitiesList'
import React from 'react'
import { CiBookmark } from "react-icons/ci";
import { GoShareAndroid } from "react-icons/go";

export default function page() {
    const onClick = () => {
        console.log("Hire me")
    }
    return (
        <main className=" min-h-screen bg-zinc-50 flex flex-col">
            <section className='bg-slate-100 fixed'>
                <Header showInput={true} className='bg-slate-100' />
                <div className=' w-full md:w-10/12 mx-auto flex items-center py-5 h-full flex-col gap-5'>
                    <div className='flex items-center justify-between w-full'>
                        <div>
                            <p className='text-2xl leading-10 -tracking-wide font-semibold'>Bonny and Clyde Street, Ajao Estate, Lagos. </p>
                            <p>Reviews (People are raving about the selected location)</p>
                        </div>
                        <div className='flex items-center gap-2'>
                            <Button onClick={onClick} className='px-7 bg-[#3366FF] rounded-sm'>LEAVE A REVIEW</Button>
                            <Button onClick={onClick} className='px-4 bg-transparent border-2 border-[#3366FF] rounded-sm'>
                                <CiBookmark className='text-2xl font-semibold text-[#3366FF]' />
                            </Button>
                            <Button onClick={onClick} className='px-4 bg-transparent border-2 border-[#3366FF] rounded-sm'>
                                <GoShareAndroid className='text-2xl font-semibold text-[#3366FF]' />
                            </Button>
                        </div>
                    </div>
                    <div className='overflow-x-auto flex w-full gap-1'>
                        {
                            facilitiesToConsider.slice(0, 22).map((facility, index) => (
                                <div className='bg-white border min-w-fit border-zinc-300 p-1 px-1.5 rounded-md text-sm gap-1' key={index}>{facility}</div>
                            ))
                        }
                    </div>
                </div>
            </section>
            <section className="w-11/12 md:w-10/12 pt-60 mx-auto flex-1 flex overflow-y-auto">
                <div className="w-full h-full flex">
                    <ListOfReviews />
                </div>
            </section>

        </main>

    )
}
