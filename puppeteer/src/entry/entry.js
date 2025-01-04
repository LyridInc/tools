'use strict';

// eslint-disable-next-line import/no-unresolved
const express = require('express');
const app = express();
const api = require('./api')
var cors = require('cors')
var corsOptions = {
    origin: [ 'https://localhost:3000', 'http://localhost:3000', 'https://app.staging.lyrid.io', 'https://app.lyrid.io' ],
    credentials: true,
  }

var morgan = require('morgan')
app.use(morgan('combined'))
app.use(cors(corsOptions))
app.use('/', api)

// Error handler
//app.use((err, req, res) => {
//  res.status(500).send('Internal Serverless Error: ' + err);
//});

module.exports = app;
