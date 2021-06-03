import getUserData from "./getUserData";

export default (baseUrl) => {
<<<<<<< HEAD
    return fetch(baseUrl + "/api/v1/listBots",{
        method: "GET"
        }).then(resp => {
=======
    return fetch(baseUrl + "/api/v1/listBots", {
        method: "POST",
        body: JSON.stringify({
            user_data: getUserData(),
        }),
    }).then(resp => {
>>>>>>> master
        return resp.json();
    });
};