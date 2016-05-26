import React from 'react';
import CSSModules from 'react-css-modules';
import { Link } from 'react-router';

import styles from './post.css';

export class Post extends React.Component {
  renderLikedAvatars() {
    return this.props.liked_by.map((liker) => {
      return (
        <Link key={liker.id} to={`/user/${liker.id}`}>
          <span styleName="liked-handler">@{liker.handle}</span>
        </Link>
      );
    });
  }
  render() {
    const likedAvatars = this.renderLikedAvatars();
    return (
      <div styleName="base">
        <div styleName="post">
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

            <div styleName="content-text">
              {this.props.content}
            </div>

            <div styleName="content-footer">
              <div>Reply</div>
              <div>Like</div>
            </div>
          </div>
        </div>

        <div styleName="liked-by-list">

          <div styleName="liked-count">
            <div styleName="liked-count-heading">
              Likes
            </div>
            <div>
              {this.props.total_likes}
            </div>
          </div>
          <div styleName="liked-avatars">
            {likedAvatars}
          </div>
        </div>
      </div>
    );
  }
}
Post.propTypes = {
  id: React.PropTypes.number,
  content: React.PropTypes.string,
  total_likes: React.PropTypes.number,
  author: React.PropTypes.object,
  liked_by: React.PropTypes.array
};
export default CSSModules(Post, styles, {
  allowMultiple: true
});
