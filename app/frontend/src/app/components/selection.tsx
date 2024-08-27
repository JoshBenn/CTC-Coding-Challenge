import { useState } from "react";
import { UserData } from "../models/user";
import Login from "./login";
import Register from "./register";
import { SelectionType } from "../common/utils";

// Contains the elements passed into the selection page
interface SelectionProps {
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
}

// The selection page for logging in or registering
const Selection = ({ updateUserData }: SelectionProps) => {
    const [selected, updateSelected] = useState<number>(SelectionType.unselected);

    // Handles the button selected
    const handleSelection = () => (
        <div className="flex items-center justify-center gap-10">
            <button
                className="h-20 w-72 bg-green-600 text-white text-xl rounded-xl hover:bg-slate-500"
                onClick={() => updateSelected(SelectionType.login)}
            >
                Log In
            </button>
            <button
                className="h-20 w-72 bg-blue-600 text-white text-xl rounded-xl hover:bg-slate-500"
                onClick={() => updateSelected(SelectionType.register)}
            >
                Register
            </button>
        </div>
    );

    return (
        <div className="flex h-full w-full items-center justify-center">
            {selected === SelectionType.unselected && handleSelection()}
            {selected === SelectionType.login && <Login updateSelection={updateSelected} updateUserData={updateUserData} />}
            {selected === SelectionType.register && <Register  updateSelection={updateSelected} updateUserData={updateUserData} />}
        </div>
    );
};

export default Selection