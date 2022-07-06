# Give Up GitHub

This project has given up GitHub.  ([See Software Freedom Conservancy's *Give Up  GitHub* site for details](https://GiveUpGitHub.org).)

You can now find this project at [https://gitea.ocram85.com/OCram85/Portainer-ServiceUpdate](https://gitea.ocram85.com/OCram85/Portainer-ServiceUpdate) instead.

Any use of this project's code by GitHub Copilot, past or present, is done without our permission.  We do not consent to GitHub's use of this project's code in Copilot.

Join us; you can [give up GitHub](https://GiveUpGitHub.org) too!

![Logo of the GiveUpGitHub campaign](https://sfconservancy.org/img/GiveUpGitHub.png)

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
