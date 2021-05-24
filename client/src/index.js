import React from "react";
import ReactDOM from "react-dom";
import { Grommet } from "grommet";
import { grommet } from "grommet/themes";
import { Main } from "./components/main";
import "react-datasheet/lib/react-datasheet.css";

const render = async () => {
  const myTheme = { global: { ...grommet.global, drop: {} } };
  myTheme.global.colors.brand = "#ca27d8";
  myTheme.global.colors.contrast = "#35D827";
  myTheme.global.colors.focus = "#b7b4e1";
  myTheme.global.colors.control = { dark: "#b7b4e1", light: "#b7b4e1" };
  myTheme.global.colors.border = { dark: "#b7b4e1", light: "#b7b4e1" };
  myTheme.global.drop.background = { dark: "#b7b4e1", light: "#b7b4e1" };

  ReactDOM.render(
    <Grommet style={{ height: "100%" }} theme={myTheme}>
      <Main />
    </Grommet>,
    document.getElementById("app"),
  );
};

render();
