"use client"
import { cn } from "@/lib/utils"
import bg from "/public/logo.png"
import React, { HTMLAttributes } from 'react'
import { useRouter } from "next/navigation"

interface LogoProps extends HTMLAttributes<HTMLDivElement> { }

export default function Logo({ className, ...props }: LogoProps) {
    const router = useRouter()

    return (
        <div onClick={() => router.push("/")}  {...props} className={cn(`w-28 h-10 flex items-center justify-center`, className)}>
            <div className={`w-full h-full bg-contain bg-center bg-no-repeat`} style={{
                backgroundImage: `url(${bg.src})`,
            }} >
            </div>
        </div>
    )
}
