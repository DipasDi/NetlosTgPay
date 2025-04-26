# NetlosTgDonation 🤖💳  
**Telegram Bot for Automated Donations and Privilege Management**  

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)](https://golang.org/)
[![Telegram Bot API](https://img.shields.io/badge/Telegram%20Bot%20API-6.7+-26A5E4?logo=telegram)](https://core.telegram.org/bots/api)

## ✨ Features
- **Payment processing** via EasyDonate API
- **Auto-privilege delivery** to Minecraft servers via RCON
- **Interactive product catalog** with dynamic pricing
- **Multi-platform support** (Telegram + Discord webhooks)
- **Session management** with thread-safe storage

## 🛠 Tech Stack
**Language**: Go 1.20+ (clean code, goroutines, channels)  
**Libraries**:  
- [go-telegram/bot](https://github.com/go-telegram/bot) - Telegram Bot API framework  
- `net/http` - EasyDonate API integration  
- `sync.Map` - Thread-safe session storage

**Architecture**: Modular (handlers, models, storage)

