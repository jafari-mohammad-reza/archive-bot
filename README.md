![Language Badge](https://img.shields.io/badge/Language-Go-blue) ![Database Badge](https://img.shields.io/badge/Database-MongoDB-green)

# Archiver Bot

*a simple but hopefully use full telegram bot developed with  golang and mongodb*

**_It is designed to archive and save notes and attachment user wat to save and upload them into cloud for future access with various of formats like html, markdown , code and simple text_**


## Features
- Archive and save messages: The bot is capable of archiving and storing messages from your Telegram account.
- Attachments backup: It stores all the attachments in your Telegram account for future access and usage.
- Compatibility with Dropbox: The bot can save backups in various formats compatible with Dropbox.


## Installation Guide
1. Clone the repository:
``git clone https://github.com/jafari-mohammad-reza/archive-bot``
2. Navigate to the directory:
``cd archiver-bot``
3. run by make file(needs to install mongodb separately):
``make build && make run``
4. or run by docker compose:
``docker-compose up -d``

