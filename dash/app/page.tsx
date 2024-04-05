import Header from "@/components/Header";
import HomePageLeft from "@/components/HomePageLeft";
import HomePageRight from "@/components/HomePageRight";

export default function Home() {
    return (
        <main className=" min-h-screen bg-zinc-50 flex flex-col">
            <Header showInput={false} />
            <section className="w-11/12 md:w-10/12 mx-auto flex-1 flex">
                <div className="w-full flex">
                    <HomePageLeft />
                    <HomePageRight />
                </div>
            </section>
        </main>
    );
}
