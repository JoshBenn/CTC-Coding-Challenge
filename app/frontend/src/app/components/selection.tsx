import { useState } from "react";
import { UserData } from "../models/user";
import Login from "./login";
import Register from "./register";

interface SelectionProps {
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
}

// Selection state for the component
const selection = Object.freeze({
    unselected: 0,
    login: 1,
    register: 2,
})

const Selection: React.FC<SelectionProps> = ({ updateUserData }) => {
    const [selected, updateSelected] = useState<number>(selection.unselected);

    const handleSelection = () => (
        <div className="flex items-center justify-center gap-10">
            <button
                className="h-20 w-72 bg-green-600 text-white text-xl"
                onClick={() => updateSelected(selection.login)}
            >
                Log In
            </button>
            <button
                className="h-20 w-72 bg-blue-600 text-white text-xl"
                onClick={() => updateSelected(selection.register)}
            >
                Register
            </button>
        </div>
    );

    return (
        <div className="flex h-full w-full items-center justify-center">
            {selected === selection.unselected && handleSelection()}
            {selected === selection.login && <Login updateUserData={updateUserData} />}
            {selected === selection.register && <Register updateUserData={updateUserData} />}
        </div>
    );
};

export default Selection