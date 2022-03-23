# linux-csgo-mute

Based on [PatrikZero's solution](https://github.com/patrikzudel/PatrikZeros-CSGO-Sound-Fix/), but reimplemented in Go
and written specifically for Linux and Pipewire.

### Usage
Copy `dist/gamestate_integration_CsgoMute.cfg` to the CSGO config directory, by default 
`$HOME/.steam/steam/steamapps/common/Counter-Strike Global Offensive/csgo/cfg`.

Run `./linux-csgo-mute`. Default sound levels can be seen in `config/config.go`. Copy `dist/config.toml` to `${USER_CACHE_DIR}/csgo-mute/config.toml`
to set your own levels/options; leave absent for defaults.

### Building
Requires Go >= 1.18.

Utilizes wrappers around pipewire cli tools for volume management so those need to be available.
