import React from 'react'
import { Button } from './ui/button'
import { IoSearch } from "react-icons/io5";

export default function ImagesInReviewPage() {
    const landscapePictures = [
        "https://cdn.pixabay.com/photo/2023/12/22/19/56/hinduism-8464313_960_720.jpg",
        "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885_1280.jpg",
        "https://cdn.pixabay.com/photo/2016/11/14/04/45/elephants-1822636_1280.jpg",
        "https://cdn.pixabay.com/photo/2024/03/13/19/06/ai-generated-8631634_960_720.jpg"
    ];



    return (
        <section className={`hidden md:w-1/2 h-full md:flex justify-center gap-3 flex-wrap text-black content-start`}>
            {landscapePictures.map((picture, index) => {

                if (index === 3) {
                    return (

                        <div className='relative'>
                            <div className='bg-black bg-opacity-50 absolute inset-0 backdrop-blur-lg hover:backdrop-blur-md transition-all duration-400  z-10 flex justify-center items-center rounded-lg cursor-pointer'>
                                <div className="text-white text-center font-semibold">See More</div>
                            </div>
                            <div key={index} className='relative w-60 h-56 rounded-lg bg-center bg-cover bg-no-repeat' style={{ backgroundImage: `url(${picture})` }}></div>
                        </div>
                    )
                }

                return (
                    <div key={index} className='bg-black  w-60 h-56 rounded-lg bg-center bg-cover bg-no-repeat' style={{ backgroundImage: `url(${picture})` }}></div>
                )
            })}
        </section>
    )
}

