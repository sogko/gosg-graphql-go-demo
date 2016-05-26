import React from 'react';
import { connect } from 'react-apollo';
import gql from 'apollo-client/gql';

import cssModules from 'react-css-modules';

import FeedItem from './../components/FeedItem'
import styles from './feed-page.css';

export class FeedPage extends React.Component {
  renderFeedItems() {
    if (this.props.feedItems && this.props.feedItems.loading) {
      return <div styleName="no-items">Loading posts...</div>
    }
    if (!this.props.feedItems || !this.props.feedItems.posts) {
      return <div styleName="no-items">No posts found.</div>
    }
    return this.props.feedItems.posts.map((post) => {
      return <FeedItem key={post.id} {...post} />;
    });
  }
  render() {
    const feedItems = this.renderFeedItems();
    return (
      <div styleName="base">
        <section styleName="feed">
          {feedItems}
        </section>
      </div>
    );
  }
}

const FeedPageWithStyles = cssModules(FeedPage, styles, {
  allowMultiple: true
});

const FeedPageWithStylesAndData = connect({
  mapQueriesToProps: () => {
    return {
      feedItems: {
        query: gql`
          query getFeedPage($limit: Int!) {
            posts(limit: $limit) {
              id
              content
              total_likes
              author {
                id
                name
                handle
                avatar_url
              }
            }
          }
        `,
        variables: {
          limit: 15
        }
      }
    };
  }
})(FeedPageWithStyles);

export default FeedPageWithStylesAndData;
