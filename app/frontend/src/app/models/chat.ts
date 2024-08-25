
export type Message = {
    User: string;
    Content: string;
};

export const NewMessageRequest = ({User, Content}: Message) => {
    return {
        user: User,
        content: Content,
    };
};