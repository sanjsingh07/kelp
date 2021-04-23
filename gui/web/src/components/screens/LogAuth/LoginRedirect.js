import React, { useEffect } from "react";
import { useAuth0 } from "@auth0/auth0-react";

const LoginRedirect = () => {
  const { isLoading,isAuthenticated, loginWithRedirect, getAccessTokenSilently } = useAuth0();

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
    const getaccesstoken = async () => {
      try {
        const accessToken = await getAccessTokenSilently();
        localStorage.setItem("accessToken", accessToken);
        if (isLoading) {
          return <></>;
        }
      } catch (e) {
        console.log(e.message);
      }
    };
  
    getaccesstoken();
  }, []);

  if (isLoading) {
    return <></>;
  }

  return (
    <></>
  );
};

export default LoginRedirect;