const {createProxyMiddleware} = require('http-proxy-middleware');
module.exports = function(app) {
    app.use('/department', createProxyMiddleware( 
        { target: 'http://localhost:3000/' }
    ));
}