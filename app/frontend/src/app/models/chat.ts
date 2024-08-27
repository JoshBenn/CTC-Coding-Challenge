
export type Message = {
    Username: string;
    Content: string;
};

export const NewMessageRequest = ({Username: User, Content}: Message) => {
    return {
        username: User,
        content: Content,
    };
};