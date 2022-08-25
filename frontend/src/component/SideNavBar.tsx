import {
  AssignmentTurnedIn,
  Forum,
  Home,
  ViewTimeline,
} from "@mui/icons-material";
import {
  Box,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography,
} from "@mui/material/";
import { ThemeProvider } from "@mui/material/styles";
import React, { useState } from "react";

import { styles, theme } from "../asset/style";
import Footer from "./Footer";

const customStyle = {
  listItem: {
    margin: "8px 0 8px 0",
  },
};

const sideMenuList = [
  {
    listIcon: <Home />,
    listText: "Course",
  },
  {
    listIcon: <ViewTimeline />,
    listText: "Todo",
  },
  {
    listIcon: <AssignmentTurnedIn />,
    listText: "Attendence",
  },
  {
    listIcon: <Forum />,
    listText: "Community",
  },
];

export interface ISideNavBarProps {}

const SideNavBar = () => {
  const [selectedIndex, setSelectedIndex] = useState<number>(0);
  const handleListItemClick = (index: number) => {
    setSelectedIndex(index);
  };
  return (
    <ThemeProvider theme={theme}>
      <Box component="div" sx={styles.sideNavContainer}>
        <List>
          {sideMenuList.map((listItem, index) => (
            <ListItemButton
              disableRipple
              key={index}
              sx={customStyle.listItem}
              selected={selectedIndex === index}
              onClick={() => handleListItemClick(index)}
            >
              <ListItemIcon>{listItem.listIcon}</ListItemIcon>
              <ListItemText>
                <Typography variant="h1" sx={styles.inputLabel}>
                  {listItem.listText}
                </Typography>
              </ListItemText>
            </ListItemButton>
          ))}
        </List>
        <Box>
          <Footer position="relative" />
        </Box>
      </Box>
    </ThemeProvider>
  );
};

export default SideNavBar;
