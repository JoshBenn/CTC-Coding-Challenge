import { useState } from "react";
import { UserData } from "../models/user";
import Selection from "./selection";
import Chatroom from "./chatroom";

// Props for the content of the main page
interface ContentProps {
    userData: UserData | undefined;
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
}

// For the main content section of the page
const Content = ({ userData, updateUserData }: ContentProps) => {
    // maintains the content on the page
    const content = () => {
        // userData set here for testing without requiring login
        if (userData === undefined) {
            return (
                <Selection updateUserData={updateUserData} />
            );
        } else {
            return (
                <Chatroom user={userData} />
            );
        }

    };

    return (
        <div className="flex flex-col p-4 w-full h-full content-center justify-center bg-slate-100">
            {content()}
        </div>
    );
};

export default Content;