import React from 'react';
import cssModules from 'react-css-modules';
import { connect } from 'react-apollo';
import gql from 'apollo-client/gql';

import styles from './user-page.css';

import UserHeader from '../components/UserHeader';
import FeedItem from '../components/FeedItem';

export class UserPage extends React.Component {
  renderUserHeader() {
    const { user } = this.props.userPage;
    return <UserHeader {...user} />
  }
  renderFeedItems() {
    const { user } = this.props.userPage;
    const { posts } = user || {};
    if (!posts) {
      return <div styleName="no-items">No posts found.</div>
    }
    return posts.map((post) => {
      return <FeedItem key={post.id} {...post} />;
    });
  }
  render() {
    let { userID } = this.props.params;
    const userHeader = this.renderUserHeader();
    const feedItems = this.renderFeedItems();
    return (
      <div styleName="base">
        <div styleName="header">
          { userHeader }
        </div>
        <section styleName="feed">
          {feedItems}
        </section>
      </div>
    );
  }
}


const UserPageWithStyles = cssModules(UserPage, styles, {
  allowMultiple: true
});

const UserPageWithStylesAndData = connect({
  mapQueriesToProps: ({ ownProps }) => {

    let { userID } = ownProps.params;
    return {
      userPage: {
        query: gql`
          query getUserPage($id: Int!) {
            user(id: $id) {
              id
              name
              handle
              avatar_url
              posts {
                id
                content
                total_likes
                liked_by {
                  id
                  avatar_url
                }
                author {
                  id
                  name
                  handle
                  avatar_url
                }
              }
            }
          }
        `,
        variables: {
          id: userID
        },
        forceFetch: true,
        returnPartialData: true
      }
    };
  }
})(UserPageWithStyles);

export default UserPageWithStylesAndData;
