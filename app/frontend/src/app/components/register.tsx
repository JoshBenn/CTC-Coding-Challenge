import { useRef, useState } from "react";
import { InOut, NewRegistrationRequest, UserData } from "../models/user";
import { ApiPath, JsonComponent, SelectionType } from "../common/utils";

export interface RegisterProps {
    updateSelection: React.Dispatch<React.SetStateAction<number>>;
    updateUserData: React.Dispatch<React.SetStateAction<UserData | undefined>>;
};

const Register = ({ updateSelection, updateUserData }: RegisterProps) => {
    const [email, updateEmail] = useState<string>("");
    const [username, updateUsername] = useState<string>("");
    const [password, updatePassword] = useState<string>("");
    const [notification, updateNotification] = useState<string[]>([]);
    const timer = useRef<NodeJS.Timeout | null>(null);

    // Handles validation and registration attempt
    const handleRegister = () => {
        // Verify that both the email and password were not blank
        var errors: string[] = [];
        if (email.length === 0) {
            errors.push("Email cannot be blank.");
        }
        if (username.length === 0) {
            errors.push("Username cannot be blank");
        }
        if (password.length === 0) {
            errors.push("Password cannot be blank.");
        }

        // Display the notification if any errors exist
        if (errors.length !== 0) {
            updateNotification(errors);
            resetNotification();
            return;
        }

        // Sanitize the email to ensure no sql attacks
        // In production this would be more restrictive and the backend would also confirm that the email is a proper domain
        const sanitizeEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!sanitizeEmail.test(email)) {
            updateNotification(["Invalid email."]);
            updateEmail("");
            updateUsername("");
            updatePassword("");
            resetNotification();
            return;
        }
        const regRequest = NewRegistrationRequest({ Email: email, Username: username, Password: password });

        // Send the registration request
        fetch(`${ApiPath.backend}${ApiPath.login}`, {
            method: JsonComponent.put,
            headers: {
                [JsonComponent.contentType]: JsonComponent.applicationJson
            },
            body: JSON.stringify(regRequest),
        }).then(response => {
            // If the response is not okay
            if (!response.ok) {
                errors.push(response.statusText);
                updateNotification(errors);
                resetNotification();
                return;
            }

            return response.json;
        }).then(data => {
            if (!data) {
                console.log("Unsuccessful");
                return;
            }

            console.log("Succes:", data);

            // Send user to login page
            updateSelection(SelectionType.login);
        }).catch((error) => {
            errors.push("Error", error);
            console.error("Error:", error);
        });
    };

    // Resets the notifications
    const resetNotification = () => {
        if (timer.current) {
            clearTimeout(timer.current);
        }

        timer.current = setTimeout(() => {
            updateNotification([]);
        }, 5000);
    };

    return (
        <div className="flex flex-col justify-between">
            <div className="flex justify-between">
                <button
                    className="border-2 border-black hover:bg-slate-500 hover:text-white h-8 self-start rounded-xl w-14"
                    onClick={() => updateSelection(SelectionType.unselected)}
                >
                    Back
                </button>
                <p className="text-center text-3xl mb-10 text-blue-700 font-extrabold">
                    Register
                </p>
                <p></p>
                <p></p>
            </div>
            <div className="text-red-400 text-base text-wrap self-center flex flex-col">
                {notification.map((message, i) => (
                    <p key={i}>{message}</p>
                ))}
            </div>
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