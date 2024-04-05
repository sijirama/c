'use client'

import { AddReview } from "@/components/modals/CreateReview";
import { useEffect, useState } from "react";
export function ModalProvider() {
    const [isMounted, setIsMounted] = useState(false);

    useEffect(() => {
        setIsMounted(true);
    }, []);

    if (!isMounted) {
        return null;
    }

    return (
        <>
            <AddReview />
        </>
    )

}
