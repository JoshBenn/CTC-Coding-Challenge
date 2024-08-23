
export type MessageProps = {
    Username: string;
    Message: string;
    Timestamp: string;
}

const Message = ({ Username, Message, Timestamp }: MessageProps) => {

    return (
        <div className="">
            <p>{Username} - {Timestamp}</p>
            <p>{Message}</p>
        </div>
    )
}