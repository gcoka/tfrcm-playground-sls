'use strict';

module.exports.hello = async (event, context) => {
  console.log(JSON.stringify(event));

  const request = event.Records[0].cf.request;

  try {
    if (request.uri !== '/index.html') {
      console.log(`changing ${request.uri} to rewrite.html`);
      request.uri = '/rewrite.html';
    }

    return request;
  } catch (error) {
    return error;
  }
};
