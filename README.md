# CRON_BACKUP

## Table of contents

1. [Prerequisties](#prerequisties)
2. [Usage](#usage)
3. [Configuration](#configuration)

## Prerequisties

To use this project, you need to have Go installed on your system. Ensure you have the latest stable version of Go, which can be downloaded from the [official Go website](https://golang.org/dl/).

## Usage

To use CRON_BACKUP, configure the `settings.json` file by specifying the target path, the desired schedule, and whether you want a full backup or an incremental one. You can choose between a total backup or a differential backup based on your needs. For more details, refer to the [settings.json](settings.json) file.

## Configuration

To configure CRON_BACKUP, you have to edit the [settings.json](settings.json) file.

| Entry        | Exemple                                                                                               |
| ------------ | ----------------------------------------------------------------------------------------------------- |
| backupdir    | The dir path of the folder you want to backup. i.e. `.` for current folder (based on executable path) |
| outdir       | The folder where you want to stock you backups                                                        |
| delay        | The delay of the execution (43200m by default, 1 month)                                               |
| differential | For differencial backuping (not available yet)                                                        |
