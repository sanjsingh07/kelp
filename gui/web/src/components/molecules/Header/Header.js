import React, { Component } from 'react';
import logo from '../../../assets/images/orunpay_logo.png';
import grid from '../../_styles/grid.module.scss';
import styles from './Header.module.scss';
import LogoutButton from '../../screens/LogAuth/LogoutButton';
import config from "../../../custom_config_ui.json";

const auth0enabled = config.auth0_enabled;

class Header extends Component {
  render() {
    return (
      <header className={styles.header}>
        <div className={grid.container}>
          <div className={styles.headerWrapper}>
            <div className={styles.logoWrapper}>
              <img src={logo} className={styles.logo} alt="Orunpay logo" />
              {/* <span className={styles.version}>{this.props.version}</span> */}
            </div>
            <div>
            {auth0enabled ? (<LogoutButton className={styles.logoutButton}/>) : (<div> </div>)}
            </div>
          </div>
        </div>
      </header>
    );
  }
}

export default Header;