import React from 'react';

import { createStore, combineReducers, applyMiddleware } from 'redux';
import ApolloClient from 'apollo-client';
import { ApolloProvider } from 'react-apollo';

const client = new ApolloClient();

import { render } from 'react-dom';
import { Router, Route, IndexRoute, hashHistory } from 'react-router';

import './global.css';
import App from './App';
import FeedPage from './pages/FeedPage';
import UserPage from './pages/UserPage';
import PostPage from './pages/PostPage';

const store = createStore(
  combineReducers({
    apollo: client.reducer()
  }),
  applyMiddleware(client.middleware())
);

render((
  <ApolloProvider store={store} client={client}>
    <Router
      history={hashHistory}
    >
      <Route path="/" component={App}>
        <IndexRoute component={FeedPage}/>
        <Route path="feed" component={FeedPage} />
        <Route path="user/:userID" component={UserPage}/>
        <Route path="post/:postID" component={PostPage}/>
      </Route>
    </Router>
  </ApolloProvider>

), document.querySelector("#app"));


