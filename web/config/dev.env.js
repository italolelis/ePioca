var merge = require('webpack-merge')
var prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  GW_URI: '"https://gw-staging.hellofresh.com"',
  SERVICE_URI: '"http://localhost:8115"',
  WS_URI: '"ws://localhost:8115/ws"',
})
