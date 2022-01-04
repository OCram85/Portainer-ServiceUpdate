# Portainer-ServiceUpdate

A simple Portainer Service Updater

## About

You can use this plugin for updating any Portainer Service within a drone.io pipeline.

-   Enable the `Service webhook` feature in your portainer service definition.
    -   copy the token part in the shown url.

Add a Drone.IO secret in your project named `TOKEN`

Add this step to your pipeline to trigger the update:

```
- name: "Trigger Service Update"
  image: ocram85/portainer-serviceupdate
  settings:
    VERBOSE: true
    URI: "https://your-portainer-url.com"
    TOKEN:
      from_secret: TOKEN
```
