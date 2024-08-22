// Single point of truth for all json object components
export const JsonComponent = Object.freeze({
    inOut: "in_out",
    email: "email",
    username: "username",
    password: "password",
    exp: "exp",
    get: "GET",
    post: "POST",
    put: "PUT",
    success: "STATUS_SUCCESS",
    failure: "STATUS_FAILURE",
    contentType: "content-type",
    applicationJson: "application/json",
    setCookie: "Set-Cookie",
});

// Selection state for the component
export const SelectionType = Object.freeze({
    unselected: 0,
    login: 1,
    register: 2,
});

// Api paths
export const ApiPath = Object.freeze({
    //backend: "localhost:5150",
    backend: "/api",
    chat: "/chat",
    login: "/login"
});