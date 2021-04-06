// This is used to determine if a user is authenticated and
// if they are allowed to visit the page they navigated to.

// If they are: they proceed to the page
// If not: they are redirected to the login page.
// import React from 'react'
// import {  Route } from 'react-router-dom'
// import { useAuth0, withAuthenticationRequired } from "@auth0/auth0-react";

// const PrivateRoute = ({ component, ...args}) => {
//     return(
//         <Route component={withAuthenticationRequired(component)} {...args} />
//     )

//   const {
//       isLoading,
//       error,
//       isAuthenticated,
//       loginWithRedirect,
//     } = useAuth0();
  
    // if (isLoading) {
    //   return <div>Loading...</div>;
    // }
    // if (error) {
    //   return <div>Oops... {error.message}</div>;
    // }

  
//   return (
//     <Route
//       {...args}
//       render={props =>
//         isAuthenticated ? (
//           <component {...props} />
//         ) : (
//           loginWithRedirect()
//         )
//       }
//     />
//   )
// }
// export default PrivateRoute;

import React, { useEffect } from 'react';
import { Route } from 'react-router-dom';
import { useAuth0 } from "@auth0/auth0-react";

const PrivateRoute = ({ component: Component, path, ...rest }) => {
  const { isLoading, isAuthenticated, loginWithRedirect, error } = useAuth0();

  if (isLoading) {
    return <div>Loading...</div>;
  }
  if (error) {
    return <div>Oops... {error.message}</div>;
  }

  useEffect(() => {
    if (isLoading || isAuthenticated) {
      return;
    }
    const fn = async () => {
      await loginWithRedirect({
        appState: { targetUrl: path }
      });
    };
    fn();
  }, [isLoading, isAuthenticated, loginWithRedirect, path]);

  const render = props =>
    isAuthenticated === true ? <Component {...props} /> : null;

  return <Route path={path} render={render} {...rest} />;
};

export default PrivateRoute;