import Image from "next/image";
import { UserData } from "../models/user";
import { useState } from "react";

type HeaderProps = {
    userData: UserData | undefined;
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
}

const Header = ({ userData, updateUserData }: HeaderProps) => {
    const [showPopup, setShowPopup] = useState<boolean>(false);

    const handleButtonClick = () => {
        setShowPopup(!showPopup);
    };

    const handleLogout = () => {
        console.log("user logged out");
        updateUserData(undefined);
        setShowPopup(false);
    }

    return (
        <div className="flex justify-between w-full bg-sky-600 py-2 px-8" >
            {/** In practice this would be a link with an image */}
            <div className="text-wrap w-16" >Company Logo</div>
            <div className="relative">
                <button
                    onClick={handleButtonClick}
                    className="border-2 w-24 border-black rounded-xl flex items-center justify-center p-4 hover:text-white hover:bg-slate-500"
                >
                    {userData === undefined ? <p>User</p> : <p>{userData.username}</p>}
                </button>
                {showPopup && userData !== undefined && (
                    <div className="absolute mt-2 right-0 bg-gray-300 border-black border-2 rounded shadow-lg">
                        <button
                            onClick={handleLogout}
                            className="w-full text-left px-4 py-2 hover:bg-gray-200"
                        >
                            Logout
                        </button>
                    </div>
                )}
            </div>
            {/* <button
                className="border-2 w-24 border-black rounded-xl flex align-center justify-center p-4 hover:text-white hover:bg-slate-500"
            >
                {userData === undefined ? <p>User</p> : <p>{userData.username}</p>}
            </button> */}
        </div>
    );
};

export default Header;