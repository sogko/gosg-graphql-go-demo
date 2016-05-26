import React from 'react';
import cssModules from 'react-css-modules';
import { connect } from 'react-apollo';
import gql from 'apollo-client/gql';

import styles from './post-page.css';

import Post from '../components/Post';

export class PostPage extends React.Component {
  render() {
    const { post } = this.props.postPage;
    if (!post) {
      return null;
    }
    return (
      <div styleName="base">
        <Post {...post} />
      </div>
    );
  }
}

const PostPageWithStyles = cssModules(PostPage, styles, {
  allowMultiple: true
});

const PostPageWithStylesAndData = connect({
  mapQueriesToProps: ({ ownProps }) => {
    let { postID } = ownProps.params;
    return {
      postPage: {
        query: gql`
          query getPostPage($id: Int!) {
            post(id: $id) {
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
        `,
        variables: {
          id: postID
        },
        forceFetch: true,
        returnPartialData: true
      }
    };
  }
})(PostPageWithStyles);

export default PostPageWithStylesAndData;
