## Discounter - Discord Bot

Discord bot to increase ranking on the server by sending number++ messages.

## Quick Start

```bash
$ chmod +x discounter # for linux and macOs
$ echo "TOKEN=YOUR_DISCORD_TOKEN" > .env
$ discounter -c 12345678 -n 123 -d 5
```

### Params

- `-c` channel id
- `-n` number for start counter. Default: 0
- `-d` delay for send messages. Default: 5

### How to retrieve access token:
- First, please open your Discord in a browser
- Go to settings > Advanced
- Enable Developer mode
- Press F12 or right click select "Inspect" (Chrome)
- Then click "Application"
- Select Local Storage, Select Discord.com
- Type "Token" in the input filter

## License

This open-sourced software is licensed under the [MIT license](LICENSE).
