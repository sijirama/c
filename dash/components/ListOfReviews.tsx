"use client"
import React from 'react'
import { Button } from './ui/button'
import { townReviews } from '@/lib/reviewData'
import { BiLike, BiDislike, BiComment } from "react-icons/bi";
import moment from "moment"
import { PiStarFill } from "react-icons/pi";

export default function ListOfReviews() {

    const onClick = () => {
    }

    return (
        <section className={`w-full md:w-2/3 lg:w-3/5 text-black flex flex-col gap-2 py-1  overflow-y-auto`}>
            {
                townReviews.map((review) => {
                    const time = new Date(review.timePosted)
                    let dateofpost = moment(time).fromNow()
                    return (
                        <div className='w-full p-3 space-y-3 border-b border-zinc-300 '>

                            <div className='w-full flex items-center justify-between'>
                                <div className='flex items-center gap-1.5'>
                                    <div className="cursor-pointer rounded-full w-7 h-7 bg-center bg-cover bg-no-repeat" style={{ backgroundImage: `url(${review.imageProfile})` }}></div>
                                    <p className='-tracking-wide font-semibold text-xs'>{review.name}</p>
                                    <p className='-tracking-wide font-xs text-xs text-gray-500'>{dateofpost}</p>
                                </div>

                                <div className='flex items-center text-sm gap-0.5'>
                                    <PiStarFill className='text-yellow-500 text-xl' />
                                    {review.starRating}
                                </div>
                            </div>

                            <div className='text-sm -tracking-wide'>
                                <p>{review.content}</p>
                            </div>

                            <div className='flex items-center gap-4 text-zinc-600 text-sm'>

                                <div className='flex items-center gap-0.5'>
                                    <BiLike className="" />
                                    <p>
                                        {review.likes}
                                    </p>
                                </div>

                                <div className='flex items-center gap-0.5'>
                                    <BiDislike className="" />
                                    <p>
                                        {review.dislikes}
                                    </p>
                                </div>

                                <div className='flex items-center gap-0.5'>
                                    <BiComment className="" />
                                    <p className='text-sm'>
                                        {review.comments}
                                    </p>
                                </div>


                            </div>

                        </div>
                    )
                })
            }
        </section>
    )
}
