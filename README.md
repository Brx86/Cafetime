# CafeTime — 曼城时光 ☕

A Windows system tray countdown timer. Set a timer (50/30/10 min), and when it hits zero — lock your screen, open a file, or launch a website. Inspired by [Catime](https://github.com/vladelaina/Catime).

## How it works

```
              ┌──────────────────────┐
 tray icon ◄──┤  systray (background) │
    ↑         └──────┬───────────────┘
    │                │ menu click
    │         ┌──────▼───────────────┐
    │         │  OnClickMenu         │
    │         │  (select loop)       │
    │         └──────┬───────────────┘
    │         ┌──────▼───────────────┐     ┌──────────────────┐
    └──────┬──┤  NewTimer(sec)       ├────►│  actionFunc()    │
           │  │  (ticker every 1s)   │     │  lock / open     │
           │  └──────────────────────┘     └──────────────────┘
           │         ▲
           │         │ cancel
      ┌────┴─────────┴────┐
      │ stopTimerCh       │
      └───────────────────┘
```

- **Countdown** runs in its own goroutine, decrementing every second.
- **Cancel** at any time via the menu — a dedicated channel signals the goroutine to exit early.
- **Action** is a pluggable function pointer: lock screen, open a file, or open a `.lnk` shortcut.

## Usage

1. **Build** — requires Go 1.24+ on Windows.

   ```
   build.bat
   ```

2. **Run** — a coffee cup icon appears in the system tray.

   | Action | How |
   |---|---|
   | Pick a timeout | Click "定时 50/30/10 min" |
   | Cancel | Click "剩余 XX，点击取消" |
   | Choose what happens | "超时动作" → lock / open file / open website |
   | Quit | Click "退出" |

## Project structure

```
cafetime/
├── main.go              # Entry point, tray lifecycle, signal handling
├── menu.go              # Tray menu layout & click event loop
├── timer.go             # Countdown engine + time formatting
├── actions/
│   ├── lock.go          # Lock screen (LockWorkStation) + DPI awareness
│   ├── run.go           # Open file/program/shortcut
│   └── win32.go         # Win32 API bindings (User32.dll)
├── asset/               # Tray icons (coffee cup on/off)
├── build.bat            # Windows build script
└── .github/workflows/   # CI
```

## Tech

- **[fyne.io/systray](https://github.com/fyne-io/systray)** — system tray
- **[sqweek/dialog](https://github.com/sqweek/dialog)** — native file picker
- **syscall + User32.dll** — `LockWorkStation` & DPI awareness (no CGo)

## License

MIT
