# Steam

Steam themeing is done via [Adwaita-for-Steam](https://github.com/tkashkin/Adwaita-for-Steam) at the commit hash [171aa65](https://github.com/tkashkin/Adwaita-for-Steam/commit/171aa65e8d186cf0fe723cafb4498d10f1ae8baf)

It **will not be installed for you**, as this package does not take that liberty.
There will be an error message instructing you on how to do it, but here it is below, if you'd like to run it yourself right now:

```sh
git clone https://github.com/tkashkin/Adwaita-for-Steam --revision 171aa65e8d186cf0fe723cafb4498d10f1ae8baf ~/.local/share/steam-adwaita
```

There is an expected path - `~/.local/share/steam-adwaita`.
However, feel free to keep clone to wherever you see fit.
Just create a symlink to the above path

As of the above hash, Adwaita-for-Steam uses hex-code and light/dark mode.
The assumption (as of the time of writing this) is that Omarchy themeing only allows for one colour palette and you can set either light or dark mode.
Therefore, whatever the theme set for Omarchy, those colours will be used as both "light" and "dark" mode.
