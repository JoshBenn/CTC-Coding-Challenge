import { JsonComponent } from "../common/utils";

// Registration Request form
export type RegisterRequest = {
    Email: string;
    Username: string;
    Password: string;
};

export const NewRegistrationRequest = ({ Email, Username, Password }: RegisterRequest) => {
    return {
        [JsonComponent.email]: Email,
        [JsonComponent.username]: Username,
        [JsonComponent.password]: Password,
    };
}

// Object for maintaining user data
export type UserData = {
    username: string;
    token: string;
    exp: number;
};

export const NewUserData = ({ username, token, exp }: UserData) => {
    return {
        username: username,
        token: token,
        exp: exp,
    };
};

export const InOut = Object.freeze({
    Err: 0,
    In: 1,
    Out: 2,
});

// export class LoginReq {
//     in_out: number;
//     email: string;
//     password: string;

//     constructor(io: number, email: string, password: string) {
//         this.in_out = io;
//         this.email = email;
//         this.password = password;
//     }
// };

export type AuthenticationRequest = {
	InOut: number;    
	Email: string;    
	Password: string; 
};

export function NewAuthRequest({ InOut, Email, Password }: AuthenticationRequest) {
    return {
        [JsonComponent.inOut]: InOut,
        [JsonComponent.email]: Email,
        [JsonComponent.password]: Password,
    }
};
