# webhook-to-bash

A simple API server that triggers a batch file when asked, based on environment variables

### Environment variables to set
Variable | Use
--- | ---
envTag | The HTTP header tag name used to pass the authentication token, e.g. `X-Gitlab-Token`
envValidToken | A token that is valid for the client to send - the API will return forbidden if it receives any other token
envPort | the TCP port to run the API listener on
envBashfile | the full path of the Bash file to run

### API Endpoints

Endpoint | Use
--- | ---
/ | simple "Ok" response to confirm API is up and running
/version | returns information on the API version
/trigger | triggers the relevant Bash file and returns a job ID
/status/{`job ID`} | returns status of job *NB. Not yet implemented*

Note: All endpoints other than `/` and `/version` require the correct authentication token