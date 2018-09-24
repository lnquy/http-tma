# http-tma
HTTP protocol and server sharing at TMA 2018: [https://goo.gl/TU2nBU](https://goo.gl/TU2nBU)

Code structure (You should read the code from top-down as it's from simpler to more advanced):
- web: HTTP/HTTPS server for serving web contents or REST APIs.
  - basic
  - static-html: Serving static HTML files. 
  - dynamic-template: Serving dynamic HTML template with data binding.
  - api: Serving APIs instead of HTML contents.
    - json: Serving JSON API.
    - xml: Serving XML API.
  - https
    - basic
    - redirect: Runs a child HTTP server on port 80 to redirect all request to HTTPS server. 
- middleware:
  - basic
  - response: Middleware on response.
  - chains: Chaining multiple middlewares together.
- 3th-party-router
  - [chi](https://github.com/go-chi/chi)
  - [mux](https://github.com/gorilla/mux)

### License

This project is under the MIT License. See the [LICENSE](https://github.com/lnquy/http-tma/blob/master/LICENSE) file for the full license text.
