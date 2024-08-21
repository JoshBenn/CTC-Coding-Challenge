import { useState } from "react";
import { UserData } from "../models/user";
import Login from "./login";
import Selection from "./selection";

// Props for the content of the main page
interface ContentProps {
    userData: UserData | undefined;
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
}

const Content = ({ userData, updateUserData }: ContentProps) => {
    const content = () => {
        if (userData === undefined) {
            return (
                <Selection updateUserData={updateUserData} />
            );
        }

    };
    return (
        <div className="flex flex-col my-10 w-3/4 h-full border-2 border-slate-600 rounded-3xl content-center justify-center">
            {content()}
        </div>
    );
};

export default Content;