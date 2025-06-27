# 🗂️ Filebrowser Tray Monitor with Animated Icon

This Go application provides a **Windows System Tray** tool to monitor the status of a local `filebrowser.exe` server, automatically start it if it's not running, and display animated icons in the tray to indicate activity. It also allows easy viewing of the server log.

---

## 🚀 Features

- ✅ **Monitors** if `filebrowser.exe` is running
- 🔄 **Auto-starts** the server via `start.bat` if it's not running
- 📄 **Displays recent log** (from `filebrowser.log`) as a tooltip
- 👀 **Opens log file** in a terminal via a tray menu
- 🎞️ **Animated system tray icon** using `.ico` frames

---

## 📁 Directory Structure

```
C:\Windows\filebrowser\
│
├── filebrowser.exe        # The server executable
├── start.bat              # BAT file to launch the server
├── filebrowser.log        # Log file for output
├── img\
│   └── icons\
│       └── animation\
│           ├── anim00.ico
│           ├── anim01.ico
│           └── ...        # Animation frames used in the tray
```

---

## 🔧 Prerequisites

- Go 1.18 or newer
- Windows OS
- Git (optional, for cloning)
- `filebrowser.exe`, `start.bat`, and icons in place

---

## 🛠️ Build Instructions

1. **Clone the repository** (or create your own and copy the script):

   ```bash
   git clone https://github.com/shamim4s/Process-Monitor-from-tray.git
   cd filebrowser-tray
   ```

2. **Place assets** in `C:\Windows\filebrowser\`:
   - `filebrowser.exe`
   - `start.bat`
   - `filebrowser.log`
   - Icon frames in: `C:\Windows\filebrowser\img\icons\animation\`

3. **Download dependencies** (first time only):

   ```bash
   go mod init filebrowser-tray
   go get github.com/getlantern/systray
   go get golang.org/x/sys/windows
   ```

4. **Build the binary**:

   ```bash
   go build -o filebrowser-tray.exe
   ```

5. **Run the application**:

   ```bash
   filebrowser-tray.exe
   ```

---

## 📌 How It Works

- The app checks every 5 seconds if `filebrowser.exe` is running.
- If not, it runs `start.bat` to launch the server.
- It reads the last 10 lines from `filebrowser.log` and shows them as a tooltip.
- The tray icon animates using `.ico` frames from the `animation` folder.
- Right-click the tray icon for:
  - **"Show Log"** — Opens a terminal window showing the log file.
  - **"Exit"** — Terminates the app and log viewer (if open).

---

## 📷 Preview

## 📷 Preview

![Animated Tray Icon Example](https://raw.githubusercontent.com/shamim4s/Process-Monitor-from-tray/master/img/icons/animation/Blocks@1x-1.0s-300px-300px.gif)


---

# 🪟 Go Windows App with PNG Icon

This guide shows how to use a **PNG icon** for your **Windows Go application** by converting it to `.ico` and embedding it during build time using `rsrc`.

---

## 📦 Prerequisites

- [Go](https://golang.org/dl/) installed (1.16+ recommended)
- A `.png` file to use as your application icon
- Windows OS for building or [Cross-compiling](https://github.com/fyne-io/fyne-cross)

---

## 🪄 Step 1: Convert PNG to ICO

Windows requires `.ico` files for executable icons.

You can convert your PNG using:

- **Online tools**:
  - [cloudconvert.com/png-to-ico](https://cloudconvert.com/png-to-ico)
  - [convertico.com](https://convertico.com/)
- **Offline tools**:
  - [GIMP](https://www.gimp.org/)
  - [IrfanView](https://www.irfanview.com/)
  - [ImageMagick](https://imagemagick.org/)

💡 Ensure the `.ico` includes multiple resolutions: **16x16, 32x32, 48x48, 64x64**.

Save your icon as:  
```
icon.ico
```
⚙️ Step 2: Install rsrc

rsrc is used to generate a .syso file embedding the icon into the final binary.

Install it via:
```
go install github.com/akavel/rsrc@latest
```
Make sure $GOPATH/bin is in your PATH.
📄 Step 3: Create rsrc.syso

Run the following in the same directory as your .ico file:
```
rsrc -ico icon.ico
```
This generates a rsrc.syso file that Go will automatically include in the final executable during build.
🛠 Step 4: Build Your Go App

Now build your Go app:
```
go build -ldflags="-H=windowsgui" -o YourApp.exe
```
    -ldflags="-H=windowsgui" hides the terminal window if your app is GUI-based.

    For console apps, you can skip -ldflags.

🗂 Example Project Structure
```
myapp/
├── main.go
├── icon.ico
├── rsrc.syso
```
🧼 Optional: Platform-Specific Resource File

To avoid embedding the icon on non-Windows platforms, rename the .syso file to:
```
rsrc_windows.syso
```
Go will only include it when building on Windows.
✅ Done!

Your built .exe file will now have the icon you provided.

## 📝 License

MIT © 2025 Shamim
