// Registration Request form
export type RegisterRequest = {
    email: string;
    username: string;
    password: string;
};

// Object for maintaining user data
export type UserData = {
    username: string;
    token: string;
    exp: number;
};