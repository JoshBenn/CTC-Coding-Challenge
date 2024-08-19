import Image from "next/image";

const Header = () => {
    return (
        <div className="flex justify-between w-full">
            <Image
                className="w-24 h-auto"
                src="/mainctclogo.png"
                alt="CTC Main Logo"
                width={0}
                height={0}
                sizes="100vw"
            />
            <div className="border-2 border-black">

            </div>
        </div>
    );
};

export default Header;