var https = require('https')
const request = require('request')

exports.handler = (event, context) => {

  try {

    if (event.session.new) {
      // New Session
      console.log("NEW SESSION")
    }

    switch (event.request.type) {

      case "LaunchRequest":
        // Launch Request
        console.log(`LAUNCH REQUEST`)
        context.succeed(
          generateResponse(
            buildSpeechletResponse("Welcome to bosserator, your personal bosh deployment robot. please choose to deploy or delete.", true),
            {}
          )
        )
        break;

      case "IntentRequest":
        // Intent Request
        console.log(`INTENT REQUEST`)

        switch(event.request.intent.name) {
          case "deploy":
          var body = {
            "text": "deploy server"
          }
          const options = {
            headers: {
              "Content-Type": "application/json"
            },
            url: process.env.URL,
            method: "POST",
            json: true,
            body: body
          }
          request(options, (error, response, body) => {
            if ( error ) {
              context.fail("Deploy failed: " + error)
              console.log(error)
            }
            context.succeed(
              generateResponse(
                buildSpeechletResponse("Server deployed.", true), {}
              )
            )
            console.log(response)
          })
          break;
          case "test":
            context.succeed(
              generateResponse(
                buildSpeechletResponse("this is a test", true), {}
              )
            )
          break;
          case "panic":
            context.succeed(
              generateResponse(
                buildSpeechletResponse("A very wise man once said, don't panic. i suggest you follow his advise.", true), {}
              )
            )
          break;

          default:
            throw "Invalid intent"
        }

        break;

      case "SessionEndedRequest":
        // Session Ended Request
        console.log(`SESSION ENDED REQUEST`)
        break;

      default:
        context.fail(`INVALID REQUEST TYPE: ${event.request.type}`)

    }

  } catch(error) { context.fail(`Exception: ${error}`) }

}

// Helpers
buildSpeechletResponse = (outputText, shouldEndSession) => {

  return {
    outputSpeech: {
      type: "PlainText",
      text: outputText
    },
    shouldEndSession: shouldEndSession
  }

}

generateResponse = (speechletResponse, sessionAttributes) => {

  return {
    version: "1.0",
    sessionAttributes: sessionAttributes,
    response: speechletResponse
  }

}

