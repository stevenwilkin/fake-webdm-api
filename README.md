# fake-webdm-api

An app to simulate the API exposed by [WebDM](https://launchpad.net/webdm).

## Running

There are no dependencies so just run:

	go run *.go

The app is then available at [ http://0.0.0.0:8080]( http://0.0.0.0:8080).

## Endpoints

### All available snaps

	http://0.0.0.0:8080/api/v2/packages/

### Installed snaps

	http://0.0.0.0:8080/api/v2/packages/?installed_only=true

### Individual snaps

	http://0.0.0.0:8080/api/v2/packages/{SNAP}

### Icons

	http://0.0.0.0:8080/icons/{SNAP}