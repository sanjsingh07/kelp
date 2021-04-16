import React, { useEffect } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import config from "../../custom_config_ui.json";
import { withAuthenticationRequired } from "@auth0/auth0-react";

const LoginRedirect = () => {
  const { getAccessTokenSilently } = useAuth0();
  const { isLoading,isAuthenticated, loginWithRedirect } = useAuth0();
  if (isLoading) {
    return <></>;
  }

  if (!isAuthenticated) {
    return loginWithRedirect();
  }

  if (isLoading) {
    return <></>;
  }

  useEffect(() => {
    const getUserMetadata = async () => {
  
      try {
        const accessToken = await getAccessTokenSilently();
        localStorage.setItem("accessToken", accessToken);
      } catch (e) {
        console.log(e.message);
      }
    };
  
    getUserMetadata();
  }, [isLoading, isAuthenticated, loginWithRedirect]);

  if (isLoading) {
    return <></>;
  }

  return (
      <>
      </>
  );
};

export default LoginRedirect;
// const customExportProfile = config.auth0_enabled ? withAuthenticationRequired(Profile) : Profile;
// export default customExportProfile;
