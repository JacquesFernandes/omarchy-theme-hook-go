# omarchy-theme-hook-go

A hobbyist attempt at trying to port the [omarchy-theme-hook](https://github.com/imbypass/omarchy-theme-hook) to Go and also bring it up-to-speed with the theme system in Omarchy 4

## Mapping

"Old" -> the naming convention in the existing omarchy 3 theme-hook package
"New" -> the mapped naming convention in omarchy 4 and this package

| Old | New |
| -------------- | --------------- |
| `primary_background` | `background` |
| `primary_foreground` | `foreground` |
| `normal_red` | `red` |
| `normal_green` | `green` |
| `normal_yellow` | `yellow` |
| `normal_blue` | `blue` |
| `normal_magenta` | `magenta` |
| `normal_cyan` | `cyan` |
| `normal_white` | `light_foreground` |
| `normal_black` | `background` |
| `bright_black` | `lighter_background` |
| `bright_red` | `bright_red` |
| `bright_green` | `bright_green` |
| `bright_yellow` | `bright_yellow` |
| `bright_blue` | `bright_blue` |
| `bright_magenta` | `bright_magenta` |
| `bright_cyan` | `bright_cyan` |
| `bright_white` | `bright_foreground` |

## To-do

I've ordered this based on (personal) priority, whether it comes pre-installed on omarchy, and then everything else

- [x] gtk
- [-] steam
  - Currently working on adapting css template to commit [171aa65](https://github.com/tkashkin/Adwaita-for-Steam/commit/171aa65e8d186cf0fe723cafb4498d10f1ae8baf) of Adwaita-for-Steam
- [] discord
- [] spotify
- [] typora
- [] vscode

- [] fish
- [] fzf
- [] qt6ct
- [] superfile
- [] vicinae
- [] nwg-dock-hyprland
- [] zed
- [] cursor
- [] windsurf
- [] cava
- [] firefox
- [] qutebrowser
- [] zen
- [] heroic
