# Foundry VTT Docker image

You can run this image with docker run abibby/foundryvtt

# Updating

To update the foundry version download the latest node.js version of foundry
from [here](https://foundryvtt.com/community/zwzn/licenses) and update the
version number in the `Dockerfile`. Finaly run `make VERSION=<version>` with the
new version