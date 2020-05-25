# WoW AH dump

This command line tool allows you to download World of Warcraft auction house dumps in batches.

## Usage

Compile the source code or download a pre compiled binary for your OS and run:

```bash
ah-dump -key <YOUR APP KEY> -secret <YOUR APP SECRET> [-output <DUMP OUTPUT DIR>] [-region <API REGION>] [-concurrency <NUMBER>]
``` 

Command line options:
* -key **required** - Your client credential key.
* -secret **required** - Your client credential secret.
* -output **optional** - Directory where files will be saved. Defaults to `./dump`.
* -region **optional** - The API region for downloading the data. Defaults to `us`.
* -concurrency **optional** - How many concurrent downloads to run. Defaults to 100 (Max allowed). If you have trouble on a slow connection try decreasing this number.

To obtain your client key and secret register at: https://develop.battle.net