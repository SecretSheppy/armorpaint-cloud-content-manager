# ArmorPaint Cloud Content Manager

apccm allows you to install, update and fully remove the ArmorPaint cloud library of textures
to you local device. Although apccm was primarily created to allow users encountering the 
"Error: Check internet connection to access the cloud" error to still use the cloud content,
it also allows users working offline to download all the cloud resources before going offline.

## Usage

apccm provides 3 simple commands

* [install](#install)
* [update](#update)
* [remove](#remove)

### `install`

The `install` command installs all the cloud content into a specified directory with `/apccm`
appended to the end. It also attempts to create a shortcut in the browser view of ArmorPaint
to that location (though at time of writing this is only available on linux).

The following example installs the content to `/your/install/location/apccm`

```console
apccm install /your/install/location
```

The `install` command also creates a `.assets_list.json` file in the installation directory.
This file contains a JSON formatted version of the internal `AssetList` datatype that is used
by the `update` command to prevent the application from downloading assets that have not been
changed.

### `update`

The `update` command updates any out of date assets currently installed and downloads any missing
or new assets.

```console
apccm update /your/install/location
```

### `remove`

The `remove` command removes all assets from your local system. This is irreversible, meaning
if you need to access the assets again you will have to install them all again.

```console
apccm remove /your/install/location
```
