import { useState } from "react";
import { UserData } from "../models/user";
import Login from "./login";
import Register from "./register";
import { SelectionType } from "../common/utils";

interface SelectionProps {
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
}


const Selection = ({ updateUserData }: SelectionProps) => {
    const [selected, updateSelected] = useState<number>(SelectionType.unselected);

    const handleSelection = () => (
        <div className="flex items-center justify-center gap-10">
            <button
                className="h-20 w-72 bg-green-600 text-white text-xl"
                onClick={() => updateSelected(SelectionType.login)}
            >
                Log In
            </button>
            <button
                className="h-20 w-72 bg-blue-600 text-white text-xl"
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