/** @type {import('tailwindcss').Config} */
// NOTE how to use requires in ECMA 6
// https://stackoverflow.com/questions/34944099/how-to-import-a-json-file-in-ecmascript-6
import { createRequire } from "module";
const require = createRequire(import.meta.url);

var dark = require("./themes/tokyo-night-dark.json")
var light = require("./themes/tokyo-night-light.json")

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
  mode: "jit",
  purge: ["./src/**/*.{js,svelte,ts}"],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        //DARK_MODE
        darkBg: dark.bg,
        darkStorm: dark.storm,
        darkText: dark.text,
        darkBlack: dark.black,
        darkGreen: dark.green,
        darkRed: dark.red,
        darkYellow: dark.yellow,
        darkBlue: dark.blue,
        darkMagenta: dark.magenta,
        darkWhite: dark.white,
        darkGui: dark.gui,
        darkOrange: dark.orange,
        //LIGHT_MODE
        Bg: light.bg,
        Text: light.text,
        Black: light.black,
        Green: light.green,
        Red: light.red,
        Yellow: light.yellow,
        Blue: light.blue,
        Magenta: light.magenta,
        White: light.white,
        Gui: light.gui,
        Orange: light.orange,
      },
    },
  },
    plugins: []
};
