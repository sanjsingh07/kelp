import React, { useEffect, useState } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import config from "../../auth_config.json";

const Profile = () => {
  const { getAccessTokenSilently } = useAuth0();
  const { isLoading,isAuthenticated, loginWithRedirect } = useAuth0();
  const audience = config.audience;
  const scope = config.scope;

  if (isLoading) {
    return <div></div>;
  }

  if (!isAuthenticated) {
    return loginWithRedirect();
  }

  if (isLoading) {
    return <div></div>;
  }

  useEffect(() => {
    const getUserMetadata = async () => {
  
      try {
        const accessToken = await getAccessTokenSilently({
          audience: audience,
          scope: scope,
        });
        localStorage.setItem("accessToken", accessToken);
      } catch (e) {
        console.log(e.message);
      }
    };
  
    getUserMetadata();
  }, [isLoading, isAuthenticated, loginWithRedirect]);

  return (
    
      <div>
      </div>
  );
};


export default Profile;