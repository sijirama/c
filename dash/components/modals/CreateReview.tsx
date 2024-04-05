import { Button } from "@/components/ui/button"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { useModal } from "@/store/ModalStore"

import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { facilitiesToConsider } from "@/lib/facilitiesList"
import { Textarea } from "@/components/ui/textarea"
import { Label } from "../ui/label"
import { useEffect, useRef } from "react"
import React, { useState } from 'react'
import Rating from '@mui/material/Rating';

export function AddReview() {

    const { type, isOpen, onClose } = useModal()
    const Open = type == "addReview" && isOpen
    const reff = useRef<HTMLDivElement>(null)

    useEffect(() => {
        reff.current?.focus()
    }, [Open])

    return (
        <Dialog open={Open} onOpenChange={onClose}  >
            <DialogContent className="sm:max-w-[625px] sm:min-h-96" ref={reff}>
                <DialogHeader>
                    <DialogTitle className="text-center">Review Location</DialogTitle>
                    <DialogDescription>
                        Add a personal review here to the specified location
                    </DialogDescription>
                </DialogHeader>
                <div className="w-full space-y-6">
                    <div className="">
                        <Select>
                            <SelectTrigger className="w-full" >
                                <SelectValue placeholder="Select Amenity" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectGroup>
                                    <SelectLabel>Amenities</SelectLabel>
                                    {
                                        facilitiesToConsider.slice(0, 10).map((facility, index) => (
                                            <SelectItem key={index} value={facility}>{facility}</SelectItem>
                                        ))
                                    }
                                </SelectGroup>
                            </SelectContent>
                        </Select>
                    </div>
                    <div className="flex flex-col gap-1.5">
                        <Label htmlFor="message">Rate Location</Label>
                        <Rating />
                    </div>
                    <div className="">
                        <Label htmlFor="message">Write Review</Label>
                        <Textarea placeholder="Type your review here." draggable={false} className="resize-none min-h-48" />
                    </div>
                    <RadioGroup>
                        <div className="flex items-center space-x-2">
                            <RadioGroupItem value="default" id="r1" />
                            <Label htmlFor="r1">Post as anonymous</Label>
                        </div>
                    </RadioGroup>
                </div>
                <DialogFooter className="w-full">
                    <Button disabled type="submit" className="w-1/2">Submit</Button>
                    <Button onClick={onClose} className='px-4 hover:bg-zinc-300 bg-transparent border-2 border-[#3366FF] text-[#3366FF] w-1/2 rounded-lg'>CANCEL</Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}

