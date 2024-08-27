import { useEffect, useRef, useState } from "react";
import { UserData } from "../models/user";
import { ApiPath, JsonComponent } from "../common/utils";
import { Message, NewMessageRequest } from "../models/chat";

export type ChatroomProps = {
    user: UserData;
}

const Chatroom = ({ user }: ChatroomProps) => {
    // Maintains update timeout reference
    const update = useRef<NodeJS.Timeout | null>(null);
    // Update interval for getting messages from the server
    const interval = 100;

    // Contains all of the messages currently tracked in the chatroom (currently temporarily filled for testing)
    const [messages, updateMessages] = useState<Message[]>([
        // { User: "a", Content: "start" },
        // { User: "b", Content: "This" },
        // { User: "c", Content: "This" },
        // { User: "d", Content: "This" },
        // { User: "e", Content: "This" },
        // { User: "f", Content: "This" },
        // { User: "g", Content: "This" },
        // { User: "h", Content: "This" },
        // { User: "i", Content: "This" },
        // { User: "j", Content: "This" },
        // { User: "k", Content: "This" },
        // { User: "l", Content: "This" },
        // { User: "m", Content: "This" },
        // { User: "n", Content: "This" },
        // { User: "o", Content: "This" },
        // { User: "p", Content: "This" },
        // { User: "q", Content: "This" },
        // { User: "r", Content: "This" },
        // { User: "s", Content: "This" },
        // { User: "t", Content: "This" },
        // { User: "u", Content: "This" },
        // { User: "v", Content: "This" },
        // { User: "w", Content: "This" },
        // { User: "x", Content: "This" },
        // { User: "y", Content: "This" },
        // { User: "z", Content: "end" },
    ]);

    // Contains the current user message
    const [message, updateMessage] = useState<string>("");

    const resetInterval = () => {
        if (update.current) {
            clearInterval(update.current);
        }
        update.current = setInterval(() => {
            getMessages();
        }, interval);
    };

    useEffect(() => {
        resetInterval();

        return () => {
            if (update.current) {
                clearInterval(update.current);
            }
        };
    }, []);

    const getMessages = () => {
        if (update.current) {
            clearInterval(update.current);
        }

        // Send the reqeust
        fetch(`${ApiPath.backend}${ApiPath.chat}`, {
            method: JsonComponent.get,
            headers: {
                [JsonComponent.contentType]: JsonComponent.applicationJson
            },
        }).then(response => {
            if (!response.ok) {
                console.log(response);
                resetInterval();
                return;
            }

            return response.json()
        }).then(data => {
            if (!data) {
                resetInterval();
                console.log("Error in sending request", data);
                return;
            }

            let msgs: Message[] = [];
            for (let msg of data.messages) {
                msgs.push({ Username: msg.username, Content: msg.content });
            }
            updateMessages(msgs);
            console.log(msgs);

            // Reset the interval
            resetInterval();
        }).catch((error) => {
            console.error("Error:", error);
        });
    };

    const postMessage = () => {
        const messageRequest = NewMessageRequest({ Username: user.username, Content: message });
        updateMessage("");
        console.log(messageRequest);

        // Send the reqeust
        fetch(`${ApiPath.backend}${ApiPath.chat}`, {
            method: JsonComponent.post,
            headers: {
                [JsonComponent.contentType]: JsonComponent.applicationJson
            },
            body: JSON.stringify(messageRequest),
        }).then(response => {
            if (!response.ok) {
                console.log(response.text);
                return;
            }

            return response.json()
        }).then(data => {
            // Log for visibility
            console.log(data);

            // Get messages
            getMessages();
        }).catch((error) => {
            console.error("Error:", error);
        });
    };

    return (
        <div className="flex flex-col h-full w-full border-4 border-black bg-slate-700 rounded-2xl self-center justify-end text-white p-2"
            style={{
                background: "#323338"
            }}
        >
            <div className="scroll overflow-auto gap-2">
                {messages.map((message, i) => (
                    <div className="rounded-2xl hover:bg-slate-600 bg-[#2f3035]" key={i}>
                        <p className="px-4 font-bold" >{`${message.Username}:`}</p>
                        <p className="px-8 font-light" >{message.Content}</p>
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