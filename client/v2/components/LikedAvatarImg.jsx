import React from 'react';
import CSSModules from 'react-css-modules';
import { Link } from 'react-router';

import styles from './liked-avatar-img.css';

export class LikedAvatarImg extends React.Component {
  render() {
    console.log(this.props);
    return (
      <Link to={`/user/${this.props.id}`}>
        <img styleName="base" src={this.props.avatar_url} />
      </Link>
    );
  }
}
LikedAvatarImg.propTypes = {
  id: React.PropTypes.number,
  avatar_url: React.PropTypes.string
};
export default CSSModules(LikedAvatarImg, styles, {
  allowMultiple: true
});
