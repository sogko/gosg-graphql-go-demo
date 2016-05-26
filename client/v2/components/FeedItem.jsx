import React from 'react';
import CSSModules from 'react-css-modules';
import { Link } from 'react-router';

import { hashHistory } from 'react-router'

import styles from './feed-item.css';

import LikedAvatarImg from './LikedAvatarImg';

export class FeedItem extends React.Component {
  onClickItem(evt) {
    evt.preventDefault();
    hashHistory.push(`/post/${this.props.id}`);
    return false;

  }
  renderLikedAvatars() {
    if (!this.props.liked_by) {
      return null;
    }
    return this.props.liked_by.map((liker) => {
      return (
        <LikedAvatarImg key={liker.id} {...liker} />
      );
    });
  }
  render() {
    const likedAvatars = this.renderLikedAvatars();
    return (
      <div styleName="base">
        <div styleName="avatar">
          <img styleName="avatar-img" src={this.props.author.avatar_url} />
        </div>
        <div styleName="content">
          <div styleName="content-header">

            <div styleName="author-name">
              <Link to={`/user/${this.props.author.id}`}>{this.props.author.name}</Link>
            </div>

            <div styleName="author-handle">
              <Link to={`/user/${this.props.author.id}`}>@{this.props.author.handle}</Link>
            </div>

            <div styleName="date">
              <Link to={`/post/${this.props.id}`}>May 20</Link>
            </div>
          </div>

          <div styleName="content-text" onClick={(e) => this.onClickItem(e)}>
            {this.props.content}
          </div>

          <div styleName="content-footer">
            <div>Reply</div>
            <div>Like</div>
            <div>{this.props.total_likes} likes</div>
            <div>{likedAvatars}</div>
          </div>
        </div>
      </div>
    );
  }
}
FeedItem.propTypes = {
  id: React.PropTypes.number,
  content: React.PropTypes.string,
  total_likes: React.PropTypes.number,
  author: React.PropTypes.object
};
export default CSSModules(FeedItem, styles, {
  allowMultiple: true
});
