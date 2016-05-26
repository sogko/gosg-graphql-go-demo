import React from 'react';
import CSSModules from 'react-css-modules';
import { Link } from 'react-router';

import styles from './app.css';

export class App extends React.Component {
  render() {
    return (
      <div styleName="base">

        <header styleName="heading">
          <Link to="/"><h1>Twooter</h1></Link>
        </header>

        <section styleName="content">
          {this.props.children}
        </section>
      </div>
    );
  }
}

export default CSSModules(App, styles, {
  allowMultiple: true
});
