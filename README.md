![Threeo.one](./static/banner.png)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Heroku App Status](http://heroku-shields.herokuapp.com/threeo1)](https://your-appname.herokuapp.com)

# About
Threeo.one (***"three-oh-one"***) is an open-sourced URL shortening server/client written in Go. Similar to **TinyURL** or **Bit.ly**, Threeo.one can take very large urls and condense them into easily manageable, bite-sized, and aesthetic short-links of the form: `https://threeo.one/to/<id>`. The name derives from the http status code `301` which redirects traffic to another site.

Link to production build: [https://threeo.one/](https://threeo.one)

# Development
The project is set up for development with [air](https://github.com/cosmtrek/air). Download **air** and run to set up development. In addition, the front-end needs to be built to serve the app: `cd client && yarn build`. This will build the React application and get it ready to be served. Reach out for environment variables required.