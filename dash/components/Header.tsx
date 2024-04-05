
import React, { HTMLAttributes } from 'react'
import Logo from './Logo'
import { cn } from '@/lib/utils'

interface Header extends HTMLAttributes<HTMLDivElement> {
    showInput?: boolean
}

export default function Header({ showInput, className, ...props }: Header) {
    const imageUrl = "https://avatars.githubusercontent.com/u/106167719?v=4"
    return (
        <nav {...props} className={cn('w-full h-16 bg-white ', className)}>
            <div className='w-10/12 mx-auto flex items-center h-full justify-between'>
                <div className='w-full h-full flex items-center gap-5'>
                    <Logo />
                    {
                        showInput ? (
                            <div className={`w-1/3 h-2/3 bg-black hidden md:block`}></div>
                        ) : (null)
                    }
                </div>
                <div className='flex items-center justify-center gap-2 '>
                    <p className='text-black dark:text-red-300 text-sm font-semibold'>Welcome!</p>
                    <div className="cursor-pointer rounded-full w-7 h-7 bg-center bg-cover bg-no-repeat" style={{ backgroundImage: `url(${imageUrl})` }}></div>
                </div>
            </div>
        </nav>
    )
}
