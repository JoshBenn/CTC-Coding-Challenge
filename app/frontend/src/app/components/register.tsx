import { useState } from "react";
import { UserData } from "../models/user";

interface RegisterProps {
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
};

const Register = ({ updateUserData }: RegisterProps) => {
    const [email, updateEmail] = useState<string>("");
    const [username, updateUsername] = useState<string>("");
    const [password, updatePassword] = useState<string>("");

    const handleRegister = () => {
        console.log(email, username, password);
    };

    return (
        <div className="flex flex-col justify-between">
            <p className="text-center text-3xl mb-10 text-blue-700 font-extrabold">
                Register
            </p>
            <div className="flex flex-col h-full w-full content-center justify-center text-black gap-3">
                <p className="text-black font-extrabold text-lg self-center -mb-2">Email</p>
                <input
                    className="bg-slate-400 h-10 w-72 p-2 rounded self-center text-white"
                    type="email"
                    value={email}
                    onChange={(e) => updateEmail(e.target.value)}
                    placeholder="Enter your email"
                />
                <p className="text-black font-extrabold text-lg self-center -mb-2">Username</p>
                <input
                    className="bg-slate-400 h-10 w-72 p-2 rounded self-center text-white"
                    type="email"
                    value={username}
                    onChange={(e) => updateUsername(e.target.value)}
                    placeholder="Enter your email"
                />
                <p className="text-black font-extrabold text-lg self-center -mb-2">Password</p>
                <input
                    className="bg-slate-400 h-10 w-72 p-2 rounded self-center"
                    type="password"
                    value={password}
                    onChange={(e) => updatePassword(e.target.value)}
                    placeholder="Enter your password"
                />
                <button
                    className="border-2 bg-purple-300 w-20 h-8 self-center rounded-xl hover:bg-blue-500"
                    onClick={handleRegister}
                >
                    Register
                </button>
            </div>
        </div>
    );
};

export default Register;