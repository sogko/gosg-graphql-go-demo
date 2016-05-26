import React from 'react';
import CSSModules from 'react-css-modules';
import { Link } from 'react-router';

import styles from './user-header.css';

export class UserHeader extends React.Component {
  render() {
    return (
      <header styleName="base">
        <img styleName="avatar" src={this.props.avatar_url} alt="avatar" />
        <h1 styleName="name">{this.props.name}</h1>
        <h2 styleName="handle">@{this.props.handle}</h2>
      </header>
    );
  }
}
UserHeader.propTypes = {
  id: React.PropTypes.number,
  avatar_url: React.PropTypes.string,
  name: React.PropTypes.string,
  handle: React.PropTypes.string
};
export default CSSModules(UserHeader, styles, {
  allowMultiple: true
});
