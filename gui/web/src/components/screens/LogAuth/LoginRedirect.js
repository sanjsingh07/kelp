import React, { useEffect, useState } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { Redirect } from "react-router";
import config from "../../../custom_config_ui.json"

const auth0enabled = config.auth0_enabled;

const LoginRedirect = () => {
  const { isLoading,isAuthenticated, loginWithRedirect, getAccessTokenSilently } = useAuth0();
  // const [ accessToken, setaccesstoken ] = useState();
  // const [some, setIsLoading] = useState(false);

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

  // try{
  //     const accessToken =  getAccessTokenSilently();
  //     // localStorage.setItem("accessToken", accessToken);
  //     console.log(accessToken)
  //     accessToken.then(function(result) {
  //       localStorage.setItem("accessToken", result);
  //       // return <Redirect to={{pathname: '/home'}}/>;
  //    })
  // } 
  // catch (e) {
  //   console.log(e.message);
  // }

      //  const getUserMetadata = async () => {
    
      //   try {
      //     const accessToken = await getAccessTokenSilently();
      //     localStorage.setItem("accessToken", accessToken);
      //   } catch (e) {
      //     console.log(e.message);
      //   }
      // };
    
      // getUserMetadata();
  
  // if (isAuthenticated) {
  //   useEffect(() => {
  //     const getUserMetadata = async () => {
    
  //       try {
  //         const accessToken = await getAccessTokenSilently();
  //         localStorage.setItem("accessToken", accessToken);
  //       } catch (e) {
  //         console.log(e.message);
  //       }
  //     };
    
  //     getUserMetadata();
  //   }, []);

  //   return <Redirect to={{pathname: '/home'}}/>;
  // }

  // try {
  //   const accessToken = await getAccessTokenSilently();
  //   localStorage.setItem("accessToken", accessToken);

  //   return <Redirect to={{pathname: '/home'}}/>;
  // } catch (e) {
  //   console.log(e.message);
  // }

  // const accessToken = getAccessTokenSilently();
  // accessToken.then(function(result) {
  //   localStorage.setItem("accessToken", result);
  //   console.log(result);
  //   return <Redirect to={{pathname: '/'}}/>;
  // })


  // useEffect(() => {
  //   const getUserMetadata = async () => {
  //     try {
  //       const accessToken = getAccessTokenSilently();
  //       accessToken.then(function(result) {
  //         localStorage.setItem("accessToken", result);
  //         // window.location.reload();
  //       })
  //       if (isLoading) {
  //         return <></>;
  //       }
  //     } catch (e) {
  //       console.log(e.message);
  //     }
  //     // window.location.reload();
  //   };
  
  //   getUserMetadata();
  // }, []);

  if (isLoading) {
    return <></>;
  }

  return (
    <></>
  );
};

export default LoginRedirect;