import {
  AssignmentTurnedIn,
  Forum,
  Home,
  ViewTimeline,
} from "@mui/icons-material";
import {
  Avatar,
  Box,
  Button,
  Divider,
  Link,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography,
} from "@mui/material/";
import { ThemeProvider } from "@mui/material/styles";
import React, { useState } from "react";

import logo from "../asset/logo.png";
import { styles, theme } from "../asset/style";

const customStyle = {
  brandLogo: {
    display: "flex",
    justifyContent: "center",
    margin: "16px",
  },
  link: {
    display: "flex",
    alignItems: "center",
  },
  listIcon: {
    marginLeft: "20px",
  },
  listItem: {
    margin: "30px 0 30px 60px",
  },
  backgroundDiv1: {
    height: "calc(100%)",
    width: "285px",
    zIndex: "-2",
    margin: "0 0 0 0",
    borderRadius: "0 32.8px 32.8px 0",
    backgroundColor: "#84B140",
  },
  backgroundDiv2: {
    height: "calc(100%)",
    width: "285px",
    zIndex: "-1",
    margin: "0 0 0 0",
    borderRadius: "0 32.8px 32.8px 0",
    backgroundColor: "#E1E3E4",
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
        <Box>
          <Box sx={customStyle.brandLogo}>
            <img src={logo} alt="logo" width={64} height={64} />
            <Link href="/" underline="none" sx={customStyle.link}>
              <Typography variant="h2" component="div" sx={styles.topNavHeader}>
                TutorKit
              </Typography>
            </Link>
          </Box>
          <List>
            {sideMenuList.map((listItem, index) => (
              <ListItemButton
                disableRipple
                key={index}
                sx={customStyle.listItem}
                selected={selectedIndex === index}
                onClick={() => handleListItemClick(index)}
              >
                <ListItemIcon sx={customStyle.listIcon}>
                  {listItem.listIcon}
                </ListItemIcon>
                <ListItemText>
                  <Typography variant="h1" sx={styles.inputLabel}>
                    {listItem.listText}
                  </Typography>
                </ListItemText>
              </ListItemButton>
            ))}
          </List>
        </Box>
        <Box sx={{ height: "200px" }}>
          <Divider variant="middle" />
          <Box
            sx={{
              margin: "0 30px 16px 30px",
            }}
          >
            <Box
              sx={{
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
              }}
            >
              <Avatar>T</Avatar>
              <Box sx={{ marginLeft: "8px" }}>
                <Typography
                  variant="subtitle1"
                  component="div"
                  sx={{ display: "flex" }}
                >
                  Test User
                </Typography>
                <Typography
                  variant="subtitle2"
                  component="div"
                  sx={{ display: "flex" }}
                >
                  testuser@gmail.com
                </Typography>
              </Box>
            </Box>
          </Box>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              margin: "16px 30px 16px 30px",
            }}
          >
            <Button variant="contained" disableRipple={true}>
              Settings
            </Button>
            <Button variant="contained" disableRipple={true}>
              Log Out
            </Button>
          </Box>
        </Box>
      </Box>
      <Box
        component="div"
        sx={{ ...styles.sideNavContainer, ...customStyle.backgroundDiv1 }}
      ></Box>
      <Box
        component="div"
        sx={{ ...styles.sideNavContainer, ...customStyle.backgroundDiv2 }}
      ></Box>
    </ThemeProvider>
  );
};

export default SideNavBar;
