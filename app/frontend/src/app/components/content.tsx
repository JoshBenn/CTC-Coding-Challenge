import { useState } from "react";
import { UserData } from "../models/user";
import Selection from "./selection";
import Chatroom from "./chatroom";

// Props for the content of the main page
interface ContentProps {
    userData: UserData | undefined;
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
}

const Content = ({ userData, updateUserData }: ContentProps) => {
    const content = () => {
        userData = { username: "test", token: "asdf", exp: 123 };
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
        <div className="flex flex-col my-10 w-3/4 h-4/5 border-2 border-slate-600 rounded-3xl content-center justify-center">
            {content()}
        </div>
    );
};

export default Content;