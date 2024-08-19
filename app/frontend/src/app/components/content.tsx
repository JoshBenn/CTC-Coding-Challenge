
const Content = (userData: any) => {
    return (
        <div className="flex flex-col my-10 w-3/4 h-full border-2 border-slate-600 rounded-3xl">
            {
                userData === null ? 
                    (<div className="self-center">asdf</div>)
                    : (<div></div>)
            }
        </div>
    );
};

export default Content;