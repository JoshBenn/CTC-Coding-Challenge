import { useEffect, useState } from "react";
import { UserData } from "../models/user";
import { ApiPath, JsonComponent } from "../common/utils";
import { Message, NewMessageRequest } from "../models/chat";

export type ChatroomProps = {
    user: UserData;
}

const Chatroom = ({ user }: ChatroomProps) => {
    const [messages, updateMessages] = useState<Message[]>([
        { User: "a", Content: "start" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "This" },
        { User: "a", Content: "end" },
    ]);
    const [message, updateMessage] = useState<string>("");

    useEffect(() => {
        getMessages();
    }, []);

    const getMessages = () => {
        // Send the reqeust
        fetch(`${ApiPath.backend}${ApiPath.login}`, {
            method: JsonComponent.get,
            headers: {
                [JsonComponent.contentType]: JsonComponent.applicationJson
            },
        }).then(response => {
            if (!response.ok) {
                console.log(response);
                return;
            }

            return response.json()
        }).then(data => {
            updateMessages(data.messages);
            console.log(data);
        }).catch((error) => {
            console.error("Error:", error);
        });
    };

    const postMessage = () => {
        const messageRequest = NewMessageRequest({ User: user.username, Content: message });
        updateMessage("");

        // Send the reqeust
        fetch(`${ApiPath.backend}${ApiPath.login}`, {
            method: JsonComponent.post,
            headers: {
                [JsonComponent.contentType]: JsonComponent.applicationJson
            },
            body: JSON.stringify(messageRequest),
        }).then(response => {
            if (!response.ok) {
                console.log(response);
                return;
            }

            return response.json()
        }).then(data => {
            updateMessages(data.messages);
            console.log(data);
        }).catch((error) => {
            console.error("Error:", error);
        });

        getMessages();
    };

    return (
        <div className="flex flex-col h-full w-full border-4 border-black bg-slate-700 rounded-2xl self-center justify-end text-white p-2"
            style={{
                background: "#323338"
            }}
        >
            <div className="scroll overflow-auto gap-2">
                {messages.map((message, i) => (
                    <div className="rounded-2xl hover:bg-slate-600 bg-[#2f3035]">
                        <p className="px-4 font-bold" key={i} >{`${message.User}:`}</p>
                        <p className="px-8 font-light" key={i} >{message.Content}</p>
                    </div>
                ))}
            </div>
            <div className="flex gap-2 self-center w-full text-white relative" >
                <textarea
                    className="w-11/12 h-24 rounded-2xl text-wrap p-2 resize-none drop-shadow-xl border-2"
                    style={{ background: "#383a3f", }}
                    onChange={e => updateMessage(e.target.value)}
                    value={message}
                />
                <button
                    className="w-1/12 h-12 m-auto border-2 rounded-2xl hover:bg-slate-500"
                    onClick={() => postMessage()}
                >
                    Send
                </button>
            </div>
        </div>
    );
};

export default Chatroom;