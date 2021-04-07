import fetchIntercept from 'fetch-intercept';
import authConfig from '../../src/auth_config.json';

const AccessToken = authConfig.auth0enabled ? localStorage.getItem('accessToken') : null;

export const interceptor = fetchIntercept.register({
    request: function (url, config) {
        // Modify the url or config here
        const withDefaults = Object.assign({}, config);
        withDefaults.headers = config.headers || new Headers({
        'AUTHORIZATION': `Bearer ${AccessToken}`
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