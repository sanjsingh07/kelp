import React, { Component } from 'react';
import styles from './PopoverMenu.module.scss';

class PopoverMenu extends Component {
  render() {
    return (
      <div className={styles.wrapper}>
        <div className={styles.list}>
          {this.props.enableOffers ? <div className={styles.item} onClick={this.props.onOffers}>Show Offers</div> : <div className={[styles.item, styles.disabled].join(' ')}>Show Offers</div>}
          {this.props.enableMarket ? <div className={styles.item} onClick={this.props.onMarket}>Show Market</div> : <div className={[styles.item, styles.disabled].join(' ')}>Show Market</div>}
          {this.props.enableEdit ? <div className={styles.item} onClick={this.props.onEdit}>Edit</div> : <div className={[styles.item, styles.disabled].join(' ')}>Edit</div>}
          {this.props.enableCopy ? <div className={styles.item} onClick={this.props.onCopy}>Copy</div> : <div className={[styles.item, styles.disabled].join(' ')}>Copy</div>}
          {this.props.enableDelete ? <div className={styles.itemDanger} onClick={this.props.onDelete}>Delete</div> : <div className={[styles.item, styles.disabled].join(' ')}>Delete</div>}
        </div>
      </div>
    );
  }
}

export default PopoverMenu;