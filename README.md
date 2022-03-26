# linux-csgo-mute

Based on [PatrikZero's solution](https://github.com/patrikzudel/PatrikZeros-CSGO-Sound-Fix/), but reimplemented in Go
and written specifically for Linux and Pipewire.

### Usage
Copy `dist/gamestate_integration_CsgoMute.cfg` to the CSGO config directory, by default 
`$HOME/.steam/steam/steamapps/common/Counter-Strike Global Offensive/csgo/cfg`.

Run `./linux-csgo-mute`. Default sound levels can be seen with the `-t` flag. Use the `-C` flag to write the config to 
`${USER_CONFIG_DIR}/csgo-mute/config.toml`; edit the values to set your own levels/options; leave absent for defaults.

The `-h` flag will print all available options.

### Building
Requires Go >= 1.18.

Utilizes wrappers around pipewire cli tools (`pw-dump` and `pw-cli`) for volume management so those need to be available.

System tray support is provided by [getlantern/systray](https://github.com/getlantern/systray), so all CGo dependencies
of that module are required as well, namely `gcc` and development headers for `gtk3` and `libayatana-appindicator3`.
