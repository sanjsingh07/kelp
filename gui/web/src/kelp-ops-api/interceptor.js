import fetchIntercept from 'fetch-intercept';
import authConfig from '../auth_config.json';
import { useAuth0 } from "@auth0/auth0-react";

// const { getAccessTokenSilently } = useAuth0();

// (async () => {
//         try {
//           const AccessToken = await getAccessTokenSilently({
//             audience: authConfig.audience,
//             scope: authConfig.scope,
//           });
//         } catch (e) {
//           console.log(e.message);
//         }
// })();


export const interceptor = fetchIntercept.register({
    request: function (url, config) {
        // Modify the url or config here
        const withDefaults = Object.assign({}, config);
        withDefaults.headers = config.headers || new Headers({
        'AUTHORIZATION': `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InVVOUxUZk00YXNmNHVLdTh1VG5JRiJ9.eyJpc3MiOiJodHRwczovL2Rldi02ZXNnMGkwMi5ldS5hdXRoMC5jb20vIiwic3ViIjoiRWtkMmdYbU11SUZ4UzNwRkdwb2pLTEJrQ0dLakxUQWNAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vbG9jYWxob3N0OjMwMDAvYXBpL3YyLyIsImlhdCI6MTYxNzYyMjcwMywiZXhwIjoxNjE3NzA5MTAzLCJhenAiOiJFa2QyZ1htTXVJRnhTM3BGR3BvaktMQmtDR0tqTFRBYyIsImd0eSI6ImNsaWVudC1jcmVkZW50aWFscyJ9.dMCUkysTf3KDAfTucBYNlBxJym7Tv9-McQx0JhilGMgyTILy49F-x7AKUtn1zd9M9wKJtpkEAq7-0QQ8rMNQUjrTh_3uN90ioIpkWSA6mXbwCNu9OoqBUVitUOIZ0QdJeDtImrQb0d4oPdBTv-hkFOAd8ru1rl97GDLhatZilLhD7wdDef26EHAv_r6nAgRYdbHkbMBqARo636T2tXS2tFti_RZPlGnGsgffjVI3iBkYlDR6ZNOKjaQxEejJqOBpictusCHnk3X5Z73YASzsFsMl1DStNJFVKULpZn3Lkm7Iq7gQk8oHtdgRwIMe9T22vJHfCB8JShBhwvgcbA5Aow`
        });
    return [url, withDefaults]
    },
 
    requestError: function (error) {
        // Called when an error occured during another 'request' interceptor call
        return Promise.reject(error);
    },
 
    response: function (response) {
        // Modify the reponse object
        return response;
    },
 
    responseError: function (error) {
        // Handle an fetch error
        return Promise.reject(error);
    }
});