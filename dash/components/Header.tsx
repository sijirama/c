"use client"
import React, { HTMLAttributes, useState } from 'react';
import Logo from './Logo';
import { cn } from '@/lib/utils';
import { IoSearch } from 'react-icons/io5';
import { IoMdClose } from 'react-icons/io';

interface Header extends HTMLAttributes<HTMLDivElement> {
    showInput?: boolean;
}

export default function Header({ showInput, className, ...props }: Header) {
    const [inputValue, setInputValue] = useState('');
    const imageUrl = 'https://avatars.githubusercontent.com/u/106167719?v=4';

    const handleInputChange = (e: any) => {
        setInputValue(e.target.value);
    };

    const handleClearInput = () => {
        setInputValue('');
    };

    return (
        <nav {...props} className={cn('w-full h-16 bg-white ', className)}>
            <div className="w-10/12 mx-auto flex items-center h-full justify-between">
                <div className="w-full h-full flex items-center gap-5">
                    <Logo />
                    {showInput ? (
                        <div className="bg-white p-1.5 w-3/5 md:p-2 hidden md:flex items-center justify-center gap-2 rounded-md border-[0.1px] border-zinc-300 ">
                            <IoSearch className="text-xl" />
                            <input
                                className="w-full outline-none focus:outline-none"
                                type="text"
                                placeholder=""
                                value={inputValue}
                                onChange={handleInputChange}
                            />
                            {(
                                <IoMdClose
                                    className="text-xl cursor-pointer"
                                    onClick={handleClearInput}
                                />
                            )}
                        </div>
                    ) : null}
                </div>
                <div className="flex items-center justify-center gap-2">
                    <p className="text-black dark:text-red-300 text-sm font-semibold">Welcome!</p>
                    <div
                        className="cursor-pointer rounded-full w-7 h-7 bg-center bg-cover bg-no-repeat"
                        style={{ backgroundImage: `url(${imageUrl})` }}
                    ></div>
                </div>
            </div>
        </nav>
    );
}

