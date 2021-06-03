import getUserData from "./getUserData";

export default (baseUrl) => {
<<<<<<< HEAD
    return fetch(baseUrl + "/api/v1/genBotName",{
        method: "GET"
        }).then(resp => {
        return resp.text();
=======
    return fetch(baseUrl + "/api/v1/genBotName", {
        method: "POST",
        body: JSON.stringify({
            user_data: getUserData(),
        }),
    }).then(resp => {
        return resp.jtext();
>>>>>>> master
    });
};