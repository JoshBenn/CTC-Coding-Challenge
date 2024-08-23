import { useState } from "react";
import { UserData } from "../models/user";

export type ChatroomProps = {
    user: UserData;
}

const Chatroom = ({ user }: ChatroomProps) => {
    const [messages, updateMessages] = useState<string[]>([])

    return (
        <div className="flex flex-col h-full w-full border-4 border-black bg-slate-700 rounded-2xl self-center">
            {messages.map((message, i) => (
                <p key={i} >{message}</p>
            ))}
        </div>
    );
};

export default Chatroom;