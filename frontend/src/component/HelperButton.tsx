import { Create } from "@mui/icons-material";
import { SpeedDial, SpeedDialAction, SpeedDialIcon } from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import React, { useState } from "react";

import { theme } from "../asset/style";

const customStyle = {
  mainButton: {
    position: "absolute",
    bottom: "32px",
    right: "32px",
  },
  tooltip: {
    backgroundColor: "#F0F0F0",
    color: "#000000",
    fontFamily: "rubik",
    fontSize: "16px",
  },
};

const actions = [{ icon: <Create />, name: "New" }];

export interface IMainButtonProps {
  currentPage: string;
}

export const MainButton: React.FC<IMainButtonProps> = ({ currentPage }) => {
  const [open, setOpen] = useState<boolean>(false);
  const toggleOpen = () => setOpen(!open);
  return (
    <ThemeProvider theme={theme}>
      <SpeedDial
        ariaLabel="SpeedDial for action"
        sx={customStyle.mainButton}
        icon={<SpeedDialIcon />}
        onClose={toggleOpen}
        onOpen={toggleOpen}
        open={open}
      >
        {actions.map((action) => (
          <SpeedDialAction
            key={action.name}
            icon={action.icon}
            tooltipTitle={action.name + ` ${currentPage}`}
            onClick={toggleOpen}
            componentsProps={{
              tooltip: {
                sx: {
                  ...customStyle.tooltip,
                },
              },
            }}
          />
        ))}
      </SpeedDial>
    </ThemeProvider>
  );
};
