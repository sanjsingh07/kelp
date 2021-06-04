import getUserData from "./getUserData";

export default (baseUrl) => {
    return fetch(baseUrl + "/api/v1/autogenerate", {
<<<<<<< HEAD
        method: "GET"
        }).then(resp => {
=======
        method: "POST",
        body: JSON.stringify({
            user_data: getUserData(),
        }),
    }).then(resp => {
>>>>>>> master
        return resp.json();
    });
};