import React from "react";
import { useHistory } from "react-router-dom";
import { Auth0Provider } from "@auth0/auth0-react";
import authConfig from '../auth0-config.json';

const Auth0ProviderWithHistory = ({ children }) => {
  // const domain = process.env.REACT_APP_AUTH0_DOMAIN;
  // const clientId = process.env.REACT_APP_AUTH0_CLIENT_ID;

  const history = useHistory();

  const onRedirectCallback = (appState) => {
    history.push(appState?.returnTo || window.location.pathname);
  };

  return (
    <Auth0Provider
      domain= {authConfig.domain}
      clientId= {authConfig.client_id}
      audience= {authConfig.audience}
      redirectUri={window.location.origin}
      onRedirectCallback={onRedirectCallback}
    >
      {children}
    </Auth0Provider>
  );
};

export default Auth0ProviderWithHistory;