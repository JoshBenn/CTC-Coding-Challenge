import Image from "next/image";

const Header = (userData: any) => {
    /**src="/mainctclogo.png"*/
    return (
        <div className="flex justify-between w-full bg-sky-600 py-2 px-8" >
            <Image
                className="w-16 h-auto"
                src=""
                alt="CTC Main Logo"
                width={0}
                height={0}
                sizes="100vw"
            />
            <div className="border-2 w-24 border-black rounded-xl flex align-center justify-center p-4">
                {userData !== null ? (<p>Login</p>) : <p>asdf</p>}
            </div>
        </div>
    );
};

export default Header;