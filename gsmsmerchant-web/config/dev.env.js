'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    service: {
        NODE_ENV: '"development"',
        service: {
            url: `"http://192.168.5.83:9091"`,
        }
    }
})